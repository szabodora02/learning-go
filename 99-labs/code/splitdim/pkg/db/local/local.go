package local


import (
    "sort"
    "sync"
    "fmt"
    "splitdim/pkg/api"
)
// localDB is a simple implementation of the DataLayer API.
type localDB struct {
    // accounts maintains the balance for each user name
    accounts map[string]int
    // The read-write mutex makes sure concurrent access is safe.
    mu sync.RWMutex
}


func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

// NewDataLayer creates a new database of accounts.
func NewDataLayer() api.DataLayer {
    return &localDB{accounts: make(map[string]int)}
}

func (db *localDB) Transfer(t api.Transfer) error {
    if t.Sender == t.Receiver {
        return fmt.Errorf("sender and receiver must be different")
    }

    db.mu.Lock()
    defer db.mu.Unlock()

    if _, ok := db.accounts[t.Sender]; !ok {
        db.accounts[t.Sender] = 0
    }
    if _, ok := db.accounts[t.Receiver]; !ok {
        db.accounts[t.Receiver] = 0
    }

    db.accounts[t.Sender] += t.Amount
    db.accounts[t.Receiver] -= t.Amount

    return nil
}
func (db *localDB) AccountList() ([]api.Account, error) {
    db.mu.RLock()
    defer db.mu.RUnlock()

    ret := []api.Account{}
    for name, balance := range db.accounts {
        ret = append(ret, api.Account{
            Holder:  name,
            Balance: balance,
        })
    }

    sort.Slice(ret, func(i, j int) bool {
        return ret[i].Holder < ret[j].Holder
    })

    return ret, nil
}
func (db *localDB) Clear() ([]api.Transfer, error) {
    db.mu.RLock()

    // 1. Ellenőrizzük az egyenlegek konzisztenciáját
    sum := 0
    for _, balance := range db.accounts {
        sum += balance
    }
    if sum != 0 {
        db.mu.RUnlock()
        return nil, fmt.Errorf("inconsistent balances: total sum is not zero")
    }

    // 2. Másoljuk el az adatokat egy ideiglenes map-be
    tempAcc := make(map[string]int)
    for user, balance := range db.accounts {
        tempAcc[user] = balance
    }

    db.mu.RUnlock()

    // 3. Üres slice a tranzakcióknak
    transfers := []api.Transfer{}

    // 4. Addig futtatjuk, amíg vannak pozitív egyenlegű felhasználók
    for {
        var debtor string
        var debtorBal int
        foundDebtor := false

        for user, bal := range tempAcc {
            if bal < 0 {
                debtor = user
                debtorBal = bal
                foundDebtor = true
                break
            }
        }

        if !foundDebtor {
            break // nincs több adós
        }

        for creditor, creditorBal := range tempAcc {
            if creditorBal <= 0 || creditor == debtor {
                continue
            }

            transferAmount := min(-debtorBal, creditorBal)

            transfers = append(transfers, api.Transfer{
                Sender:   debtor,
                Receiver: creditor,
                Amount:   transferAmount,
            })

            tempAcc[debtor] += transferAmount
            tempAcc[creditor] -= transferAmount

            if tempAcc[debtor] == 0 {
                break
            }
        }
    }

    return transfers, nil
}
func (db *localDB) Reset() error {
    db.mu.Lock()
    defer db.mu.Unlock()

    db.accounts = make(map[string]int)

    return nil
}
// SetBalance updates the balance of a user by amount (can be negative)
func (db *localDB) SetBalance(user string, amount int) error {
    db.mu.Lock()
    defer db.mu.Unlock()

    if _, ok := db.accounts[user]; !ok {
        db.accounts[user] = 0
    }
    db.accounts[user] += amount
    return nil
}
