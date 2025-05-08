package models

import (
    "fmt"
	// "log"
    "strings"
	"database/sql"
    "webtechproject/config"
)

type DeviceFilter struct {
    CategoryID   *int    `json:"category_id"`
    TypeID       *int    `json:"type_id"`
    UserID       *int    `json:"user_id"`       // ownership
    OwnerName    *string `json:"owner_name"`    // ownership by full name
    AddedBy      *int    `json:"added_by"`
    LocationID   *int    `json:"location_id"`
}

type DeviceView struct {
    DeviceID     int     `json:"device_id"`
    DeviceName   string  `json:"device_name"`
    Category     string  `json:"category_name"`
    Type         string  `json:"type_name"`
    SerialNo     string  `json:"serial_no"`
    ModelNo      string  `json:"model_no"`
    Owner        string  `json:"owner"`
    AddedByName  string  `json:"added_by_name"`
    Location     string  `json:"location"`
}

func GetAllChildLocations(locID int) ([]int, error) {
    var ids []int
    query := `
        WITH RECURSIVE loc_hierarchy AS (
            SELECT location_id FROM Location WHERE location_id = ?
            UNION ALL
            SELECT l.location_id FROM Location l
            INNER JOIN loc_hierarchy lh ON lh.location_id = l.parent_location_id
        )
        SELECT location_id FROM loc_hierarchy;
    `
    rows, err := config.DB.Query(query, locID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()
    for rows.Next() {
        var id int
        if err := rows.Scan(&id); err != nil {
            return nil, err
        }
        ids = append(ids, id)
    }
    return ids, nil
}

func FetchFilteredDevices(filter DeviceFilter) ([]DeviceView, error) {
    query := `
        SELECT 
            d.device_id, d.device_name, dc.category_name, dt.type_name,
            d.serial_no, d.model_no, 
            u1.full_name as owner,
            u2.full_name as added_by,
            l.location_name
        FROM Device d
        JOIN Device_Type dt ON d.type_id = dt.type_id
        JOIN Device_Category dc ON dt.category_id = dc.category_id
        JOIN Location l ON d.location_id = l.location_id
        LEFT JOIN Ownership o ON d.device_id = o.device_id AND o.status = 1
        LEFT JOIN User u1 ON o.user_id = u1.user_id
        LEFT JOIN User u2 ON d.added_by = u2.user_id
        WHERE d.status_flag = 1
    `
    var args []interface{}
    var conditions []string

    if filter.CategoryID != nil {
        conditions = append(conditions, "dc.category_id = ?")
        args = append(args, *filter.CategoryID)
    }

    if filter.TypeID != nil {
        conditions = append(conditions, "dt.type_id = ?")
        args = append(args, *filter.TypeID)
    }

    if filter.UserID != nil {
        conditions = append(conditions, "u1.user_id = ?")
        args = append(args, *filter.UserID)
    }

    if filter.OwnerName != nil {
        conditions = append(conditions, "u1.full_name LIKE ?")
        args = append(args, "%"+*filter.OwnerName+"%")
    }

    if filter.AddedBy != nil {
        conditions = append(conditions, "d.added_by = ?")
        args = append(args, *filter.AddedBy)
    }

    if filter.LocationID != nil {
        locIDs, err := GetAllChildLocations(*filter.LocationID)
        if err != nil {
            return nil, err
        }
        if len(locIDs) > 0 {
            placeholders := strings.TrimSuffix(strings.Repeat("?,", len(locIDs)), ",")
            conditions = append(conditions, fmt.Sprintf("d.location_id IN (%s)", placeholders))
            for _, id := range locIDs {
                args = append(args, id)
            }
        }
    }

    if len(conditions) > 0 {
        query += " AND " + strings.Join(conditions, " AND ")
    }

    rows, err := config.DB.Query(query, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var devices []DeviceView
    for rows.Next() {
        var (
            d                DeviceView
            owner            sql.NullString
            addedByFullName  sql.NullString
        )

        if err := rows.Scan(
            &d.DeviceID, &d.DeviceName, &d.Category, &d.Type,
            &d.SerialNo, &d.ModelNo, &owner, &addedByFullName, &d.Location,
        ); err != nil {
            return nil, err
        }

        d.Owner = stringOrEmpty(owner)
        d.AddedByName = stringOrEmpty(addedByFullName)

        devices = append(devices, d)
    }

    return devices, nil
}




func stringOrEmpty(ns sql.NullString) string {
    if ns.Valid {
        return ns.String
    }
    return ""
}
