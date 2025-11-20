package kvstore

import (
	"fmt"
	"strconv"
	"time"

	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"

	"splitdim/pkg/api"
	"splitdim/pkg/clear"

	"resilient" // Ez a laborhoz szükséges új csomag
)

type kvstore struct {
	client.Client
}

// NewDataLayer creates a new database of scores.
func NewDataLayer(kvStoreAddr string) api.DataLayer {
	return &kvstore{Client: client.NewClient(kvStoreAddr)}
}

// Segédfüggvény a resilient csomaghoz (Closure)
func (db *kvstore) setBalanceForUser(user string, amount int) resilient.Closure {
	return func() error {
		return db.setBalance(user, amount)
	}
}

func (db *kvstore) setBalance(user string, amount int) error {
	vv, err := db.Get(user)
	if err != nil {
		return err
	}
	balance, _ := strconv.Atoi(vv.Value)              // Convert to integer: should be safe
	vv.Value = fmt.Sprintf("%d", balance+amount)      // Set the new balance in the value
	vkv := clientapi.VersionedKeyValue{Key: user, VersionedValue: vv} // Create a VersionedKeyValue
	err = db.Put(vkv)                                 // Put the new value in the database
	if err != nil {
		return err
	}
	return nil
}

// Transfer - A JAVÍTOTT, RESILIENT VERZIÓ
func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == t.Receiver || t.Sender == "" || t.Receiver == "" {
		return fmt.Errorf("invalid transfer")
	}

	// 1. Retry policy beállítása (6 próba, 150ms kezdés, 2s limit)
	backoff := resilient.Backoff{
		Base:      150 * time.Millisecond,
		Cap:       2 * time.Second,
		Jitter:    0.1,
		NumTrials: 6,
	}

	// 2. Sender egyenleg növelése (+Amount) retry-al
	senderOp := db.setBalanceForUser(t.Sender, t.Amount)
	err := resilient.WithRetry(senderOp, backoff)()
	if err != nil {
		// Ha ez nem sikerül, még nem történt semmi, visszatérünk a hibával.
		return fmt.Errorf("transfer failed at sender: %v", err)
	}

	// 3. Receiver egyenleg csökkentése (-Amount) retry-al
	receiverOp := db.setBalanceForUser(t.Receiver, -t.Amount)
	err = resilient.WithRetry(receiverOp, backoff)()
	if err == nil {
		// Siker! Mindkét oldal kész.
		return nil
	}

	// 4. UNDO (Kompenzálás): Ha a Receiver oldala sikertelen, vissza kell vonni a Sender-től.
	// Agresszívebb retry policy a visszavonáshoz
	undoBackoff := backoff
	undoBackoff.NumTrials = 10 

	undoOp := db.setBalanceForUser(t.Sender, -t.Amount) 
	undoErr := resilient.WithRetry(undoOp, undoBackoff)()

	if undoErr != nil {
		// 5. Inconsistent state: Ha az Undo is sikertelen
		return fmt.Errorf("CRITICAL: Inconsistent database state! Transfer failed and undo failed. Cause: %v", undoErr)
	}

	return fmt.Errorf("transfer failed, but rolled back successfully: %v", err)
}

func (db *kvstore) AccountList() ([]api.Account, error) {
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

func (db *kvstore) Clear() ([]api.Transfer, error) {
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
func (db *kvstore) Reset() error {
	return db.Client.Reset()
}
