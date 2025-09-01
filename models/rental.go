package models

import "time"

// Rental reprezentuje wypożyczenie pojazdu
type Rental struct {
    ID            string    `json:"id" firestore:"id"`
    UserID        string    `json:"user_id" firestore:"user_id"`
    VehicleID     string    `json:"vehicle_id" firestore:"vehicle_id"`
    StartTime     time.Time `json:"start_time" firestore:"start_time"`
    EndTime       *time.Time `json:"end_time" firestore:"end_time"` // moze byc null wiec pointer
    StartLocation Location  `json:"start_location" firestore:"start_location"`
    EndLocation   *Location `json:"end_location" firestore:"end_location"`
    Cost          float64   `json:"cost" firestore:"cost"`
    Status        string    `json:"status" firestore:"status"` // "active", "completed", "cancelled"
    Distance      float64   `json:"distance" firestore:"distance"` //jednostka: km
}

// RentalStatus - mozliwe statusy wypożyczenia
const (
    RentalStatusActive    = "active"
    RentalStatusCompleted = "completed"
    RentalStatusCancelled = "cancelled"
)

// StartRentalRequest - request rozpoczęcia wypożyczenia
type StartRentalRequest struct {
    VehicleID string `json:"vehicle_id" validate:"required"`
}

// EndRentalRequest - request zakończenia wypożyczenia
type EndRentalRequest struct {
    EndLocation Location `json:"end_location" validate:"required"`
}