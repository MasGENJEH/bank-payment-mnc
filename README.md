# Prerequisites

- Go (Golang) is installed on your system.
- PostgreSQL is installed, and you have created the tables as specified in the `ddl.sql` file. Then, insert the table contents from the `dml.sql` file as dummy data.
- An active internet connection is required to download Go dependencies.

## API Spec

### Login API

Request :

- Method : `POST`
- Endpoint : `/login`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Body :

```json
{
  "username": "string",
  "password": "string"
}
```

### Create Payment

Request :

- Method : POST
- Endpoint : `/payment`
- Header :
  - Content-Type : application/json
  - Accept : application/json
- Authorization : Bearer Token
- Body :

```json
{
    "customers_id": "string",
    "merchant_id": "string",
    "amount": float64,
    "transaction_type": "string",
    "description": "string"
}
```

Response :

- Status : 201 Created
- Body :

```json
{
    "status": {
        "code": 201,
        "message": "Created"
    },
    "data": {
        "id": "TRX-string",
        "customers_id": "string",
        "merchant_id": "string",
        "date": "2000-01-01T12:00:00Z", (curent time)
        "amount": float64,
        "transaction_type": "string",
        "current_balance": float64,
        "description": "string",
        "created_at": "2000-01-01T12:00:00Z", (curent time)
        "updated_at": "2000-01-01T12:00:00Z" (curent time)
    }
}
```

### Logout API

Request :

- Method : `POST`
- Endpoint : `/logout`
- Header :

  - Content-Type : application/json
  - Accept : application/json
  - Authorization : Bearer Token

- Response :

```json
{
  "status": {
    "code": 201,
    "message": "Logout successful"
  },
  "data": null
}
```
