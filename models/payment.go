package models

import "time"

// Payment - płatność za wypożyczenie
type Payment struct {
    ID            string    `json:"id" firestore:"id"`
    RentalID      string    `json:"rental_id" firestore:"rental_id"`
    UserID        string    `json:"user_id" firestore:"user_id"`
    Amount        float64   `json:"amount" firestore:"amount"`
    Currency      string    `json:"currency" firestore:"currency"` // "PLN"
    PaymentMethod string    `json:"payment_method" firestore:"payment_method"` // "card", "blik", "paypal"
    Status        string    `json:"status" firestore:"status"` // "pending", "completed", "failed"
    CreatedAt     time.Time `json:"created_at" firestore:"created_at"`
}

// PaymentStatus - dostępne statusy płatności
const (
    PaymentStatusPending   = "pending"
    PaymentStatusCompleted = "completed"
    PaymentStatusFailed    = "failed"
)

// PaymentMethod - metody płatności
const (
    PaymentMethodCard   = "Bank card"
    PaymentMethodApplePay   = "Apple Pay"
    PaymentMethodBlik   = "Blik"
    PaymentMethodPayPal = "Paypal"
)