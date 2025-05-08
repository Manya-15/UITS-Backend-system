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
    Name string `json:"type_name"`
}

type DeviceInput struct {
    DeviceName     string `json:"device_name"`
    TypeID         int    `json:"type_id"`
    SerialNo       string `json:"serial_no"`
    ModelNo        string `json:"model_no"`
    PurchaseDate   string `json:"purchase_date"`
    WarrantyExpiry string `json:"warranty_expiry"`
    LocationID     int    `json:"location_id"`
    ReplacedID     *int   `json:"replaced_device_id"`
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

func InsertDevice(device DeviceInput, addedBy int) error {
    _, err := config.DB.Exec("CALL sp_add_device(?, ?, ?, ?, ?, ?, ?, ?, ?)",
        device.DeviceName,
        device.TypeID,
        device.SerialNo,
        device.ModelNo,
        device.PurchaseDate,
        device.WarrantyExpiry,
        addedBy,
        device.LocationID,
        device.ReplacedID,
    )
    return err
}
