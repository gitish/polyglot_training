package main

import (
	"bufio"
	"fmt"
	"go-test/payment"
	"os"
	"strconv"
	"strings"
)

func RunPayments() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter payment amount: ")
	amountStr, _ := reader.ReadString('\n')
	amountStr = strings.TrimSpace(amountStr)
	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount")
		return
	}

	fmt.Print("Enter payment method (credit, paypal, cash, upi): ")
	method, _ := reader.ReadString('\n')
	method = strings.TrimSpace(strings.ToLower(method))

	paymentContext := &payment.PaymentContext{}

	switch method {
	case "credit":
		paymentContext.SetStrategy(payment.CreditCardPayment{})
	case "paypal":
		paymentContext.SetStrategy(payment.PayPalPayment{})
	case "cash":
		paymentContext.SetStrategy(payment.CashPayment{})
	case "upi":
		paymentContext.SetStrategy(payment.UPIPayment{})
	default:
		fmt.Println("Invalid payment method")
		return
	}

	result, err := paymentContext.ExecutePayment(amount)
	if err != nil {
		fmt.Println("payment.ERROR", err)
	} else {
		fmt.Println(result)
	}
}

func main() {
	RunPayments()
}
