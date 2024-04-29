package budget

import (
	"kniru/src/controller/application/utils"
	"kniru/src/controller/enums"
	"time"
)

var Bdgt = make(map[int64]*Budget)
var BudgetCount int64

//var BCat map[Budgeter]*Budget

//type Budgeter struct {
//	Id       int64
//	Category enums.Category
//}
//
//func GetBudgetByUserIdCategory(id int64, category enums.Category) *Budget {
//	t := Budgeter{
//		Id:       id,
//		Category: category,
//	}
//	return BCat[t]
//}

func GetBudgetById(id int64) *Budget {
	//fmt.Println(Bdgt)
	return Bdgt[id]
}

func GetAllBudgets() []*Budget {
	b := make([]*Budget, 0)
	for _, v := range Bdgt {
		b = append(b, v)
	}
	return b
}
func CreateBudget(users []int64, category enums.Category, budgetForDuration float64, currency enums.Currency, budgetingDuration time.Duration) error {
	b := Budget{
		BudgetId:               BudgetCount,
		Users:                  nil,
		TotalBudgetForDuration: budgetForDuration,
		Currency:               currency,
		Category:               category,
		BudgetingDuration:      budgetingDuration,
		CurrentBudget:          nil,
		OldTransactions:        nil,
		CurrentSpend:           0,
		BudgetStartTime:        time.Now(),
	}
	b.BudgetStartTime = time.Now()
	b.Users = make([]int64, 0)
	for _, u := range users {
		b.Users = append(b.Users, u)
	}
	b.CurrentBudget = make([]int64, 0)
	b.OldTransactions = make([]int64, 0)
	bt := utils.Budgeter{
		Id:       b.BudgetId,
		Category: category,
	}
	utils.BCat[bt] = b.BudgetId
	//bCat[bt] = &b
	Bdgt[b.BudgetId] = &b
	BudgetCount++
	return nil
}
