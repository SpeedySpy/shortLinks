/***
Omma Habiba BIPLOB
M2 Master Manager Dev
Grp 1
Projet Golang : Elections 2022
***/

package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type URL struct {
	ID           string    `json:"id" bson:"id"`
	LongUrl      string    `json:"original_url" bson:"long_url"`
	ShortUrl     string    `json:"short_url" bson:"short_url"`
	ExpirationAt time.Time `json:"expiration_at" bson:"expiration_at"`
}

const (
	mongoURI       = "mongodb://localhost:27017/"
	dbName         = "url_shortener"
	collectionName = "Urls"
)

var client *mongo.Client
var collection *mongo.Collection

func init() {
	var err error
	client, err = mongo.Connect(context.Background(), options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(dbName).Collection(collectionName)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/shorten", shortenUrl).Methods("POST")
	r.HandleFunc("/get-long-url", redirectToLongURL).Methods("POST")
	r.HandleFunc("/{shortUrl}", redirectHandler).Methods("GET")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func shortenUrl(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	longURL := data["original_url"]

	uuidShortID := uuid.New().String()
	shortID := uuidShortID[:8]

	expirationTime := time.Now().Add(24 * time.Hour)

	shortURL := "http://localhost:8080/" + shortID

	_, err = collection.InsertOne(context.Background(), URL{
		ID:           shortID,
		LongUrl:      longURL,
		ShortUrl:     shortURL,
		ExpirationAt: expirationTime,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"short_url": shortURL})
}

func redirectToLongURL(w http.ResponseWriter, r *http.Request) {
	var requestData map[string]string
	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	shortURL := requestData["short_url"]

	var url URL
	err = collection.FindOne(context.Background(), bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	if url.ExpirationAt.Before(time.Now()) {
		http.Error(w, "URL expired", http.StatusGone)
		return
	}

	http.Redirect(w, r, url.LongUrl, http.StatusFound)
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := r.URL.Path[1:]

	var url URL
	err := collection.FindOne(context.Background(), bson.M{"short_url": shortURL}).Decode(&url)
	if err != nil {
		http.Error(w, "URL not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.LongUrl, http.StatusSeeOther)
}
