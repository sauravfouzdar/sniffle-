# sniffle-
Boilerplate to quickly build REST APIs in Go using Gin Gonic

## Technology Stack

 - Go 1.16
 - Gin Gonic
 - Docker 
 - Unit tests
 - Swagger
 - RDS - Postgres

## Project Structure

```
sniffle
|____config
|____controllers
|____middlewares
|____models
|____utils
|
|____main.go
```


## API Endpoints

List of available endpoints:

**app health routes**:\
`GET /health` - Health check endpoint\
`GET /test-email` - Test email endpoint

**user app routes**:\
`POST /login` - User Login\
`POST /register` - Register a user\
`POST /health` - Health check endpoint\
`POST /posts` - Create a post

