package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Client struct {
	Id   int
	Name string
}

// Usando Map apenas para simplificar o uso de banco de dados
var db map[int]Client

func saveClient(w http.ResponseWriter, r *http.Request) {
	var c Client
	w.Header().Add("Content-Type", "application/json")

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	db[c.Id] = c
	w.WriteHeader(http.StatusCreated)
}

func readClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if ok {
		index, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		c, exists := db[index]

		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	for _, client := range db {
		if err := json.NewEncoder(w).Encode(&client); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

}

func editClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if ok {
		index, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		var c Client
		err = json.NewDecoder(r.Body).Decode(&c)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Id = index
		db[index] = c
		if err := json.NewEncoder(w).Encode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}
	w.WriteHeader(http.StatusNotFound)
}

func deleteClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if ok {
		index, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		_, exists := db[index]
		if !exists {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		delete(db, index)
		w.WriteHeader(http.StatusOK)
	}
	w.WriteHeader(http.StatusNotFound)
}

func main() {
	db = make(map[int]Client)
	r := mux.NewRouter()

	r.HandleFunc("/client", readClient).Methods("GET")
	r.HandleFunc("/client/{id}", readClient).Methods("GET")
	r.HandleFunc("/client", saveClient).Methods("POST")
	r.HandleFunc("/client/{id}", editClient).Methods("PUT")
	r.HandleFunc("/client/{id}", deleteClient).Methods("DELETE")

	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
	}

	log.Fatal(srv.ListenAndServe())
}
