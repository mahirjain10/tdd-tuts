package wallet

import (
	"errors"
	"testing"
)

func assertMessage(t *testing.T, got, want any) {
	t.Helper()
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func assertErrorMessage(t *testing.T, got, want error) {
	t.Helper()
	if !errors.Is(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestWallet(t *testing.T) {
	t.Run("Deposit amount into the wallet", func(t *testing.T) {
		w := Wallet{}
		w.Deposit(10)
		w.Deposit(10)
		got := w.Balance();
		assertMessage(t, got, Rupee(20))
	})

	t.Run("Withdraw amount from the wallet unsuccessfull", func(t *testing.T) {
		w := Wallet{BalanceAmount: 20}
		err := w.Withdraw(30)
		assertErrorMessage(t, err, ErrInsufficientFunds)
		assertMessage(t, w.BalanceAmount, Rupee(20))
	})

	t.Run("Withdraw amount within balance", func(t *testing.T) {
		w := Wallet{BalanceAmount: 20}

		err:=w.Withdraw(10)
		assertErrorMessage(t, err, nil)
		assertMessage(t, w.BalanceAmount, Rupee(10))
	})

	t.Run("Transferred money successfully", func(t *testing.T) {
		rw := Wallet{BalanceAmount: 100}
		sw := Wallet{BalanceAmount: 100}

		amountToTransfer := Rupee(100)

		err := sw.Transfer(&rw, amountToTransfer)
		assertMessage(t, err, nil)
		assertMessage(t, rw.BalanceAmount, Rupee(200))
		assertMessage(t, sw.BalanceAmount, Rupee(0))
	})
	t.Run("Transferred money unsuccessfully", func(t *testing.T) {
		rw := Wallet{BalanceAmount: 100}
		sw := Wallet{BalanceAmount: 100}

		amountToTransfer := Rupee(120)

		err := sw.Transfer(&rw, amountToTransfer)
		assertErrorMessage(t, err, ErrInsufficientFunds)
		assertMessage(t, rw.BalanceAmount, Rupee(100))
		assertMessage(t, sw.BalanceAmount, Rupee(100))
	})
}
