package models

import "time"

// Vehicle reprezentuje pojazd (hulajnoga lub rower)
type Vehicle struct {
    ID           string    `json:"id" firestore:"id"`
    Type         string    `json:"type" firestore:"type"` // "scooter" lub "bike"
    QRCode       string    `json:"qr_code" firestore:"qr_code"`
    BatteryLevel int       `json:"battery_level" firestore:"battery_level"` // 0-100
    Location     Location  `json:"location" firestore:"location"`
    Status       string    `json:"status" firestore:"status"` // "available", "rented", "maintenance"
    LastUpdate   time.Time `json:"last_update" firestore:"last_update"`
}

// Location reprezentuje współrzędne GPS
type Location struct {
    Lat float64 `json:"lat" firestore:"lat"`
    Lng float64 `json:"lng" firestore:"lng"`
}

// VehicleStatus - możliwe statusy pojazdu
const (
    VehicleStatusAvailable   = "available"
    VehicleStatusRented      = "rented"
    VehicleStatusMaintenance = "maintenance"
)

// VehicleType - typy pojazdów
const (
    VehicleTypeScooter = "scooter"
    VehicleTypeBike    = "bike"
)