package resilientkvstore

import (
	"fmt"
	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"
	"log"
	"strconv"
	"time"

	"resilient"

	"splitdim/pkg/api"
	"splitdim/pkg/clear"
)

type resilientkvstore struct {
	client.Client
	resilient.Backoff
}

// NewDataLayer creates a new database of scores.
func NewDataLayer(kvStoreAddr string) api.DataLayer {
	backoff := resilient.Backoff{
		Base:      150 * time.Millisecond,
		Cap:       2 * time.Second,
		Jitter:    3,
		NumTrials: 6,
	}
	return &resilientkvstore{Client: client.NewClient(kvStoreAddr), Backoff: backoff}
}

func (db *resilientkvstore) setBalance(user string, amount int) error {
	vv, err := db.Get(user)
	if err != nil {
		return err
	}
	balance, _ := strconv.Atoi(vv.Value)                              // Convert to integer: should be safe
	vv.Value = fmt.Sprintf("%d", balance+amount)                      // Set the new balance in the value
	vkv := clientapi.VersionedKeyValue{Key: user, VersionedValue: vv} // Create a VersionedKeyValue
	err = db.Put(vkv)                                                 // Put the new value in the database
	if err != nil {
		return err
	}
	return nil
}

func (db *resilientkvstore) Transfer(t api.Transfer) error {
	if t.Sender == t.Receiver || t.Sender == "" || t.Receiver == "" {
		return fmt.Errorf("invalid transfer")
	}
	err := resilient.WithRetry(func() error {
		return db.setBalance(t.Sender, t.Amount)
	}, db.Backoff)()
	if err != nil {
		return err
	}
	err = resilient.WithRetry(func() error {
		return db.setBalance(t.Receiver, -t.Amount)
	}, db.Backoff)()
	if err != nil {
		log.Println("Transfer failed, rolling back")
		// Rollback the sender's balance
		err = resilient.WithRetry(func() error {
			return db.setBalance(t.Sender, -t.Amount)
		}, resilient.Backoff{
			Base:      250 * time.Millisecond,
			Cap:       50 * time.Second,
			Jitter:    2,
			NumTrials: 10,
		})()
		if err != nil {
			return fmt.Errorf("transfer failed and rollback failed: %v", err)
		}
	}
	return nil
}

func (db *resilientkvstore) AccountList() ([]api.Account, error) {
	vkvs, err := db.List()
	if err != nil {
		return nil, err
	}
	ret := []api.Account{}
	for _, vkv := range vkvs {
		balance, _ := strconv.Atoi(vkv.Value)
		ret = append(ret, api.Account{Holder: vkv.Key, Balance: balance})
	}
	return ret, nil
}

func (db *resilientkvstore) Clear() ([]api.Transfer, error) {
	accounts, err := db.AccountList()
	if err != nil {
		return nil, err
	}
	transformedAccounts := make(map[string]int)
	for _, account := range accounts {
		transformedAccounts[account.Holder] = account.Balance
	}
	transfers, err := clear.Clear(transformedAccounts)
	return transfers, err
}

// Reset sets all balances to zero.
func (db *resilientkvstore) Reset() error {
	return db.Client.Reset()
}
