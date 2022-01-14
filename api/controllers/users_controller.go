package controllers

import (
	"net/http"
	"github.com/wentyocta/majoo/api/models"
	"github.com/wentyocta/majoo/api/responses"
	"io/ioutil"
	"encoding/json"
)

// func (server *Server) GetUsers(w http.ResponseWriter, r *http.Request) {
//
// 	user := models.User{}
//
// 	users, err := user.FindAllUsers(server.DB)
// 	if err != nil {
// 		responses.ERROR(w, http.StatusInternalServerError, err)
// 		return
// 	}
// 	responses.JSON(w, http.StatusOK, users)
// }

func (server *Server) Login(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	user := models.User{}
	err = json.Unmarshal(body, &user)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user.Prepare()
	err = user.Validate("login")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIn(user.Username, user.Password)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err.Error())
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIn(Username, password string) (string, error) {

	var err error

	user := models.User{}

	err = server.DB.Debug().Model(models.User{}).Where("username = ?", Username).Take(&user).Error
	if err != nil {
		return "", err
	}
	err = models.VerifyPassword(user.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateToken(user.ID)
}