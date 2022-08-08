package handler

import (
	"encoding/json"
	"net/http"
	"tugasregisterjwt/database"

	"fmt"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

var SqlConnect *database.Database
var JwtSigningMethod = jwt.SigningMethodHS256

type response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

const (
	statusSuccess int = 0
	statusError   int = 1
)

var SecretKey = config.SecretKey

func writeJsonResp(w http.ResponseWriter, status int, obj interface{}) {

	resp := response{
		Status: status,
		Data:   obj,
	}
	w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(resp)
}

func HashingPassword(pwd string) (string, error) {
	// Hashing the password with the default cost of 10
	securePassword, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(securePassword), nil
}

func GetConnectionString() string {
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%d?database=%s",
		config.ConnectionString.Sqluser, config.ConnectionString.Sqlpassword, config.ConnectionString.Sqlserver, config.ConnectionString.Sqlport, config.ConnectionString.SqldbName)
	return connString
}
