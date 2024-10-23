package payment

import (
	"errors"
	"fmt"
)

const INVALID_AMOUNT string = "invalid amount"

// PaymentStrategy defines the interface that all payment methods will implement
type PaymentStrategy interface {
	Pay(amount float64) (string, error)
}

// CreditCardPayment is one of the strategies
type CreditCardPayment struct{}

func (c CreditCardPayment) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New(INVALID_AMOUNT)
	}
	return fmt.Sprintf("Paid %.2f using Credit Card", amount), nil
}

// PayPalPayment is another strategy
type PayPalPayment struct{}

func (p PayPalPayment) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New(INVALID_AMOUNT)
	}
	return fmt.Sprintf("Paid %.2f using PayPal", amount), nil
}

// CashPayment is another strategy
type CashPayment struct{}

func (c CashPayment) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New(INVALID_AMOUNT)
	}
	return fmt.Sprintf("Paid %.2f using Cash", amount), nil
}

// UPIPayment is the new strategy for UPI payments
type UPIPayment struct{}

func (u UPIPayment) Pay(amount float64) (string, error) {
	if amount <= 0 {
		return "", errors.New(INVALID_AMOUNT)
	}
	return fmt.Sprintf("Paid %.2f using UPI", amount), nil
}

// PaymentContext is responsible for holding the current strategy
type PaymentContext struct {
	strategy PaymentStrategy
}

// SetStrategy allows us to change the payment strategy
func (p *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// ExecutePayment executes the chosen payment strategy
func (p *PaymentContext) ExecutePayment(amount float64) (string, error) {
	return p.strategy.Pay(amount)
}
