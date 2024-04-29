package transactions

import (
	"kniru/src/controller/application/utils"
	"kniru/src/controller/budget"
	"kniru/src/controller/enums"
	"time"
)

var Trans = make(map[int64]*Transactions)
var TransCount int64

func AddTransaction(user int64, amount float64, currency enums.Currency, category enums.Category, transTime time.Time) (*Transactions, error) {
	t := Transactions{
		TransId:         TransCount,
		UserId:          user,
		Amount:          amount,
		Currency:        currency,
		TransactionTime: transTime,
		Category:        category,
	}

	Trans[t.TransId] = &t
	TransCount++
	return &t, nil
}

func ProcessTransactions(t *Transactions) error {
	id := utils.GetBudgetByUserIdCategory((*t).UserId, (*t).Category)
	b := budget.GetBudgetById(id)
	if b != nil { // not required to process if it is not part of budget
		b.CurrentBudget = append(b.CurrentBudget, t.TransId)
		b.CurrentSpend += (*t).Amount
	}

	return nil
}
