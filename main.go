package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/williamchang80/sea-apd/usecase/user"

	userCtrl "github.com/williamchang80/sea-apd/controller/http/user"

	"github.com/labstack/echo"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository/postgres"
	product2 "github.com/williamchang80/sea-apd/usecase/product"
)

func main() {
	e := echo.New()
	db := db.Postgres()

	k := postgres.NewProductRepository(db)
	t := product2.NewProductUseCase(k)
	product.NewProductController(e, t)

	userRepo := postgres.NewUserRepository(db)
	adminUC := user.NewAdminUseCase(userRepo)
	userCtrl.NewAdminController(e, adminUC)

	appPort := ":" + os.Getenv("APP_PORT")
	appHost := fmt.Sprintf("http://%s%v", os.Getenv("APP_HOST"), appPort)
	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(appPort, e)
}
