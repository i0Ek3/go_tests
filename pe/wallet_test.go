package pe

import (
	"fmt"
	"testing"
)

func TestWallet(t *testing.T) {

	// Use assert functions replace assistant anonymous functions
	/*
		assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
			got := w.Balance()
			if got != want {
				t.Errorf("got '%s' want '%s'", got, want)
			}
		}

		assertError := func(t *testing.T, got error, want error) {
			if got == nil {
				t.Fatal("we need an error")
			}

			if got != want {
				t.Errorf("got '%s' want '%s'", got, want)
			}
		}
	*/

	w := Wallet{}

	t.Run("Depoist", func(t *testing.T) {
		w.Deposit(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, w, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		err := w.Withdraw(Bitcoin(10))
		want := Bitcoin(0)

		assertBalance(t, w, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw with unenough balance", func(t *testing.T) {
		initBalance := Bitcoin(10)
		w = Wallet{initBalance}
		err := w.Withdraw(Bitcoin(100))

		assertBalance(t, w, initBalance)
		assertError(t, err, InsufficientBalanceErr)
	})
}

func assertBalance(t *testing.T, w Wallet, want Bitcoin) {
	got := w.Balance()
	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but we don't need it")
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("we need an error")
	}

	if got != want {
		t.Errorf("got '%s' want '%s'", got, want)
	}
}

func ExampleDeposit() {
	w := Wallet{}
	w.Deposit(Bitcoin(10))
	ret := w.Balance()
	fmt.Println(ret)
	// Output: 10 BTC
}

func ExampleWithdraw() {
	w := Wallet{}
	w.Withdraw(Bitcoin(0))
	ret := w.Balance()
	fmt.Println(ret)
	// Output: 0 BTC
}
