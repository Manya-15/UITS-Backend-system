package models

import (
    "log"
    "webtechproject/config"
)

type DeviceCategory struct {
    ID   int    `json:"id"`
    Name string `json:"category_name"`
}

type DeviceType struct {
    ID   int    `json:"id"`
    Name string `json:"typename"`
}

type DeviceIDResponse struct {
    DeviceID int `json:"deviceId"`
}


// type DeviceInput struct {
//     DeviceName     string `json:"device_name"`
//     TypeID         int    `json:"type_id"`
//     SerialNo       string `json:"serial_no"`
//     ModelNo        string `json:"model_no"`
//     PurchaseDate   string `json:"purchase_date"`
//     WarrantyExpiry string `json:"warranty_expiry"`
//     LocationID     int    `json:"location_id"`
//     ReplacedID     *int   `json:"replaced_device_id"`
// }

type DeviceInput struct {
    DeviceName       string  `json:"device_name"`
    TypeID           int     `json:"type_id"`
    SerialNo         string  `json:"serial_no"`
    ModelNo          string  `json:"model_no"`
    PurchaseDate     string  `json:"purchase_date"`
    WarrantyExpiry   string  `json:"warranty_expiry"`
    LocationID       *int    `json:"location_id"`         // Optional, defaults to 1
    NewLocationName  *string `json:"new_location_name"`   // Optional
    LevelID          *int    `json:"level_id"`            // Optional
    ParentLocationID *int    `json:"parent_location_id"`  // Optional
    ReplacedID       *int    `json:"replaced_device_id"`
}


func FetchDeviceCategories() ([]DeviceCategory, error) {
	log.Printf("hehehe")
    rows, err := config.DB.Query("SELECT category_id, category_name FROM Device_Category WHERE status = 1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var categories []DeviceCategory
    for rows.Next() {
        var cat DeviceCategory
        if err := rows.Scan(&cat.ID, &cat.Name); err != nil {
            return nil, err
        }
        categories = append(categories, cat)
    }
    return categories, nil
}

func FetchDeviceTypes() ([]DeviceType, error) {
	
    rows, err := config.DB.Query("SELECT type_id, type_name FROM device_type WHERE status = 1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var types []DeviceType
    for rows.Next() {
        var t DeviceType
        if err := rows.Scan(&t.ID, &t.Name); err != nil {
            return nil, err
        }
        types = append(types, t)
    }
    return types, nil
}

func FetchDeviceTypesByCategory(catID int) ([]DeviceType, error) {
    rows, err := config.DB.Query("SELECT type_id, type_name FROM Device_Type WHERE category_id = ? AND status = 1", catID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var types []DeviceType
    for rows.Next() {
        var t DeviceType
        if err := rows.Scan(&t.ID, &t.Name); err != nil {
            return nil, err
        }
        types = append(types, t)
    }
    return types, nil
}

func InsertDeviceCategory(name string) error {
    _, err := config.DB.Exec("CALL sp_add_device_category(?)", name)
    return err
}

func InsertDeviceType(catID int, typeName string) error {
    _, err := config.DB.Exec("CALL sp_add_device_type(?, ?)", catID, typeName)
    return err
}

// func InsertDevice(device DeviceInput, addedBy int) error {
//     _, err := config.DB.Exec("CALL sp_add_device(?, ?, ?, ?, ?, ?, ?, ?, ?)",
//         device.DeviceName,
//         device.TypeID,
//         device.SerialNo,
//         device.ModelNo,
//         device.PurchaseDate,
//         device.WarrantyExpiry,
//         addedBy,
//         device.LocationID,
//         device.ReplacedID,
//     )
//     return err
// }


// func InsertDevice(device DeviceInput, addedBy int) error {
//     locID := 1 // default

//     if device.LocationID != nil {
//         locID = *device.LocationID
//     }

//     _, err := config.DB.Exec("CALL sp_add_device(?, ?, ?, ?, ?, ?, ?, ?, ?)",
//         device.DeviceName,
//         device.TypeID,
//         device.SerialNo,
//         device.ModelNo,
//         device.PurchaseDate,
//         device.WarrantyExpiry,
//         addedBy,
//         locID,
//         device.ReplacedID,
//     )
//     return err
// }

func InsertDevice(device DeviceInput, addedBy int) (DeviceIDResponse, error) {
	locID := 1 // default location

	if device.LocationID != nil {
		locID = *device.LocationID
	}

	_, err := config.DB.Exec("CALL sp_add_device(?, ?, ?, ?, ?, ?, ?, ?, ?, @device_id)",
		device.DeviceName,
		device.TypeID,
		device.SerialNo,
		device.ModelNo,
		device.PurchaseDate,
		device.WarrantyExpiry,
		addedBy,
		locID,
		device.ReplacedID,
	)

	if err != nil {
		log.Printf("Error calling stored procedure: %v", err)
		return DeviceIDResponse{}, err
	}

	var id int
	err = config.DB.QueryRow("SELECT @device_id").Scan(&id)
	if err != nil {
		log.Printf("Error retrieving output parameters: %v", err)
		return DeviceIDResponse{}, err
	}

	log.Printf("Device added with ID: %d\n", id)
	return DeviceIDResponse{DeviceID: id}, nil
}
