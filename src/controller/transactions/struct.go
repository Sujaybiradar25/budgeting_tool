package transactions

import (
	"kniru/src/controller/enums"
	"time"
)

type Transactions struct {
	TransId         int64
	UserId          int64
	Amount          float64
	Currency        enums.Currency
	TransactionTime time.Time
	Category        enums.Category
}
