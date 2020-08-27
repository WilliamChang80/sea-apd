package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/williamchang80/sea-apd/controller/http/product"
	"github.com/williamchang80/sea-apd/controller/http/user"
	"github.com/williamchang80/sea-apd/infrastructure/config"
	"github.com/williamchang80/sea-apd/infrastructure/db"
	"github.com/williamchang80/sea-apd/repository"
	product2 "github.com/williamchang80/sea-apd/usecase/product"
	ucUser "github.com/williamchang80/sea-apd/usecase/user"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Cannot read .env file")
	}
}

func main() {
	conf := config.New()
	appHost := fmt.Sprintf("http://%s:%s", conf.AppHost, conf.AppPort)

	r := mux.NewRouter()
	db := db.Postgres()

	k := repository.NewProductRepository(db)
	t := product2.NewProductUseCaseImpl(k)
	product.NewProductController(r, t)

	userRepo := repository.NewUserRepository(db)
	userCImpl := ucUser.NewAdminUseCaseImpl(userRepo)
	user.NewAdminController(r, userCImpl)

	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(conf.AppPort, r)
}
