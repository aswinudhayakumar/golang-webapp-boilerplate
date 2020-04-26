# golang-boilerplate

This is a boilerplate to getting started with a golang web application

## What's Included ?

This boilerplate was developed using ```gorilla/mux``` for routing and ```jinzhu/gorm``` for database connectivity. 

    * Fully dockerized with postgres
    * Totally structured setup
    * Quick and easy production ready setup
    * An api router
    * Quick setup of database driver for postgres
    * Setup to write your authentication middleware on the go

## Configurations

To run this boilerplate you must add a ``` .env ``` file to the root directory which should have

    SERVER_PORT="port number for server to run"
    DB_USERNAME="postgres username"
    DB_PASSWORD="postgres password"
    DB_NAME="postgres database name"
    DB_HOST="postgres host" 
    DB_PORT="postgres port number" (from port [6002:5432] in docker-compose.yml - where the port 5432 should be used in .env)

you can find everything in ```Dockerfile``` and if you want you can also change all the attributes.

## How to run ?

Here are the command to getting started with this boilerplate with docker,

    docker-compose up --build

To run without docker, you can use the local database and run the app using the below commands,

    go run main.go

Or you can also run by building your go project by using the below commands,

    1) go build
    2) ./golang-boilerplate

You are ready to rock with go !
