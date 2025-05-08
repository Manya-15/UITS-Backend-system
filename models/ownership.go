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
    var deviceArgs []interface{} // For filters only

    if input.LocationID != nil {
        locIDs, err := GetAllChildLocationIDs(*input.LocationID)
        if err != nil {
            return err
        }
        if len(locIDs) > 0 {
            placeholders := strings.Repeat("?,", len(locIDs))
            filters = append(filters, fmt.Sprintf("d.location_id IN (%s)", placeholders[:len(placeholders)-1]))
            for _, id := range locIDs {
                deviceArgs = append(deviceArgs, id)
            }
        }
    }

    if input.TypeID != nil {
        filters = append(filters, "d.type_id = ?")
        deviceArgs = append(deviceArgs, *input.TypeID)
    }

    baseQuery := `
        INSERT INTO Ownership (device_id, user_id, start_datetime, end_datetime, status)
        SELECT d.device_id, ?, ?, ?, 1
        FROM Device d
        WHERE d.status_flag = 1
          AND d.device_id NOT IN (
              SELECT device_id FROM Ownership WHERE status = 1
          )
    `
    if len(filters) > 0 {
        baseQuery += " AND " + strings.Join(filters, " AND ")
    }

    // Final argument order: userID, start_datetime, end_datetime, [device filters...]
    finalArgs := []interface{}{input.UserID, input.StartDate}
    if input.EndDate != nil {
        finalArgs = append(finalArgs, *input.EndDate)
    } else {
        finalArgs = append(finalArgs, nil)
    }
    finalArgs = append(finalArgs, deviceArgs...)

    _, err := config.DB.Exec(baseQuery, finalArgs...)
    return err
}


