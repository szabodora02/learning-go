package main

import (
	"context"
	"encoding/json"
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"splitdim/pkg/api"
	"splitdim/pkg/db/kvstore"
	"splitdim/pkg/db/local"
	"splitdim/pkg/db/resilientkvstore"
)

// KVStoreMode defines the data layer mode (local/redis/kvstore).
var KVStoreMode string

// KVStoreAddr stores the key-value store address as a DNS domain name or IP address.
var KVStoreAddr string

var db api.DataLayer

// TransferHandler is a HTTP handler that implements the money transfer API.
func TransferHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	transfer := api.Transfer{}
	err := json.NewDecoder(r.Body).Decode(&transfer)
	if err != nil {
		// Return HTTP 400
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	log.Println("Transfer:", transfer)
	err = db.Transfer(transfer)
	if err != nil {
		log.Println("Error transferring money:", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// AccountListHandler is a HTTP handler that returns the current balance of each registered user.
func AccountListHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Listing accounts")
	accounts, err := db.AccountList()
	if err != nil {
		log.Println("Error listing accounts:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(accounts)
	if err != nil {
		log.Println("Error encoding accounts:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Accounts: %v", accounts)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// ClearHandler is a HTTP handler that returns a list of transfers to clear the balance of each user.
func ClearHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Clearing accounts")
	transfers, err := db.Clear()
	if err != nil {
		log.Println("Error clearing accounts:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(transfers)
	if err != nil {
		log.Println("Error encoding transfers:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("Transfers: %v", transfers)
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(json)
}

// ResetHandler is a HTTP handler that allows to zero out all balances.
func ResetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		// Return HTTP 405
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	log.Println("Resetting all balances")
	err := db.Reset()
	if err != nil {
		log.Println("Error resetting balances:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func main() {
	// Set the default logger to a fancier log format.
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Determine defaults based on ENV vars or hardcoded fallbacks
	defaultMode := os.Getenv("KVSTORE_MODE")
	if defaultMode == "" {
		defaultMode = "local"
	}

	defaultAddr := os.Getenv("KVSTORE_ADDR")
	if defaultAddr == "" {
		defaultAddr = "localhost:8081"
	}

	// Define flags using the calculated defaults
	flag.StringVar(&KVStoreMode, "mode", defaultMode, "Data layer mode (local/kvstore/resilientkvstore)")
	flag.StringVar(&KVStoreAddr, "addr", defaultAddr, "Key-value store address")

	// IMPORTANT: Parse the flags!
	flag.Parse()

	switch KVStoreMode {
	case "kvstore":
		log.Printf("Using the kvstore datalayer using at %q", KVStoreAddr)
		db = kvstore.NewDataLayer(KVStoreAddr)
	case "resilientkvstore":
		log.Printf("Using the resilientkvstore datalayer using at %q", KVStoreAddr)
		db = resilientkvstore.NewDataLayer(KVStoreAddr)
	case "local":
		fallthrough
	default:
		log.Println("Using the local datalayer")
		db = local.NewDataLayer()
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "static/index.html")
	})

	http.HandleFunc("/api/transfer", TransferHandler)
	http.HandleFunc("/api/accounts", AccountListHandler)
	http.HandleFunc("/api/clear", ClearHandler)
	http.HandleFunc("/api/reset", ResetHandler)
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	log.Println("Server listening on http://:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	s := &http.Server{Addr: ":8080"}
	go func() {
		if err := s.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("HTTP server error: %v", err)
		}
	}()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Minute)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Shutdown error: %s", err)
	}
}
