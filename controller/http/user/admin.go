package user

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/mux"
	domainUser "github.com/williamchang80/sea-apd/domain/user"
	"github.com/williamchang80/sea-apd/usecase/user"

	"net/http"
)

// AdminController ...
type AdminController struct {
	ac user.AdminUseCase
}

// NewAdminController ...
func NewAdminController(r *mux.Router, a user.AdminUseCase) {
	c := &AdminController{
		ac: a,
	}
	r.HandleFunc("/users/v1/register-admin", c.RegisterAdmin)
}

// RegisterAdmin ...
func (a *AdminController) RegisterAdmin(r http.ResponseWriter, req *http.Request) {
	var u domainUser.User

	err := json.NewDecoder(req.Body).Decode(&u)
	if err != nil {
		http.Error(r, err.Error(), http.StatusBadRequest)
	}

	validateErr := u.Validate("")
	if validateErr != "" {
		http.Error(r, validateErr, http.StatusBadRequest)
	}

	user, err := a.ac.RegisterAdmin(u)
	if err != nil {
		log.Panic("Error")
	}

	o, err := json.Marshal(user)
	if err != nil {
		log.Panic("Error")
	}

	fmt.Fprintf(r, string(o))
}
