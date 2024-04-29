package utils

import "kniru/src/controller/enums"

var BCat = make(map[Budgeter]int64)

type Budgeter struct {
	Id       int64
	Category enums.Category
}

func GetBudgetByUserIdCategory(id int64, category enums.Category) int64 {
	t := Budgeter{
		Id:       id,
		Category: category,
	}
	return BCat[t]
}
