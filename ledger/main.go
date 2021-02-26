package main

import (
    "log"
    "time"
    "net/http"
    "github.com/gorilla/mux"
)

func main() {
    router := mux.NewRouter()
    router.Use(middleware)
    router.PathPrefix("/did/").Handler(http.StripPrefix("/did/", http.FileServer(http.Dir("./did"))))

    server := &http.Server{
        Handler: router,
        Addr: "127.0.0.1:3000",
        WriteTimeout: 15 * time.Second,
        ReadTimeout: 15 * time.Second,
    }

    log.Fatal(server.ListenAndServe()) 
}

func middleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Add("Content-Type", "application/did+ld+json")
        next.ServeHTTP(w, r)
    })
}
