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
    price INT,
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
URL: http://localhost:8080/api/rice
Body: raw/json

{
    "name": "mapan 05",
    "price": 15000,
    "location":"indonesia"
}

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093219.png?raw=true](https://github.com/suandedev/go-rice-api-postgres/blob/main/Screenshot%20(335).png?raw=true)">

## Get a user (GET)
URL: http://localhost:8080/api/user/4

/api/user/{id}

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093424.png?raw=true](https://github.com/suandedev/go-rice-api-postgres/blob/main/Screenshot%20(333).png?raw=true)">

## Get all user (GET)
URL: http://localhost:8080/api/user

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093511.png?raw=true](https://github.com/suandedev/go-rice-api-postgres/blob/main/Screenshot%20(332).png?raw=true)">

## Update a user (PUT)
URL: http://localhost:8080/api/user/4
Body: raw/json

{
    "name": "mapan 03",
    "age": 160000,
    "location":"indonesia"
}

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093559.png?raw=true](https://github.com/suandedev/go-rice-api-postgres/blob/main/Screenshot%20(336).png?raw=true)">

## Delete a user (DELETE)
URL: http://localhost:8080/api/deleteuser/4

<img width="45%" src="https://github.com/suandedev/go-api-postgres/blob/main/Screenshot%202022-07-03%20093641.png?raw=true](https://github.com/suandedev/go-rice-api-postgres/blob/main/Screenshot%20(337).png?raw=true)">
