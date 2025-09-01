# Dokumentacja API

## Endpointy

### Health Check
- **GET** `/health`
- Zwraca status serwera
- Odpowiedź: `{"status":"ok"}`

### Autoryzacja
- **POST** `/auth/register` - Rejestracja użytkownika
- **POST** `/auth/login` - Logowanie użytkownika
- **POST** `/auth/refresh` - Odświeżenie tokena

### Pojazdy
- **GET** `/vehicles` - Lista dostępnych pojazdów
- **GET** `/vehicles/:id` - Szczegóły pojazdu
- **PUT** `/vehicles/:id/location` - Aktualizacja lokalizacji

### Wypożyczenia
- **POST** `/rentals/start` - Rozpoczęcie wypożyczenia
- **POST** `/rentals/:id/end` - Zakończenie wypożyczenia
- **GET** `/rentals/my` - Historia wypożyczeń użytkownika

### Admin
- **POST** `/admin/vehicles` - Dodanie pojazdu
- **GET** `/admin/stats` - Statystyki systemu