package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres"
	"github.com/williamchang80/sea-apd/usecase"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot read .env file")
	}
}

func main() {
	e := echo.New()
	db := db.Postgres()
	k := postgres.NewProductRepository(db)
	t := usecase.NewProductUseCase(k)
	product.NewProductController(e, t)
	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, e)
}
