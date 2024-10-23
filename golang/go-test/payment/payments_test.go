package payment_test

//normal test using testify and assert
// import (
// 	"go-test/payment"
// 	"testing"
// )

// const (
// 	NO_ERROR_MSG     = "Expected no error, got %v"
// 	EXPECTED_MSG     = "Expected %v, got %v"
// 	EXPECTED_ERR_MSG = "Expected error %v, got %v"
// )

// func TestCreditCardPayment(t *testing.T) {
// 	cc := payment.CreditCardPayment{}

// 	// Test valid amount
// 	result, err := cc.Pay(100.50)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected := "Paid 100.50 using Credit Card"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test invalid amount
// 	_, err = cc.Pay(-10)
// 	if err == nil || err.Error() != payment.INVALID_AMOUNT {
// 		t.Errorf(EXPECTED_ERR_MSG, payment.INVALID_AMOUNT, err)
// 	}
// }

// func TestPayPalPayment(t *testing.T) {
// 	pp := payment.PayPalPayment{}

// 	// Test valid amount
// 	result, err := pp.Pay(200.75)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected := "Paid 200.75 using PayPal"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test invalid amount
// 	_, err = pp.Pay(-20)
// 	if err == nil || err.Error() != payment.INVALID_AMOUNT {
// 		t.Errorf(EXPECTED_ERR_MSG, payment.INVALID_AMOUNT, err)
// 	}
// }

// func TestCashPayment(t *testing.T) {
// 	cp := payment.CashPayment{}

// 	// Test valid amount
// 	result, err := cp.Pay(50.00)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected := "Paid 50.00 using Cash"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test invalid amount
// 	_, err = cp.Pay(-5)
// 	if err == nil || err.Error() != payment.INVALID_AMOUNT {
// 		t.Errorf(EXPECTED_ERR_MSG, payment.INVALID_AMOUNT, err)
// 	}
// }

// func TestUPIPayment(t *testing.T) {
// 	upi := payment.UPIPayment{}

// 	// Test valid amount
// 	result, err := upi.Pay(150.25)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected := "Paid 150.25 using UPI"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test invalid amount
// 	_, err = upi.Pay(-15)
// 	if err == nil || err.Error() != payment.INVALID_AMOUNT {
// 		t.Errorf(EXPECTED_ERR_MSG, payment.INVALID_AMOUNT, err)
// 	}
// }

// func TestPaymentContext(t *testing.T) {
// 	pc := &payment.PaymentContext{}

// 	// Test CreditCardPayment strategy
// 	pc.SetStrategy(payment.CreditCardPayment{})
// 	result, err := pc.ExecutePayment(100.50)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected := "Paid 100.50 using Credit Card"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test PayPalPayment strategy
// 	pc.SetStrategy(payment.PayPalPayment{})
// 	result, err = pc.ExecutePayment(200.75)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected = "Paid 200.75 using PayPal"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test CashPayment strategy
// 	pc.SetStrategy(payment.CashPayment{})
// 	result, err = pc.ExecutePayment(50.00)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected = "Paid 50.00 using Cash"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}

// 	// Test UPIPayment strategy
// 	pc.SetStrategy(payment.UPIPayment{})
// 	result, err = pc.ExecutePayment(150.25)
// 	if err != nil {
// 		t.Errorf(NO_ERROR_MSG, err)
// 	}
// 	expected = "Paid 150.25 using UPI"
// 	if result != expected {
// 		t.Errorf(EXPECTED_MSG, expected, result)
// 	}
// }
