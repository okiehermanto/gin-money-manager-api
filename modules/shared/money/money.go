package money

import "fmt"

type Money struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

func (m Money) Add(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf(
			"cannot add %s to %s",
			other.Currency,
			m.Currency,
		)
	}

	return Money{
		Amount:   m.Amount + other.Amount,
		Currency: m.Currency,
	}, nil
}

func (m Money) Subtract(other Money) (Money, error) {
	if m.Currency != other.Currency {
		return Money{}, fmt.Errorf(
			"cannot subtract %s from %s",
			other.Currency,
			m.Currency,
		)
	}

	return Money{
		Amount:   m.Amount - other.Amount,
		Currency: m.Currency,
	}, nil
}
