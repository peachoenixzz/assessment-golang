package api

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/peachoenixz/assessment/internal/database"
	"github.com/peachoenixz/assessment/internal/expense"
	"github.com/peachoenixz/assessment/pkg/log"
)

func serviceRouter() {
	e := echo.New()
	postgresDBClient := database.NewPostgres()
	expensePostgresRepo := expense.NewPostgres(postgresDBClient.Client)
	expenseServiceAPI := expense.NewService(expensePostgresRepo)
	expenseEndpoint := expense.NewEndpoint(expenseServiceAPI)
	e.POST("/expenses", expenseEndpoint.AddExpense)
	log.InfoLog("ECHO PREPARE TO START", "ECHO API")
	log.ErrorLog(e.Start(":2565"), "ECHO API")
}

func EchoStart() {
	serviceRouter()
	log.InfoLog("ECHO API STOP", "ECHO API")
	fmt.Println("start at port:", os.Getenv("PORT"))
}
