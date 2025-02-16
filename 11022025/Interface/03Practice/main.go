/*
you need to create an online payment system with two types of payment processors:
1.Basic Professor - can only process payments.
2.Advanced Processor - can process payments, issue refunds ,and show transactionn history.

TODO:
Create an interface PaymentProcessor with a method:

	-ProcessPayment(amount float64)

Crate another interface

	AdvancedPaymentProcessor that embeds PaymentProcessor and adds two more methods:
		-Refund(amount float64)
		-TransactionHistory()

Create two structs:

	-BasicProcessor - Implements only ProcessPayment().
	-PremiumProcessor-Implements all three methods.

In the main() function , create instances of both processors and demonstrate their usage.
*/
package main

import "fmt"

type PaymentProcessor interface {
	ProcessPayment(amount float64)
}
type AdvancedPaymentProcessor interface{
	PaymentProcessor  
	Refund(amount float64)
	TransactionHistory()
}

type BasicProcessor struct{}


func (bp BasicProcessor)ProcessPayment(amount float64){
	fmt.Printf("Basic Processor :Process Payment of amount:%.2f\n",amount)
}

type PremiumProcessor struct{}

func (pp PremiumProcessor)ProcessPayment(amount float64){
	fmt.Printf("PremiumProcessor :Process Payment of amount:%.2f\n",amount)
}
func (pp PremiumProcessor)Refund(amount float64){
	fmt.Printf("Refunded Amount:%.2f\n",amount)
}
func (pp PremiumProcessor)TransactionHistory(){
	fmt.Println("Showing transaction history ...")
}

func main() {
	var Basic PaymentProcessor=BasicProcessor{}
	Basic.ProcessPayment(4225.222)
	var Premium AdvancedPaymentProcessor = PremiumProcessor{}
 
	Premium.ProcessPayment(1222.00)
	Premium.Refund(122.1223)
	Premium.TransactionHistory()
}
