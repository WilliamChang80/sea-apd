package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/williamchang80/sea-apd/controller/http/user"

	ucUser "github.com/williamchang80/sea-apd/usecase/user"

	"github.com/williamchang80/sea-apd/repository"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/williamchang80/sea-apd/infrastructure/config"
	"github.com/williamchang80/sea-apd/infrastructure/db"
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

	userRepo := repository.NewUserRepository(db)
	userCImpl := ucUser.NewAdminUseCaseImpl(userRepo)
	user.NewAdminController(r, userCImpl)

	fmt.Println("App is running on " + appHost)
	http.ListenAndServe(conf.AppPort, r)
}
