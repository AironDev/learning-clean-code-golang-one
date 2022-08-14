package main

import (
	"database/sql"
	"fmt"
	"net/url"

	cfg "github.com/airondev/learning-clean-code-golang-one/config"
	httpDeliver "github.com/airondev/learning-clean-code-golang-one/interface/http"
	"github.com/airondev/learning-clean-code-golang-one/repository"
	"github.com/airondev/learning-clean-code-golang-one/usecase"
	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"
)

func main() {
	config := cfg.Init()

	dbHost := config.DatabaseHost
	dbPort := config.DatabasePort
	dbUser := config.DatabaseUser
	dbPass := config.DatabasePassword
	dbName := config.DatabaseName

	connection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	val := url.Values{}
	val.Add("parseTime", "1")
	val.Add("loc", "Asia/Jakarta")
	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`mysql`, dsn)
	if err != nil {
		fmt.Println(err)
	}
	defer dbConn.Close()

	e := echo.New()

	ar := repository.NewMysqlArticleRepository(dbConn)
	au := usecase.NewArticleUsecase(ar)
	httpDeliver.NewArticleHttpHandler(e, au)

	err = e.Start(":8000")
	if err != nil {
		fmt.Println(err)
	}
}
