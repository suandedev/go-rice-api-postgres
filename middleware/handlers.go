package middleware

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-rice-api-postgres/models"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// response format
type response struct {
	ID int64 `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// create connection with postgres db 
func createConnection() *sql.DB {
	// load .env file
	err := godotenv.Load(".env")
	
	if err != nil {
		log.Fatalf("Eddor loading .env file")
	}

	// open the connecction
	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	// check the connection
	err = db.Ping()

	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")
	// retun the db connection
	return db
}

// createRice create a rice in the postgres db
func CreateRice(w http.ResponseWriter, r *http.Request) {
	// set the header to connect type x-www-form-urlencoded
	// allow all origin to handle cors issue
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// create an empty rice of type models.Rice
	var rice models.Rice

	// decode the json request to rice
	err := json.NewDecoder(r.Body).Decode(&rice)

	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}

	// cal insert rice functions and pass the rice
	insertID := insertRice(rice)

	// format a response object
	res := response {
		ID : insertID,
		Message : "Rice created successfully",
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

// getAllRice get all rice from the db
func GetAllRice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// get all the rice in the db
	rice, err := getAllRice()

	if err != nil {
		log.Fatalf("unable to get all rice. %v", err)
	}

	// send all the rice as response
	json.NewEncoder(w).Encode(rice)
}

func getAllRice() ([]models.Rice, error) {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	var rices []models.Rice
	
	// create sql statement sql query
	sqlStatement := `SELECT * FROM rice`

	// execute the sql query
	rows, err := db.Query(sqlStatement)

	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)
	}

	// close the statement
	defer rows.Close()

	// iterate over the rows
	for rows.Next() {
		var rice models.Rice

		// unmarshal the row object to rice
		err = rows.Scan(&rice.ID, &rice.Name, &rice.Price, &rice.Location)

		if err != nil {
			log.Fatalf("unable to scan the row object. %v", err)
		}

		// append the rice in the slice
		rices = append(rices, rice)
	}

	// return empty rice on error
	return rices, err
}

func GetRice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// get the riceid from the reqquest param, key is id
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert the id to int. %v", err)
	}

	// call the get user function with id to retrieve a single rice
	rice, err := getRice(int64(id))

	if err != nil {
		log.Fatalf("unable to get the rice. %v", err)
	}

	// send the response
	json.NewEncoder(w).Encode(rice)
}

func UpdateRice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the rice id from the request params , key is id
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert the id to int. %v", err)
	}

	// create an empty rice of type models.Rice
	var rice models.Rice

	// decode the json request to rice
	err = json.NewDecoder(r.Body).Decode(&rice)

	if err != nil {
		log.Fatalf("unable to decode the request body. %v", err)
	}

	// call update rice to update the rice
	updateRows := updateRice(int64(id), rice)

	// format the message string
	msg := fmt.Sprintf("Rice updated successfully %v", updateRows)

	// format the response message
	res := response {
		ID : int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)

}

func DeleteRice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	// get the rice id from the request params , key is id
	params := mux.Vars(r)

	// convert the id type from string to int
	id, err := strconv.Atoi(params["id"])

	if err != nil {
		log.Fatalf("unable to convert the id to int. %v", err)
	}

	// call delete rice to delete the rice
	deleteRows := deleteRice(int64(id))

	// format the message string
	msg := fmt.Sprintf("Rice deleted successfully %v", deleteRows)

	// format the response message
	res := response {
		ID : int64(id),
		Message: msg,
	}

	// send the response
	json.NewEncoder(w).Encode(res)
}

func deleteRice(id int64) int64 {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create sql statement sql query
	sqlStatement := `DELETE FROM rice WHERE riceid=$1`

	// execute the sql query
	rows, err := db.Exec(sqlStatement, id)

	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)
	}

	// get the number of rows affected
	rowsAffected, err := rows.RowsAffected()

	if err != nil {
		log.Fatalf("unable to get the rows affected. %v", err)
	}

	// return the number of rows affected
	return rowsAffected
}

func updateRice(id int64, rice models.Rice) int64 {

	//  create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the sql statement
	sqlStatement := `UPDATE rice SET name=$2, price=$3, location=$4 WHERE riceid=$1`

	// execute the sql query
	res, err := db.Exec(sqlStatement, id, rice.Name, rice.Price, rice.Location)

	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)
	}

	// get the number of rows affected
	rowsAffected, err := res.RowsAffected()

	if err != nil {
		log.Fatalf("unable to get the rows affected. %v", err)
	}

	fmt.Printf("total rows/record affected %v", rowsAffected)

	return rowsAffected
}

// getRice get a rice from the db
func getRice(id int64) (models.Rice, error) {
	// create the postgres connection
	db := createConnection()

	// close connection
	defer db.Close()

	// create rice models.Rice type
	var rice models.Rice

	// create the sql select sql query
	sqlStatement := `SELECT * FROM rice WHERE riceid = $1`

	// execute the sql query
	row := db.QueryRow(sqlStatement, id)

	// unmarchal the row object to rice
	err := row.Scan(&rice.ID, &rice.Name, &rice.Price, &rice.Location)

	switch err {
		case sql.ErrNoRows:
			fmt.Println("no rows were returned")
			return rice, nil
		case nil:
			return rice, nil
		default:
			log.Fatalf("unable to scan the row. %v", err)
	}

	// return empty rice on error
	return rice, err
}

// handler functions
// insert rice in the db
func insertRice(rice models.Rice) int64 {
	// create the postgres db connection
	db := createConnection()

	// close the db connection
	defer db.Close()

	// create the query
	// returning rice is will return the id of the inserted rice
	sqlStatement := `INSERT INTO rice (name, price, location) VALUES ($1, $2, $3) RETURNING riceid`

	// the inserted id will store in this id
	var id int64

	// execute the sql statement
	// scan function will save the insert id in the id
	err := db.QueryRow(sqlStatement, rice.Name, rice.Price, rice.Location).Scan(&id)

	if err != nil {
		log.Fatalf("unable to execute the query. %v", err)
	}

	fmt.Printf("inserted a singgle record %v", id)

	// return the inserted id
	return id
}