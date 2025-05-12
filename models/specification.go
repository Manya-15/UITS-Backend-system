package models

import (
    "webtechproject/config"
    "log"
)

type SpecificationTemplate struct {
    ID      int    `json:"template_id"`
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

type TemplateIDResponse struct {
    SpecTemplateID int `json:"spec_template_id"`
}

type SpecificationIDResponse struct {
    SpecMasterID int `json:"spec_master_id"`
}


// func FetchTemplatesByDeviceID(deviceID int) ([]SpecificationTemplate, error) {
//     query := `
//         SELECT st.spec_template_id, st.type_id, st.spec_name 
//         FROM Specification_Template st 
//         JOIN Device d ON d.type_id = st.type_id 
//         WHERE d.device_id = ? AND st.status = 1`
//     rows, err := config.DB.Query(query, deviceID)
//     if err != nil {
//         return nil, err
//     }
//     defer rows.Close()

//     var templates []SpecificationTemplate
//     for rows.Next() {
//         var t SpecificationTemplate
//         rows.Scan(&t.ID, &t.TypeID, &t.Name)
//         templates = append(templates, t)
//     }
//     return templates, nil
// }

func FetchTemplatesByTypeID(TypeID int) ([]SpecificationTemplate, error) {
    query := `
        SELECT st.spec_template_id, st.spec_name 
        FROM Specification_Template st 
        JOIN device_type d ON d.type_id = st.type_id 
        WHERE d.type_id = ? AND st.status = 1`
    rows, err := config.DB.Query(query, TypeID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var templates []SpecificationTemplate
    for rows.Next() {
        var t SpecificationTemplate
        rows.Scan(&t.ID, &t.Name)
        templates = append(templates, t)
    }
    return templates, nil
}

func InsertSpecificationTemplate(typeID int, name string) (TemplateIDResponse, error) {
    _, err := config.DB.Exec("CALL sp_add_specification_template(?, ?, @template_id)", typeID, name)
    if err != nil {
        log.Printf("Error executing stored procedure: %v", err)
        return TemplateIDResponse{}, err
    }

    var id int
    err = config.DB.QueryRow("SELECT @template_id").Scan(&id)
    if err != nil {
        log.Printf("Error fetching template ID: %v", err)
        return TemplateIDResponse{}, err
    }

    log.Printf("Specification template added with ID: %d", id)
    return TemplateIDResponse{SpecTemplateID: id}, nil
}


func FetchSpecificationValues(TemplateID int) ([]SpecificationMaster, error) {
    log.Println("Fetching specification values for template ID:", TemplateID)
    query := `
        SELECT st.spec_master_id, st.spec_value 
        FROM Specification_master st 
        JOIN Specification_template s ON st.spec_template_id = s.spec_template_id
        WHERE s.spec_template_id = ? AND st.status = 1 AND s.status = 1`
    rows, err := config.DB.Query(query, TemplateID)

    log.Println(rows, err)
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


func InsertSpecificationValue(value string, templateID int) (SpecificationIDResponse, error) {
    _, err := config.DB.Exec("CALL sp_add_specification_master(?, ?, @spec_master_id)", value, templateID)
    if err != nil {
        log.Printf("Error executing stored procedure: %v", err)
        return SpecificationIDResponse{}, err
    }

    var id int
    err = config.DB.QueryRow("SELECT @spec_master_id").Scan(&id)
    if err != nil {
        log.Printf("Error retrieving specification ID: %v", err)
        return SpecificationIDResponse{}, err
    }

    log.Printf("Specification value added with ID: %d", id)
    return SpecificationIDResponse{SpecMasterID: id}, nil
}


// func InsertDeviceSpecifications(specs []DeviceSpecificationInput) error {
//     tx, err := config.DB.Begin()
//     if err != nil {
//         return err
//     }
//     stmt, err := tx.Prepare(`
//         INSERT INTO Device_Specification (device_id, spec_template_id, spec_master_id) 
//         VALUES (?, ?, ?)`)
//     if err != nil {
//         tx.Rollback()
//         return err
//     }
//     defer stmt.Close()

//     for _, spec := range specs {
//         _, err := stmt.Exec(spec.DeviceID, spec.SpecTemplateID, spec.SpecMasterID)
//         if err != nil {
//             tx.Rollback()
//             return err
//         }
//     }
//     return tx.Commit()
// }
func InsertDeviceSpecifications(specs []DeviceSpecificationInput) error {
    tx, err := config.DB.Begin()
    if err != nil {
        return err
    }

    for _, spec := range specs {
        _, err := tx.Exec("CALL sp_add_device_specification(?, ?, ?)",
            spec.DeviceID, spec.SpecTemplateID, spec.SpecMasterID)
        if err != nil {
            tx.Rollback()
            return err
        }
    }

    return tx.Commit()
}
