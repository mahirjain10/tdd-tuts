package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds")

type Rupee int64

func (r Rupee) String() string {
	return fmt.Sprintf("%d Rupees", int64(r))
}

type Wallet struct {
	BalanceAmount Rupee
}

func (wallet *Wallet) Balance() Rupee {
	return wallet.BalanceAmount
}

func (wallet *Wallet) Deposit(amount Rupee) {
	wallet.BalanceAmount += amount
}

func (wallet *Wallet) Withdraw(amount Rupee) error {
	if wallet.BalanceAmount < amount {
		return fmt.Errorf("%w: cannot withdraw more than %d", ErrInsufficientFunds, wallet.BalanceAmount)
	}
	wallet.BalanceAmount -= amount
	return nil
}

func (wallet *Wallet) Transfer(to *Wallet, amount Rupee) error {
	if amount > wallet.BalanceAmount {
		return fmt.Errorf("%w: cannot transfer than more than %d", ErrInsufficientFunds, wallet.BalanceAmount)
	}
	wallet.BalanceAmount -= amount
	to.BalanceAmount += amount
	return nil
}
