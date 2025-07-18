# HySkySFC

## Instalation
Need some library to running this app.
- chi
- goose
- pgx

## Running docker
Need docker run to running database.
```bash
docker compose build up
```

## Migration
Running migration with run main.go

## API Documentation
This full documentation to access api endpoint.

### Get All PLTD
GET /pltd

### Get PLTD By ID
GET /pltd/:id

### Create PLTD
POST /pltd
body:
```json
{
    "name": "Mesin 1",
    "status": "tersedia",
    "efisitensi": {
        "100%": 0.2555,
        "75%": 0.2455,
        "50%": 0.2333,
    },
    "batas_beban": 750
}
```
### Update PLTD By ID
PUT /pltd/:id
body
```json
{
    "name": "Mesin 1",
    "status": "tersedia",
    "efisitensi": {
        "100%": 0.2555,
        "75%": 0.2455,
        "50%": 0.2333,
    },
    "batas_beban": 750
}
```
### Delete PLTD By ID
DELETE /pltd/:id

