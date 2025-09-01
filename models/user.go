package models

import "time"

// User oznacza użytkownika systemu
type User struct {
    ID           string    `json:"id" firestore:"id"`
    Email        string    `json:"email" firestore:"email"`
    Name         string    `json:"name" firestore:"name"`
    Phone        string    `json:"phone" firestore:"phone"`
    PasswordHash string    `json:"-" firestore:"password_hash"` // "-" oznacza że nie zwracamy w JSON
    CreatedAt    time.Time `json:"created_at" firestore:"created_at"`
    UpdatedAt    time.Time `json:"updated_at" firestore:"updated_at"`
}

// UserRegistration czyli dane przy rejestracji
type UserRegistration struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
    Name     string `json:"name" validate:"required"`
    Phone    string `json:"phone" validate:"required"`
}

// UserLogin tzn. dane przy logowaniu
type UserLogin struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
}