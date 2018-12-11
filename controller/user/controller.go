package user

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/onkiit/db-monitor/lib/helper"

	"github.com/onkiit/db-monitor/model/user"
	"golang.org/x/crypto/bcrypt"
)

type Controller struct {
	store user.Store
}

func hashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), 13)
	if err != nil {
		return "", err
	}
	return string(bytePassword), nil
}

func (c *Controller) Register(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	u := new(user.User)
	if err := d.Decode(u); err != nil {
		log.Println("decode", err)
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	uuid, err := helper.GenerateUUID()
	if err != nil {
		log.Println("generate uuid", err)
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	hashed, err := hashPassword(u.Password)
	if err != nil {
		log.Println("hash", err)
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	u.UUID = uuid
	u.Password = hashed
	if err := c.store.AddUser(r.Context(), u); err != nil {
		log.Println("add user", err)
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, nil)
}

func (c *Controller) Login(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	l := new(user.Login)
	if err := d.Decode(l); err != nil {
		log.Println("decode", err)
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	u, err := c.store.GetByUsername(r.Context(), l.Username)
	if err != nil {
		log.Println("not found", err)
		helper.ErrorResponse(w, http.StatusNotFound, err)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(l.Password)); err != nil {
		log.Println("compare hash", err)
		helper.ErrorResponse(w, http.StatusForbidden, err)
		return
	}

	helper.RespondWithJSON(w, http.StatusOK, u)
}

func (c *Controller) CheckAvailable(w http.ResponseWriter, r *http.Request) {
	d := json.NewDecoder(r.Body)
	u := new(user.AvailableUser)
	if err := d.Decode(u); err != nil {
		log.Println("decode", err)
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	u.IsRegistered = false
	user, err := c.store.GetByUsername(r.Context(), u.Username)
	if err != nil {
		helper.ErrorResponse(w, http.StatusBadRequest, err)
		return
	}

	if user != nil {
		u.IsRegistered = true
	}

	helper.RespondWithJSON(w, http.StatusOK, u)

}
