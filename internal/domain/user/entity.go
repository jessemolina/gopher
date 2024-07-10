package user

import (
    "time"

    "github.com/google/uuid"
)

// User represents an individual who interacts with the machine learning system.
type User struct {
    ID          uuid.UUID
    Name        string
    Email       string
    Roles       []Role
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

type Role string

const (
    RoleAdmin Role = "admin"
    RoleUser  Role = "user"
)
