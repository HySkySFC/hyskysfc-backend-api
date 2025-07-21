# HySkySFC

## Instalation
Need some library to running this app.
- chi
- goose
- pgx

If you don't want develope this api, you can just running the binary file.
```bash
./skyhysfc-backend-app
```
running the above command on the root of this project.
Note: need docker running before.

## Running docker
Need docker run to running database.
```bash
docker compose build up
```

## Migration
Running migration with run main.go

Note (this just a note for someone developing this api):
```bash
goose -dir migrations postgres "postgres://hyskysfc:hyskysfc@localhost:5432/hyskysfc?sslmode=disable" up
```

## API Documentation
This full documentation to access api endpoint.

### Create User
POST /users
```bash
{
    "username": "admin",
    "email": "admin@gmail.com",
    "password": "admin123"
}
```

### Create Token
POST /token/authentication
```bash
{
    "username": "admin",
    "password": "admin123"
}
```

## Authenticated Routes
All of the route below need beare token to access, we can get the token exactly on create token.

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
    "efisiensi": {
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
    "efisiensi": {
        "100%": 0.2555,
        "75%": 0.2455,
        "50%": 0.2333,
    },
    "batas_beban": 750
}
```
### Delete PLTD By ID
DELETE /pltd/:id

