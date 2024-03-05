package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Issues reading from the env file ,", err.Error())
	}

	router := mux.NewRouter() //initializing the router

	//handling the endpoints
	router.HandleFunc("/upload", UploadFileHandler).Methods("POST")
	router.HandleFunc("/download/{file-name}", RetrieveFileHandler).Methods("GET")
	router.HandleFunc("/list", GetAllFilesHandler).Methods("GET")
	router.HandleFunc("/download/{file-name}", DeleteFileHandler).Methods("DELETE")

	//getting the required values from env files
	address := os.Getenv("APP_Addr")
	port := os.Getenv("APP_Port")

	//starting the server

	fmt.Println("Starting the server on the port ", port)
	srv := &http.Server{
		Handler: router,
		Addr:    fmt.Sprintf("%s%s", address, port),
	}

	err = srv.ListenAndServe()

	if err != nil {
		log.Fatal("Issue with starting the server at port ", port, " ", err.Error())
	}
}
