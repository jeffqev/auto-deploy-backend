package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Mensaje envia un mensaje desde el servidor
type Mensaje struct {
	Msg   string `json:"msg"`
	Fecha string `json:"fecha"`
}

// FechaDB guarda la fecha de la consulta de la db
var FechaDB string

func main() {

	route := mux.NewRouter()
	route.HandleFunc("/", handleResponseFecha)
	handler := cors.Default().Handler(route)

	PORT := os.Getenv("Port")
	if PORT == "" {
		PORT = "1323"
	}

	log.Fatal(http.ListenAndServe(":1323", handler))

}

func handleResponseFecha(w http.ResponseWriter, r *http.Request) {
	fecha, err := handleFechaDB()

	if err != nil {
		mensaje := Mensaje{
			Msg:   "Error al consultar en la db:",
			Fecha: "",
		}
		w.WriteHeader(http.StatusInternalServerError)
		w.Header().Set("Content-type", "Application/json")
		json.NewEncoder(w).Encode(mensaje)
		return
	}

	mensaje := Mensaje{
		Msg:   "La fecha y hora desde la db es",
		Fecha: fecha,
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-type", "Application/json")
	json.NewEncoder(w).Encode(mensaje)
}

func handleFechaDB() (string, error) {

	USER := os.Getenv("USER_MYSQL")
	if USER == "" {
		USER = "root"
	}

	PASSWORD := os.Getenv("PASSWORD_MYSQL")
	if PASSWORD == "" {
		PASSWORD = "root"
	}

	HOST := os.Getenv("HOST_MYSQL")
	if HOST == "" {
		HOST = "localhost:1517"
	}

	// Conexion
	db, err := sql.Open("mysql", USER+":"+PASSWORD+"@tcp("+HOST+")/mysql")

	if err != nil {
		return "", err
	}
	defer db.Close()

	// Consulta
	results, err := db.Query("SELECT NOW()")

	for results.Next() {
		err = results.Scan(&FechaDB)
		if err != nil {
			return "", err
		}
	}

	if err != nil {
		return "", err
	}

	return FechaDB, nil
}
