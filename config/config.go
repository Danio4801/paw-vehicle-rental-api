package config

import (
    "log"
    "os"
    "time"
)

// Config przechowuje konfigurację aplikacji
type Config struct {
    Port             string
    JWTSecret        string
    JWTExpiration    time.Duration
    FirebaseProject  string
    Environment      string // "development", "production"
    PricePerMinute   float64
    UnlockPrice      float64
}

// LoadConfig wczytuje konfigurację ze zmiennych środowiskowych
func LoadConfig() *Config {
    config := &Config{
        Port:            getEnv("PORT", "8080"),
        JWTSecret:       getEnv("JWT_SECRET", "key"),
        FirebaseProject: getEnv("FIREBASE_PROJECT", ""),
        Environment:     getEnv("ENVIRONMENT", "development"),
        PricePerMinute:  1.20,
        UnlockPrice:     1.50,
    }

    // Czas ważności tokena JWT 
    config.JWTExpiration = 24 * time.Hour

    if config.Environment == "production" && config.JWTSecret == "key" {
        log.Fatal("JWT_SECRET set w prod")
    }

    return config
}

// getEnv pobiera zmienną środowiskową lub zwraca wartość domyślną
func getEnv(key, defaultValue string) string {
    if value := os.Getenv(key); value != "" {
        return value
    }
    return defaultValue
}