package main

import (
    "encoding/json"
    "log"
    "net/http"
    "splitdim/pkg/api"
    "splitdim/pkg/db/local"
)

var db api.DataLayer

// TransferHandler is a HTTP handler that implements the money transfer API.
func TransferHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    var t api.Transfer
    if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
        http.Error(w, "Invalid JSON", http.StatusBadRequest)
        return
    }

    log.Printf("Transfer requested: %+v", t)

    if err := db.Transfer(t); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    w.WriteHeader(http.StatusOK)
}
// AccountListHandler is a HTTP handler that returns the current balance of each registered user.
func AccountListHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    log.Println("Account list requested")

    accountList, err := db.AccountList()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(accountList)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// ClearHandler is a HTTP handler that returns a list of transfers to clear the balance of each user.
func ClearHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    log.Println("Clear requested")

    transfers, err := db.Clear()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    jsonData, err := json.Marshal(transfers)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write(jsonData)
}

// ResetHandler is a HTTP handler that allows to zero out all balances.
func ResetHandler(w http.ResponseWriter, r *http.Request) {
 if r.Method != http.MethodGet {
        w.WriteHeader(http.StatusMethodNotAllowed)
        return
    }

    if err := db.Reset(); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
}

func main() {
    // Set the default logger to a fancier log format.
    log.SetFlags(log.LstdFlags | log.Lshortfile)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "static/index.html")
})
   http.HandleFunc("/api/transfer", TransferHandler)
   http.HandleFunc("/api/accounts", AccountListHandler)
   http.HandleFunc("/api/clear", ClearHandler)
   http.HandleFunc("/api/reset", ResetHandler)

   db = local.NewDataLayer()

   log.Println("Server listening on http://:8080")
   log.Fatal(http.ListenAndServe(":8080", nil)) 
}
