package domain

import (
	"time"
	uuid "github.com/satori/go.uuid"
)

type CreditCard struct {
	ID string
	Name string
	Number string
	EspirationMonth int32
	EspirationYear int32
	CVV int32
	Balance float64
	Limit float64 
	CreatedAt time.time
}

func NewCreditCard() CreditCard {
	c := CreditCard()
	c.ID = uuid.NewV4().String()
	c.CreatedAt = time.now()
	return c
}