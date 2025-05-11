package models

import (
    "webtechproject/config"
)

type Location struct {
    ID       int     `json:"id"`
    Name     string  `json:"name"`
    LevelID  int     `json:"level_id"`
    ParentID *int    `json:"parent_id"`
}

func FetchAllLocations() ([]Location, error) {
    rows, err := config.DB.Query("SELECT location_id, location_name, level_id, parent_location_id FROM Location")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var locations []Location
    for rows.Next() {
        var loc Location
        if err := rows.Scan(&loc.ID, &loc.Name, &loc.LevelID, &loc.ParentID); err != nil {
            return nil, err
        }
        locations = append(locations, loc)
    }
    return locations, nil
}
