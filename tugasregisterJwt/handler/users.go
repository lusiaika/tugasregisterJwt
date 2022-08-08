package handler

import (
	"context"
	"encoding/json"
	"net/http"
	"tugasregisterjwt/entity"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"

	"github.com/gorilla/mux"
)

type userAPI struct{}

//url/api
func InstallUserAPI(r *mux.Router) {
	s := r.PathPrefix("/user").Subrouter()
	api := userAPI{}

	s.HandleFunc("/login", api.login).Methods("POST")
	s.HandleFunc("/register", api.registerUsersHandler).Methods("POST")
}

func (h *userAPI) registerUsersHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	decoder := json.NewDecoder(r.Body)
	var user entity.User
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	pwd, err := HashingPassword(user.Password)
	if err != nil {
		w.Write([]byte("error hashing password "))
		return
	}

	user.Password = pwd
	users, err := SqlConnect.Register(ctx, user)
	if err != nil {
		writeJsonResp(w, statusError, err.Error())
		return
	}
	writeJsonResp(w, statusSuccess, users)
}

func (h *userAPI) login(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	decoder := json.NewDecoder(r.Body)
	var user entity.UserLogin
	if err := decoder.Decode(&user); err != nil {
		w.Write([]byte("error decoding json body"))
		return
	}

	users, err := SqlConnect.Login(ctx, user.Username)
	if err != nil {
		writeJsonResp(w, statusError, err.Error())
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(user.Password))
	if err != nil {
		w.Write([]byte(`wrong password`))
		return
	}

	claims := entity.JwtClaims{
		Uid: user.Username,
		//Pwd: user.Password,
	}

	token := jwt.NewWithClaims(
		JwtSigningMethod,
		claims,
	)

	retVal, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		writeJsonResp(w, statusError, "BAD_REQUEST")
		return
	}

	result := map[string]string{
		"token": retVal,
	}
	writeJsonResp(w, statusSuccess, result)

}
