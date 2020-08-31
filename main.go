package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/williamchang80/sea-apd/controller/http/user"

	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres"
	"github.com/williamchang80/sea-apd/usecase"
)

func main() {
	e := echo.New()
	db := db.Postgres()

	k := postgres.NewProductRepository(db)
	t := usecase.NewProductUseCase(k)
	product.NewProductController(e, t)

	userRepo := postgres.NewUserRepository(db)
	adminUC := usecase.NewAdminUseCase(userRepo)
	user.NewAdminController(e, adminUC)

	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, e)
}
