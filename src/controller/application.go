package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"kniru/src/controller/budget"
	"kniru/src/controller/enums"
	"kniru/src/controller/transactions"
	"kniru/src/controller/users"
	"strconv"
	"time"
)

type User struct {
	Name string `json:"name"`
}

type Budget struct {
	Users    []int64 `json:"users"`
	Duration int64   `json:"duration_in_days"`
	Category string  `json:"category"`
	Budget   float64 `json:"budget"`
	Currency string  `json:"currency"`
}

type Trans struct {
	TransId         int64   `json:"trans_id"`
	UserId          int64   `json:"user_id"`
	Amount          float64 `json:"amount"`
	Currency        string  `json:"currency"`
	TransactionTime string  `json:"transaction_time"`
	Category        string  `json:"category"`
}
type BudgetId struct {
	Id int64 `json:"id"`
}

func CreateBudget(c *fiber.Ctx) error {
	data := Budget{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var cat enums.Category
	if data.Category == "Shopping" {
		cat = enums.Shopping
	} else if data.Category == "Food" {
		cat = enums.Food
	} else {
		cat = enums.Misc
	}
	var cur enums.Currency
	if data.Currency == "USD" {
		cur = enums.USD
	} else {
		cur = enums.INR
	}
	err := budget.CreateBudget(data.Users, cat, data.Budget, cur, time.Duration(data.Duration)*time.Hour*24)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("OK")
}

func GetBudget(c *fiber.Ctx) error {
	stringId := c.Params("id", "-1")
	if stringId == "all" {
		stringId = "-1"
	}
	id, err := strconv.ParseInt(stringId, 10, 64)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	if id == -1 {
		b := budget.GetAllBudgets()
		bJson, err := json.Marshal(b)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		c.Set(fiber.HeaderContentType, "application/json")
		return c.Status(fiber.StatusOK).Send(bJson)
	}

	b := budget.GetBudgetById(id)
	bJson, err := json.Marshal(b)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	c.Set(fiber.HeaderContentType, "application/json")
	return c.Status(fiber.StatusOK).Send(bJson)
}

func AddTransaction(c *fiber.Ctx) error {
	data := Trans{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	var cat enums.Category
	if data.Category == "Shopping" {
		cat = enums.Shopping
	} else if data.Category == "Food" {
		cat = enums.Food
	} else {
		cat = enums.Misc
	}
	var cur enums.Currency
	if data.Currency == "USD" {
		cur = enums.USD
	} else {
		cur = enums.INR
	}
	transTime, err := time.Parse(time.RFC3339, data.TransactionTime)
	t, err := transactions.AddTransaction(data.UserId, data.Amount, cur, cat, transTime)
	err = transactions.ProcessTransactions(t)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString("OK")
}

func CreateUser(c *fiber.Ctx) error {
	data := User{}
	if err := c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	id, err := users.CreateUser(data.Name)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	return c.Status(fiber.StatusOK).SendString(fmt.Sprintf("Created Id : %d", id))
}
