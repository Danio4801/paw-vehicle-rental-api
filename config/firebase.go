package config

import (
    "context"
    "log"
    
    firebase "firebase.google.com/go"
    "cloud.google.com/go/firestore"
    "google.golang.org/api/option"
)

// FirebaseApp przechowuje instancje Firebase
type FirebaseApp struct {
    App       *firebase.App
    Firestore *firestore.Client
}

// InitFirebase inicjalizuje połączenie z Firebase
func InitFirebase(projectID string) *FirebaseApp {
    ctx := context.Background()
    
    var app *firebase.App
    var err error

    // Używamy emulatora w dev env
    if getEnv("FIRESTORE_EMULATOR_HOST", "") != "" {
        log.Println("Używam emulatora Firestore")
        conf := &firebase.Config{ProjectID: projectID}
        app, err = firebase.NewApp(ctx, conf)
    } else {
        // W prod service account
        opt := option.WithCredentialsFile("firebase-key.json")
        app, err = firebase.NewApp(ctx, nil, opt)
    }
    
    if err != nil {
        log.Fatalf("Błąd inicjalizacji Firebase: %v", err)
    }
    
    // Inicjalizacja klienta Firestore
    client, err := app.Firestore(ctx)
    if err != nil {
        log.Fatalf("Błąd inicjalizacji Firestore: %v", err)
    }
    
    return &FirebaseApp{
        App:       app,
        Firestore: client,
    }
}