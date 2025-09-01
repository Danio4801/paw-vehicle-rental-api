package main

import (
    "log"
    "net/http"
    "os"
)

func main() {
    // Pobiermay port z zmiennych środowiskowych
    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    // Podstawowy router
    mux := http.NewServeMux()
    
    // Sprawdzamy status serwera
    mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`{"status":"ok"}`))
    })

    // Uruchamiamy serwer
    log.Printf("Serwer uruchomiony na porcie %s", port)
    if err := http.ListenAndServe(":"+port, mux); err != nil {
        log.Fatal(err)
    }
}