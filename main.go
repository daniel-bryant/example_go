package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)

const port = ":3001"

type Person struct {
  Name string
}

func main() {
  r := mux.NewRouter()
  r.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
    var people = []Person {
      Person {"Joe"},
      Person {"Jane"},
      Person {"John"},
    }

    bytes, err := json.Marshal(people)

    if err != nil {
      fmt.Fprintf(w, `{"Error":"Error encoding json"}`)
      return
    }

    w.Header().Set("Access-Control-Allow-Origin", "*")
    fmt.Fprintf(w, string(bytes))
  })

  fmt.Printf("* Listening on tcp://0.0.0.0%s\n", port)
  fmt.Print("Use Ctrl-C to stop\n")
  http.ListenAndServe(port, r)
}
