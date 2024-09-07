package main

import (
	"context"
	"fmt"
	"justcgh9/spotify_clone/server/config"
	"justcgh9/spotify_clone/server/repositories"
	"justcgh9/spotify_clone/server/routers"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	clientOptions := options.Client().ApplyURI(config.MongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	repositories.Initialize(client)

	r := mux.NewRouter()

	r.HandleFunc("/registration", routers.HandleCORS).Methods("OPTIONS")
	r.HandleFunc("/registration", routers.PostSignUp).Methods("POST")
	r.HandleFunc("/users", routers.GetUsers).Methods("GET")
	r.HandleFunc("/user/{user_id}", routers.GetUser).Methods("GET")
	r.HandleFunc("/refresh", routers.GetRefreshedToken).Methods("POST")
	r.HandleFunc("/activate/{link}", routers.GetActivation).Methods("GET")
	r.HandleFunc("/logout", routers.PostSignOut).Methods("POST")
	r.HandleFunc("/login", routers.HandleCORS).Methods("OPTIONS")
	r.HandleFunc("/login", routers.PostSignIn).Methods("POST")

	r.HandleFunc("/tracks/upload", routers.PostTrack).Methods("POST")
	r.HandleFunc("/tracks/search", routers.SearchTrack).Methods("GET")
	r.HandleFunc("/tracks", routers.GetTracksHandler).Methods("GET")
	r.HandleFunc("/tracks/{track_id}", routers.GetTrackHandler).Methods("GET")
	r.HandleFunc("/tracks/{track_id}", routers.DeleteTrack).Methods("DELETE")
	r.HandleFunc("/tracks/{track_id}/comment", routers.HandleCORS).Methods("OPTIONS")
	r.HandleFunc("/tracks/{track_id}/comment", routers.CreateComment).Methods("POST")
	r.HandleFunc("/tracks/{track_id}/comment", routers.GetComments).Methods("GET")
	r.HandleFunc("/tracks/{track_id}/comment/{comment_id}", routers.EditComment).Methods("PUT")
	r.HandleFunc("/tracks/{track_id}/comment/{comment_id}", routers.DeleteComment).Methods("DELETE")

	r.HandleFunc("/albums", routers.PostAlbum).Methods("POST")
	r.HandleFunc("/albums/{album_id}", routers.PostToAlbum).Methods("POST")
	r.HandleFunc("/albums/{album_id}", routers.GetAlbum).Methods("GET")
	r.HandleFunc("/albums/{album_id}", routers.DeleteAlbum).Methods("DELETE")

	staticDir := "./files/"
	fs := http.FileServer(http.Dir(staticDir))
	r.PathPrefix("/files/").Handler(http.StripPrefix("/files/", fs))

	server := http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Starting server on :8080")
	err = server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}
