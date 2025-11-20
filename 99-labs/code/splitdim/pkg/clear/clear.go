package clear

import (
	"errors"
	"splitdim/pkg/api"
)

// Clear clears the debts for the accounts given as argument. Meanwhile, it updates "accounts", so always pass a copy to this function.
func Clear(accounts map[string]int) ([]api.Transfer, error) {
	total := 0
	for _, balance := range accounts {
		total += balance
	}
	if total != 0 {
		return nil, errors.New("total balance is not zero")
	}
	tempAccounts := make(map[string]int)
	for name, balance := range accounts {
		tempAccounts[name] = balance
	}
	transfers := []api.Transfer{}
	for sender, balance := range tempAccounts {
		if balance < 0 {
			for receiver, receiverBalance := range tempAccounts {
				if receiverBalance > 0 {
					amount := min(-balance, receiverBalance)
					transfers = append(transfers, api.Transfer{Sender: sender, Receiver: receiver, Amount: amount})
					balance += amount
					tempAccounts[receiver] -= amount
				}
				if tempAccounts[sender] == 0 {
					break
				}
			}
		}
	}
	return transfers, nil
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
