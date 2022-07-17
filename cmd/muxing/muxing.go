package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/name/{PARAM}", getParamFunc).Methods(http.MethodGet)
	router.HandleFunc("/bad", getBadFunc).Methods(http.MethodGet)
	router.HandleFunc("/data", postParamFunc).Methods(http.MethodPost)
	router.HandleFunc("/headers", postHeaderFunc).Methods(http.MethodPost)
	router.HandleFunc("/", anyNotDefined)

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

func getParamFunc(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, "+params["PARAM"]+"!")
}

func getBadFunc(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusInternalServerError)
}

func postParamFunc(w http.ResponseWriter, r *http.Request) {
	bodyData, _ := io.ReadAll(r.Body)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "I got message:\n"+string(bodyData))
}

func postHeaderFunc(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.Header.Get("a"))
	b, _ := strconv.Atoi(r.Header.Get("b"))
	w.Header().Add("a+b", strconv.Itoa(a+b))
	w.WriteHeader(http.StatusOK)
}

func anyNotDefined(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
