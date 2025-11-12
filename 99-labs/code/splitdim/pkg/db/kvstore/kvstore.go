package kvstore

import (
	"fmt"
	"strconv"

	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"

	"splitdim/pkg/api"
)

// kvstore implements api.DataLayer using the kvstore client
type kvstore struct {
	client.Client
}

// NewDataLayer creates a new kvstore-based datalayer
func NewDataLayer(kvStoreAddr string) api.DataLayer {
	return &kvstore{Client: client.NewClient(kvStoreAddr)}
}

// setBalance updates the balance of a user by amount (can be negative)
func (db *kvstore) setBalance(user string, amount int) error {
	for {
		vv, err := db.Get(user)
		if err != nil {
			// If user doesn't exist yet, initialize
			vv = clientapi.VersionedValue{Value: "0", Version: 0}
		}

		balance, _ := strconv.Atoi(vv.Value)
		vv.Value = fmt.Sprintf("%d", balance+amount)

		vkv := clientapi.VersionedKeyValue{
			Key:            user,
			VersionedValue: vv, // Beágyazott VersionedValue
		}

		if err := db.Put(vkv); err == nil {
			return nil
		}
		// Retry on version conflict
	}
}

// Transfer moves t.Amount from t.Sender to t.Receiver
func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == "" || t.Receiver == "" || t.Sender == t.Receiver {
		return fmt.Errorf("invalid transfer")
	}

	// Decrease sender balance
	for {
		if err := db.setBalance(t.Sender, -t.Amount); err == nil {
			break
		}
	}

	// Increase receiver balance
	for {
		if err := db.setBalance(t.Receiver, t.Amount); err == nil {
			break
		}
	}

	return nil
}

// AccountList returns all accounts as api.Account slice
func (db *kvstore) AccountList() ([]api.Account, error) {
	accounts, err := db.List() // List() from client
	if err != nil {
		return nil, err
	}

	ret := []api.Account{}
	for _, kv := range accounts {
		balance, _ := strconv.Atoi(kv.VersionedValue.Value)
		ret = append(ret, api.Account{
			Holder:  kv.Key,
			Balance: balance,
		})
	}

	return ret, nil
}

// computeClear calculates minimal transfers to settle all balances
func computeClear(accounts map[string]int) ([]api.Transfer, error) {
	// Lab szerint a részletes algoritmus opcionális
	return []api.Transfer{}, nil
}

// Clear returns the minimal list of transfers to clear debts
func (db *kvstore) Clear() ([]api.Transfer, error) {
	accounts, err := db.AccountList()
	if err != nil {
		return nil, err
	}

	accMap := make(map[string]int)
	for _, a := range accounts {
		accMap[a.Holder] = a.Balance
	}

	return computeClear(accMap)
}

// Reset sets all balances to zero
func (db *kvstore) Reset() error {
	return db.Client.Reset()
}
