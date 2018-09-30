package service

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"log"
	"strings"
	"time"
	"encoding/json"
	"golang_tes/depLearing/model"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("url", r.Method+"=====", r.Body)
	err :=r.ParseForm()
	fmt.Println(r.Form)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "error in request")
	}
	user := new(model.UserCredentials)
	user.Username = r.Form["username"][0]
	user.Password = r.Form["password"][0]
	log.Println(user.Password + "    fffff")

	log.Println("第二步")
	if strings.ToLower(user.Username) != "liliangbin" {
		if user.Password != "liliangbin" {
			w.WriteHeader(http.StatusForbidden)
			fmt.Println("Error login ")
			fmt.Fprintf(w, "Invalid credentials")

			return
		}
	}
	log.Println("第三步")
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = "123456"
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("lilili"))
	fmt.Println(tokenString, err)
	log.Println(tokenString + "   " + w.Header().Get("status"))
	response := model.Token{tokenString}
	JsonResponse(response, w)

}

func ValidateToeknMiddleWare(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	r.ParseForm() //这个地方不会出问题了
	fmt.Println(r.Form)
	var tokenString = r.Form["token"][0]
	//tokenString := ""
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte("lilili"), nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["id"], claims["id"])
		} else {
			fmt.Println(err)
		}

		fmt.Println("You look nice today")
		fmt.Fprintf(w, "index")
		next(w, r)
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "token is not valid ")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "token is not valid ")
		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "token is not valid ")
			fmt.Println("Couldn't handle this token:", err)
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
	}

}

func JsonResponse(response interface{}, w http.ResponseWriter) {

	log.Println("jsonresponse")
	json, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}
