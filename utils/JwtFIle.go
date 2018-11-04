package utils

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"net/http"
	"log"
	"time"
	"encoding/json"
	"golang_tes/depLearing/model"
	"golang_tes/depLearing/service"
)

const appName = "石大请销假"
const jwtSecret = "lililili"

func OAuthHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	fmt.Println(r.Form)
	if err != nil {
		w.WriteHeader(http.StatusForbidden)
		fmt.Fprintf(w, "error in request")
	}
	/*vq appName */
	vq := r.Form["vq"][0]

	userInfo, err := service.GetStuInfoByVq(appName, vq)
	if err != nil {
		panic(err.Error())
	}
	log.Println("第三步")
	claims := make(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
	claims["iat"] = time.Now().Unix()
	claims["id"] = "123456"
	claims["student_id"] = userInfo.YbStudentid
	claims["vq"] = vq
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(jwtSecret))
	fmt.Println(tokenString, err)
	log.Println(tokenString + "   " + w.Header().Get("status"))
	response := model.Token{tokenString}
	JsonResponse(response, w)

}

func ValidateTokenMiddleWare(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	r.ParseForm()
	fmt.Println(r.Form)
	var tokenString = r.Form["token"][0]
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
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

func GetStuInfoByToken(tokenString string) (*model.YBUserInfo, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})

	if token.Valid {
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			fmt.Println(claims["id"], claims["student_id"])
			vq := claims["vq"]
			v := vq.(string)
			userInfo, err := service.GetStuInfoByVq(appName, v)
			return userInfo, err
		} else {
			fmt.Println(err)
			return nil,err
		}

	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			fmt.Println("That's not even a token")
			return nil, err
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			fmt.Println("Timing is everything")
			return nil, err
		} else {
			fmt.Println("Couldn't handle this token:", err)
			return nil, err
		}
	} else {
		fmt.Println("Couldn't handle this token:", err)
		return nil, err
	}
}

func JsonResponse(response interface{}, w http.ResponseWriter)  {

	log.Println("jsonresponse")
	json, err := json.Marshal(response)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

