package main

import "fmt"

type paymentMethod interface {
	processPayment(totalAmount float64) string
	calculateFee(amount float64) float64
}

type creditCard struct {
	cardNumber    string
	feePercentage float64
}

func (c *creditCard) processPayment(totalAmount float64) string {
	// Your code goes here
	return ""
}

func (c *creditCard) calculateFee(amount float64) float64 {
	// Your code goes here
	return 0
}

type paypal struct {
	email   string
	flatFee float64
}

func (p *paypal) processPayment(totalAmount float64) string {
	// Your code goes here
	return ""
}

func (p *paypal) calculateFee(amount float64) float64 {
	// Your code goes here
	return 0
}

type bankTransfer struct {
	accountNumber string
}

func (b *bankTransfer) processPayment(totalAmount float64) string {
	return fmt.Sprintf("Processed bank transfer of $%.2f", totalAmount)
}

func (b *bankTransfer) calculateFee(amount float64) float64 {
	return 0.0
}

func completePayment(p paymentMethod, amount float64) {
	// Your code goes here
}

func main() {
	cc := &creditCard{cardNumber: "1234 5678 9012 3456", feePercentage: 0.03}
	pp := &paypal{email: "johndoe@example.com", flatFee: 0.5}
	bt := &bankTransfer{accountNumber: "1234567890"}

	completePayment(cc, 100.0)
	completePayment(pp, 100.0)
	completePayment(bt, 100.0)
}
