package strategy

import "fmt"

type Payment struct {
	context  *PaymentContext
	strategy PaymentStrategy
}

type PaymentContext struct {
	Name, CardID string
	Money        int
}

type PaymentStrategy interface {
	Pay(*PaymentContext)
}

func NewPayment(name, cardID string, money int, strategy PaymentStrategy) *Payment {
	return &Payment{
		context: &PaymentContext{
			Name:   name,
			CardID: cardID,
			Money:  money,
		},
		strategy: strategy,
	}
}

func (p *Payment) Pay() {
	p.strategy.Pay(p.context)
}

type Cash struct{}

func (c *Cash) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s vy cash", ctx.Money, ctx.Name)
}

type Bank struct{}

func (*Bank) Pay(ctx *PaymentContext) {
	fmt.Printf("Pay $%d to %s vy back account %s", ctx.Money, ctx.Name, ctx.CardID)
}
