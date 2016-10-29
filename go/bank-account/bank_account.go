package account

import (
	"sync"
)

// Account ..
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open ..
func Open(initalDeposit int64) *Account {
	if initalDeposit >= 0 {
		return &Account{balance: initalDeposit}
	}

	return nil
}

// Close ..
func (acc *Account) Close() (payout int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed || acc.balance < 0 {
		return 0, false
	}

	ok = true
	payout = acc.balance

	acc.balance = 0
	acc.closed = true

	return
}

// Balance ..
func (acc *Account) Balance() (balance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed {
		return 0, false
	}

	return acc.balance, true
}

// Deposit ..
func (acc *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.Lock()
	defer acc.Unlock()

	if acc.closed || acc.balance+amount < 0 {
		return 0, false
	}

	acc.balance += amount

	return acc.balance, true
}
