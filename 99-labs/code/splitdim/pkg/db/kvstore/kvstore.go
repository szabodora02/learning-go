package kvstore

import (
	"fmt"
	"strconv"
	"time"

	clientapi "kvstore/pkg/api"
	"kvstore/pkg/client"

	"splitdim/pkg/api"
	"splitdim/pkg/clear"

	"resilient"
)

type kvstore struct {
	client.Client
}

func NewDataLayer(kvStoreAddr string) api.DataLayer {
	return &kvstore{Client: client.NewClient(kvStoreAddr)}
}

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
	balance, _ := strconv.Atoi(vv.Value)
	vv.Value = fmt.Sprintf("%d", balance+amount)
	vkv := clientapi.VersionedKeyValue{Key: user, VersionedValue: vv}
	err = db.Put(vkv)
	if err != nil {
		return err
	}
	return nil
}

func (db *kvstore) Transfer(t api.Transfer) error {
	if t.Sender == t.Receiver || t.Sender == "" || t.Receiver == "" {
		return fmt.Errorf("invalid transfer")
	}

	backoff := resilient.Backoff{
		Base:      150 * time.Millisecond,
		Cap:       2 * time.Second,
		Jitter:    0.1,
		NumTrials: 6,
	}

	// JAVÍTVA (VISSZAÁLLÍTVA): Sender egyenleg NŐ (+t.Amount), mert ő fizetett (neki tartoznak)
	senderOp := db.setBalanceForUser(t.Sender, t.Amount)
	err := resilient.WithRetry(senderOp, backoff)()
	if err != nil {
		return fmt.Errorf("transfer failed at sender: %v", err)
	}

	// JAVÍTVA (VISSZAÁLLÍTVA): Receiver egyenleg CSÖKKEN (-t.Amount), mert helyette fizettek
	receiverOp := db.setBalanceForUser(t.Receiver, -t.Amount)
	err = resilient.WithRetry(receiverOp, backoff)()
	if err == nil {
		return nil
	}

	// UNDO: Ha a Receiver művelet sikertelen, vissza kell vonni a Sender növekedését (-t.Amount)
	undoBackoff := backoff
	undoBackoff.NumTrials = 10

	undoOp := db.setBalanceForUser(t.Sender, -t.Amount)
	undoErr := resilient.WithRetry(undoOp, undoBackoff)()

	if undoErr != nil {
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

func (db *kvstore) Reset() error {
	return db.Client.Reset()
}
