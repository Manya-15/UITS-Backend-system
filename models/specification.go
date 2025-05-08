package models

import "webtechproject/config"

type SpecificationTemplate struct {
    ID      int    `json:"id"`
    TypeID  int    `json:"type_id"`
    Name    string `json:"name"`
}

type SpecificationMaster struct {
    ID    int    `json:"id"`
    Value string `json:"value"`
}

type DeviceSpecificationInput struct {
    DeviceID        int `json:"device_id"`
    SpecTemplateID  int `json:"spec_template_id"`
    SpecMasterID    int `json:"spec_master_id"`
}

func FetchTemplatesByDeviceID(deviceID int) ([]SpecificationTemplate, error) {
    query := `
        SELECT st.spec_template_id, st.type_id, st.spec_name 
        FROM Specification_Template st 
        JOIN Device d ON d.type_id = st.type_id 
        WHERE d.device_id = ? AND st.status = 1`
    rows, err := config.DB.Query(query, deviceID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var templates []SpecificationTemplate
    for rows.Next() {
        var t SpecificationTemplate
        rows.Scan(&t.ID, &t.TypeID, &t.Name)
        templates = append(templates, t)
    }
    return templates, nil
}

func InsertSpecificationTemplate(typeID int, name string) error {
    _, err := config.DB.Exec("INSERT INTO Specification_Template (type_id, spec_name) VALUES (?, ?)", typeID, name)
    return err
}

func FetchSpecificationValues() ([]SpecificationMaster, error) {
    rows, err := config.DB.Query("SELECT spec_master_id, spec_value FROM Specification_Master WHERE status = 1")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var values []SpecificationMaster
    for rows.Next() {
        var v SpecificationMaster
        rows.Scan(&v.ID, &v.Value)
        values = append(values, v)
    }
    return values, nil
}

func InsertSpecificationValue(value string) error {
    _, err := config.DB.Exec("INSERT INTO Specification_Master (spec_value) VALUES (?)", value)
    return err
}

func InsertDeviceSpecifications(specs []DeviceSpecificationInput) error {
    tx, err := config.DB.Begin()
    if err != nil {
        return err
    }
    stmt, err := tx.Prepare(`
        INSERT INTO Device_Specification (device_id, spec_template_id, spec_master_id) 
        VALUES (?, ?, ?)`)
    if err != nil {
        tx.Rollback()
        return err
    }
    defer stmt.Close()

    for _, spec := range specs {
        _, err := stmt.Exec(spec.DeviceID, spec.SpecTemplateID, spec.SpecMasterID)
        if err != nil {
            tx.Rollback()
            return err
        }
    }
    return tx.Commit()
}