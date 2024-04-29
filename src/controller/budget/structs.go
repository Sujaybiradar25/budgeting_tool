package budget

import (
	"kniru/src/controller/enums"
	"time"
)

type Budget struct {
	BudgetId               int64          `json:"budget_id"`
	Users                  []int64        `json:"users"`
	TotalBudgetForDuration float64        `json:"total_budget_for_duration"`
	Currency               enums.Currency `json:"currency"`
	Category               enums.Category `json:"category"`
	BudgetingDuration      time.Duration  `json:"budgetingDurationInMilliSec"`
	CurrentBudget          []int64        `json:"transactions_in_current_budget"`
	OldTransactions        []int64        `json:"transactions_in_old_budget"`
	CurrentSpend           float64        `json:"current_spend"`
	BudgetStartTime        time.Time      `json:"budget_start_time"`
}
