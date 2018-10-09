package main

import (
  "fmt"
  "net/http"
  "encoding/json"

  "github.com/gorilla/mux"
)

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

  http.ListenAndServe(":3001", r)
}
