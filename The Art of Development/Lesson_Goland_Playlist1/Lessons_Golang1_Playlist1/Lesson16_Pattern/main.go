package main

func main() {
	product := "vahicle"
	payWay := 3
	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("12345", "123")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}
	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	// ... implementation
	err := payment.Pay() // .varCheckError
	if err != nil {
		return
	}
}

type Payment interface { // interface не в этом месте по своей архитектуре, где он используется там и должен быть.
	Pay() error
}

type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment{
		cardNumber: cardNumber,
		cvv:        cvv,
	}
}

func (p *cardPayment) Pay() error {
	// ... implementation
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}
func (p *payPalPayment) Pay() error {
	// ... implementation
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}
func (p *qiwiPayment) Pay() error {
	// ... implementation
	return nil
}
