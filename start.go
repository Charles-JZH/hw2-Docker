package main

import (
        "fmt"
        "log"
        "net/http"
        "os"
)

func main() {
        log.Print("starting server...")
        log.Print("Hello World!")
        http.HandleFunc("/", handler)

        // Determine port for HTTP service.
        port := os.Getenv("PORT")
        if port == "" {
            log.Fatal("PORT is empty")
        }

        // Start HTTP server.
        log.Printf("listening on port %s", port)
        if err := http.ListenAndServe(":" + port, nil); err != nil {
            log.Fatal(err)
        }
}

func handler(w http.ResponseWriter, r *http.Request) {
        log.Print("Hello World!")
        fmt.Fprintf(w, "Hello World!\n")
}
