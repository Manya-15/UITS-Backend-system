package models

import (
    "webtechproject/config"
)

type User struct {
    ID          uint
    Email       string
    Password    string
    Role        string
    FullName    string
    Designation string
    Username    string
    Status      int
}

func CreateUser(user *User) error {
    query := `
        INSERT INTO user (email, password, role, full_name, designation, username, status)
        VALUES (?, ?, ?, ?, ?, ?, ?)`

    result, err := config.DB.Exec(query,
        user.Email, user.Password, user.Role,
        user.FullName, user.Designation, user.Username, user.Status)
    if err != nil {
        return err
    }

    insertID, err := result.LastInsertId()
    if err != nil {
        return err
    }

    user.ID = uint(insertID)
    return nil
}

func GetUserByEmail(email string) (*User, error) {
    query := `
        SELECT user_id, email, password, role, full_name, designation, username, status
        FROM user WHERE email = ?`

    row := config.DB.QueryRow(query, email)

    var user User
    err := row.Scan(
        &user.ID, &user.Email, &user.Password, &user.Role,
        &user.FullName, &user.Designation, &user.Username, &user.Status)
    if err != nil {
        return nil, err
    }

    return &user, nil
}

func GetUserByID(id uint) (*User, error) {
    query := `
        SELECT user_id, email, role, full_name, designation, username, status
        FROM user WHERE user_id = ?`

    row := config.DB.QueryRow(query, id)

    var user User
    err := row.Scan(
        &user.ID, &user.Email, &user.Role,
        &user.FullName, &user.Designation, &user.Username, &user.Status)
    if err != nil {
        return nil, err
    }

    return &user, nil
}
