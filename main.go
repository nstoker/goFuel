package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	landing "github.com/nstoker/gofuel/web/landing/templates"
	"github.com/sirupsen/logrus"
)

func main() {
	fmt.Println("Initialising")

	godotenv.Load(".env")
	port := os.Getenv("PORT")
	if port == "" {
		logrus.Fatal("Environmental variable 'PORT' missing")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", landing.PageHandler)
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static",
			http.FileServer(http.Dir("./static"))))
	r.PathPrefix("/vendor").Handler(
		http.StripPrefix("/vendor",
			http.FileServer(http.Dir("./static/vendor"))))
	r.Use(loggingMiddleware)
	logrus.Infof("Listening on %s", port)
	logrus.Fatalln(http.ListenAndServe(":"+port, r))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logrus.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}
