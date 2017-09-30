package account_test

import (
	"testing"

	"github.com/GlynOwenHanmer/GOHMoney/account"
)

func TestAccountFieldError_Equal(t *testing.T) {
	testSets := []struct {
		errA, errB account.FieldError
		equal      bool
	}{
		{
			errA:  account.FieldError{},
			errB:  account.FieldError{},
			equal: true,
		},
		{
			errA: account.FieldError{
				account.EmptyNameError,
			},
			errB:  account.FieldError{},
			equal: false,
		},
		{
			errA: account.FieldError{
				account.EmptyNameError,
			},
			errB: account.FieldError{
				account.EmptyNameError,
			},
			equal: true,
		},
		{
			errA: account.FieldError{
				account.ZeroDateOpenedError,
				account.EmptyNameError,
			},
			errB: account.FieldError{
				account.EmptyNameError,
				account.ZeroDateOpenedError,
			},
			equal: false,
		},
		{
			errA: account.FieldError{
				account.ZeroDateOpenedError,
				account.EmptyNameError,
				account.ZeroValidDateClosedError,
			},
			errB: account.FieldError{
				account.ZeroDateOpenedError,
				account.EmptyNameError,
				account.ZeroValidDateClosedError,
			},
			equal: true,
		},
	}
	for _, testSet := range testSets {
		equalA := testSet.errA.Equal(testSet.errB)
		equalB := testSet.errB.Equal(testSet.errA)
		if equalA != equalB {
			t.Fatalf("Equal did not return same value when comparing account.FieldError to other account.FieldError the reverse way around.")
		}
		if testSet.equal != equalA {
			t.Errorf("Unexpected Equal value.\n\tExpected: %t\n\tActual  : %t", testSet.equal, equalA)
		}
	}
}
