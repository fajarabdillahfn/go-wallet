# go-wallet
wallet app with kafka

# Pre-Requisite
* Docker
* Go > 17.0

# Initialization
* `export APP_PORT=9000` or any port number you want
* `docker-compose up -d`

# Start The Application
`go run main.go`

# API Contract

## Insert a New Deposit
`POST /wallets`

### Request Data (application/json)
example:
```
{
    "id": 1, //required, integer
    "amount": 2000 //required, float
}
```

## Get Wallet Details
`GET /wallets/:wallet_id` with `wallet_id` is an integer
