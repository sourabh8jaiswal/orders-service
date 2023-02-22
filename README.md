# Order Microservice : RESTful API service in golang using Gin Gorm
## Introduction

A simple RESTful API microservice written in Go using the gin framework. This service exposes the APIs for the following operations related to orders :
1. Add an order
2. Fetch all orders
3. Update an order

All requests made to this service store their data in PostgreSQL.

## Pre-requisites / Service Dependency

Make sure you have the following installed:

1. PostgreSQL server, and running on local machine or a working environment
2. Go Language (1.8 or above)
3. Clone the application
4. Install Go package dependencies by running the below commands from order-service folder

    `go get -u github.com/gin-gonic/gin`
    
    `go get -u gorm.io/gorm`
    
    `go get -u gorm.io/driver/postgres`
    
    `go get github.com/thedevsaddam/govalidator`
    
    
5. Log into PostgreSQL DB, and create a `order_service` database.

## Setting up DB Migrations for Order service

**Before running migrations make sure, you specify correct host, port, dbname, user and password in initializer/database.go file

e.g. "host=localhost user=postgres dbname=order_service port=5432 sslmode=disable"

**For running migrations, run the following file: migrations/migrate.go

    go run migrations/migrate.go


## Starting the Service
1. Start Go order service locally. Run the following command from the directory in which **`main.go`** is located in order-service folder. This command will run the executable to start the service.
    
    `go run main.go`
    
2. Play with order endpoints to add/update/get order.

The endpoints can be invoked using the popular `curl` command or any REST client such as `POSTMAN`, etc.

* _Ping to check service health_
    
    
    HTTP METHOD : GET
    URL : http://localhost:5050

* _Create a new order_

    
    HTTP METHOD : POST
    
    URL : http://localhost:5050/orders
    
    Content-Type : application/json; charset=utf-8
    
    REQUEST PAYLOAD :
    
    `
    {
    "status": "PENDING_INVOICE",
    "items" : [{
        "id": "1234",
        "description": "some description",
        "price": 10,
        "quantity": 1
    },
    {
        "id": "1234",
        "description": "any description",
        "price": 29,
        "quantity": 1
    }],
    "total": 46,
    "currency_unit": "USD"
}
`
* _Fetch all existing orders_
    
    
    HTTP METHOD : GET
    
    URL : http://localhost:5050/orders


* Update an existing order
    
    
    HTTP METHOD : PUT
    
    URL : http://localhost:5050/orders/{order_id}
    
    SAMPLE URL : http://localhost:5050/orders/1
    
    Content-Type : application/json; charset=utf-8
    
    REQUEST PAYLOAD :
    
   `{
    "status": "SUCCESS",
    "items" : [{
        "id": "1234",
        "description": "sourabh",
        "price": 10,
        "quantity": 1
    }],
    "total": 98,
    "currency_unit": "USD"
}`

 
