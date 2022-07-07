# go-api-postgres

## setup new project
- install dependences :
  - go get -u github.com/gorilla/mux
  - go get github.com/lib/pq
  - go get github.com/joho/godotenv
 - setup elephant sql :
  - https://www.elephantsql.com/
  - free plan
  - create new
  - sql : CREATE TABLE users (
    userid SERIAL PRIMARY KEY,
    name TEXT,
    age INT,
    location TEXT
);
  - copy url to .env

## install
- git clone https://github.com/suandedev/go-api-postgres.git
- cd go-api-postgres
- go run main.go

## test 
- go run main.go

## Create a new user (POST)
URL: http://localhost:8080/api/newuser
Body: raw/json

{
    "name": "gopher",
    "age":25,
    "location":"India"
}

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093219.png?raw=true">

## Get a user (GET)
URL: http://localhost:8080/api/user/1

/api/user/{id}

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093424.png?raw=true">

## Get all user (GET)
URL: http://localhost:8080/api/user

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093511.png?raw=true">

## Update a user (PUT)
URL: http://localhost:8080/api/user/1
Body: raw/json

{
    "name": "golang gopher",
    "age":24,
    "location":"Hyderabad, India"
}

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093559.png?raw=true">

## Delete a user (DELETE)
URL: http://localhost:8080/api/deleteuser/1

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093641.png?raw=true">
