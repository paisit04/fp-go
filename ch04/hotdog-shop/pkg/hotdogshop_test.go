package pkg

import (
	"errors"
	"testing"
)

func TestCharge(t *testing.T) {
	tests := []struct {
		name       string
		inputCard  CreditCard
		amount     int
		outputCard CreditCard
		err        CreditError
	}{
		{
			"more credit",
			CreditCard{1000},
			500,
			CreditCard{500},
			nil,
		},
		{
			"exact credit",
			CreditCard{20},
			20,
			CreditCard{0},
			nil,
		},
		{
			"not enough credit",
			CreditCard{150},
			1000,
			CreditCard{150},   // no money is withdrawn
			NOT_ENOUGH_CREDIT, // payment fails with this error
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			output, err := Charge(tt.inputCard, tt.amount)
			if output != tt.outputCard {
				t.Errorf("expected %v, got %v", tt.outputCard, output)
			}
			if !errors.Is(err, tt.err) {
				t.Errorf("expected %v, got %v", tt.err, err)
			}
		})
	}
}

func TestOrderHotdog(t *testing.T) {
	testCC := CreditCard{1000}
	calledInnerFunction := false
	mockPayment := func(card CreditCard, amount int) (CreditCard, CreditError) {
		calledInnerFunction = true
		testCC.credit -= amount
		return testCC, nil
	}

	hotdog, resultF := OrderHotdog(testCC, mockPayment)
	if hotdog != NewHotdog() {
		t.Errorf("expected %v, got %v", NewHotdog(), hotdog)
	}

	_, err := resultF()
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
	if calledInnerFunction == false {
		t.Errorf("expected inner function to be called")
	}
}
