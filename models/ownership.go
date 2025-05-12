package models

import (
    "fmt"
    "strings"
    "webtechproject/config"
)

type OwnershipAssignment struct {
    UserID     int     `json:"user_id"`
    LocationID *int    `json:"location_id,omitempty"`
    TypeID     *int    `json:"type_id,omitempty"`
    StartDate  string  `json:"start_datetime"`          // e.g., "2025-05-08 10:00:00"
    EndDate    *string `json:"end_datetime,omitempty"`  // nullable
}


func GetAllChildLocationIDs(parentID int) ([]int, error) {
    query := `
        WITH RECURSIVE location_tree AS (
            SELECT location_id FROM Location WHERE location_id = ?
            UNION ALL
            SELECT l.location_id FROM Location l
            INNER JOIN location_tree lt ON lt.location_id = l.parent_location_id
        )
        SELECT location_id FROM location_tree;
    `
    rows, err := config.DB.Query(query, parentID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var ids []int
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return nil, err
        }
        ids = append(ids, id)
    }
    return ids, nil
}

func AssignOwnershipBulk(input OwnershipAssignment) error {
    var filters []string
    var args []interface{}

    if input.LocationID != nil {
        locIDs, err := GetAllChildLocationIDs(*input.LocationID)
        if err != nil {
            return err
        }
        if len(locIDs) > 0 {
            placeholders := strings.Repeat("?,", len(locIDs))
            filters = append(filters, fmt.Sprintf("location_id IN (%s)", placeholders[:len(placeholders)-1]))
            for _, id := range locIDs {
                args = append(args, id)
            }
        }
    }

    if input.TypeID != nil {
        filters = append(filters, "type_id = ?")
        args = append(args, *input.TypeID)
    }

    filterQuery := "SELECT device_id FROM Device WHERE status_flag = 1"
    if len(filters) > 0 {
        filterQuery += " AND " + strings.Join(filters, " AND ")
    }

    rows, err := config.DB.Query(filterQuery, args...)
    if err != nil {
        return err
    }
    defer rows.Close()

    var deviceIDs []int
    for rows.Next() {
        var deviceID int
        if err := rows.Scan(&deviceID); err != nil {
            return err
        }
        deviceIDs = append(deviceIDs, deviceID)
    }

    tx, err := config.DB.Begin()
    if err != nil {
        return err
    }

    for _, deviceID := range deviceIDs {
        // Set old ownerships to inactive
        _, err := tx.Exec("UPDATE Ownership SET status = 0 WHERE device_id = ? AND status = 1", deviceID)
        if err != nil {
            tx.Rollback()
            return err
        }

        // Insert new ownership
        _, err = tx.Exec(`
            INSERT INTO Ownership (device_id, user_id, start_datetime, end_datetime, status)
            VALUES (?, ?, ?, ?, 1)
        `, deviceID, input.UserID, input.StartDate, input.EndDate)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    err = tx.Commit()
    if err != nil {
        return err
    }

    return nil
}

