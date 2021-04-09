package account

import s "sync"

type Account interface {

	// Open(initialDeposit int64) *Account
	Close() (payout int64, ok bool)
	Balance() (balance int64, ok bool)
	Deposit(amount int64) (newBalance int64, ok bool)
}

type AccountImpl struct {
	closed bool
	amount int64
	mtx    s.Mutex
}

func Open(initialDeposit int64) Account {
	if initialDeposit < 0 {
		return nil
	}
	return &AccountImpl{closed: false, amount: initialDeposit}
}
func (acc *AccountImpl) Close() (payout int64, ok bool) {
	acc.mtx.Lock()
	defer acc.mtx.Unlock()
	if acc.closed {
		ok = false
		payout = 0
		return
	}
	acc.closed = true
	payout = acc.amount
	acc.amount = 0
	ok = true
	return
}
func (acc AccountImpl) Balance() (balance int64, ok bool) {
	if acc.closed {
		ok = false
		balance = 0
		return
	}
	balance = acc.amount
	ok = true
	return
}
func (acc *AccountImpl) Deposit(amount int64) (newBalance int64, ok bool) {
	acc.mtx.Lock()
	defer acc.mtx.Unlock()
	if acc.closed {
		ok = false
		newBalance = 0
		return
	}
  if acc.amount+amount < 0 {
    ok = false
    newBalance = acc.amount
  } else {
    ok = true
    acc.amount += amount
    newBalance = acc.amount
  }
  return
}

