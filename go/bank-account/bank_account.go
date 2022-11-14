package account

import "sync"

// Account represents bank account, which is concurrency-safe.
type Account struct {
	active  bool
	balance int64
	mux     sync.Mutex
}

func Open(amount int64) *Account {
	if amount < 0 {
		return nil
	}
	return &Account{
		balance: amount,
		active:  true,
	}
}

func (a *Account) Balance() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.active {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.balance+amount < 0 || !a.active {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if !a.active {
		return 0, false
	}
	a.active = false
	return a.balance, true
}
