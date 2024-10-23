// Table-Driven Tests
package payment_test

import (
	"errors"
	"go-test/payment"
	"testing"
)

const (
	NO_ERROR_MSG     = "Expected no error, got %v"
	EXPECTED_MSG     = "Expected %v, got %v"
	EXPECTED_ERR_MSG = "Expected error %v, got %v"
)

func TestCreditCardPayment(t *testing.T) {
	cc := payment.CreditCardPayment{}
	tests := []struct {
		amount   float64
		expected string
		err      error
	}{
		{100.50, "Paid 100.50 using Credit Card", nil},
		{0, "", errors.New(payment.INVALID_AMOUNT)},
		{-10, "", errors.New(payment.INVALID_AMOUNT)},
	}

	for _, tt := range tests {
		result, err := cc.Pay(tt.amount)
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf(EXPECTED_ERR_MSG, tt.err, err)
		}
		if result != tt.expected {
			t.Errorf(EXPECTED_MSG, tt.expected, result)
		}
	}
}

func TestPayPalPayment(t *testing.T) {
	pp := payment.PayPalPayment{}
	tests := []struct {
		amount   float64
		expected string
		err      error
	}{
		{200.75, "Paid 200.75 using PayPal", nil},
		{0, "", errors.New(payment.INVALID_AMOUNT)},
		{-20, "", errors.New(payment.INVALID_AMOUNT)},
	}

	for _, tt := range tests {
		result, err := pp.Pay(tt.amount)
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf(EXPECTED_ERR_MSG, tt.err, err)
		}
		if result != tt.expected {
			t.Errorf(EXPECTED_MSG, tt.expected, result)
		}
	}
}

func TestCashPayment(t *testing.T) {
	cp := payment.CashPayment{}
	tests := []struct {
		amount   float64
		expected string
		err      error
	}{
		{50.00, "Paid 50.00 using Cash", nil},
		{0, "", errors.New(payment.INVALID_AMOUNT)},
		{-5, "", errors.New(payment.INVALID_AMOUNT)},
	}

	for _, tt := range tests {
		result, err := cp.Pay(tt.amount)
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf(EXPECTED_ERR_MSG, tt.err, err)
		}
		if result != tt.expected {
			t.Errorf(EXPECTED_MSG, tt.expected, result)
		}
	}
}

func TestUPIPayment(t *testing.T) {
	upi := payment.UPIPayment{}
	tests := []struct {
		amount   float64
		expected string
		err      error
	}{
		{150.25, "Paid 150.25 using UPI", nil},
		{0, "", errors.New(payment.INVALID_AMOUNT)},
		{-15, "", errors.New(payment.INVALID_AMOUNT)},
	}

	for _, tt := range tests {
		result, err := upi.Pay(tt.amount)
		if err != nil && err.Error() != tt.err.Error() {
			t.Errorf(EXPECTED_ERR_MSG, tt.err, err)
		}
		if result != tt.expected {
			t.Errorf(EXPECTED_MSG, tt.expected, result)
		}
	}
}

func TestPaymentContext(t *testing.T) {
	tests := []struct {
		name     string
		strategy payment.PaymentStrategy
		amount   float64
		expected string
		err      error
	}{
		{"CreditCardPayment", payment.CreditCardPayment{}, 100.50, "Paid 100.50 using Credit Card", nil},
		{"PayPalPayment", payment.PayPalPayment{}, 200.75, "Paid 200.75 using PayPal", nil},
		{"CashPayment", payment.CashPayment{}, 50.00, "Paid 50.00 using Cash", nil},
		{"UPIPayment", payment.UPIPayment{}, 150.25, "Paid 150.25 using UPI", nil},

		{"InvalidCreditCardPayment", payment.CreditCardPayment{}, -100.50, "", errors.New(payment.INVALID_AMOUNT)},
		{"InvalidPayPalPayment", payment.PayPalPayment{}, -200.75, "", errors.New(payment.INVALID_AMOUNT)},
		{"InvalidCashPayment", payment.CashPayment{}, -50.00, "", errors.New(payment.INVALID_AMOUNT)},
		{"InvalidUPIPayment", payment.UPIPayment{}, -150.25, "", errors.New(payment.INVALID_AMOUNT)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pc := &payment.PaymentContext{}
			pc.SetStrategy(tt.strategy)
			result, err := pc.ExecutePayment(tt.amount)
			if err != nil && err.Error() != tt.err.Error() {
				t.Errorf(EXPECTED_ERR_MSG, tt.err, err)
			}
			if result != tt.expected {
				t.Errorf(EXPECTED_MSG, tt.expected, result)
			}
		})
	}
}
