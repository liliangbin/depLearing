package utils

import (
	"testing"
	"fmt"
	"encoding/json"
)

type YBUser struct {
	VisitTime  string     `json:"visit_time"`
	VisitUser  VisitUser  `json:"visit_user"`
	VisitOauth VisitOauth `json:"visit_oauth"`
}
type VisitUser struct {
	Userid   string `json:"userid"`
	Username string `json:"username"`
	Usernick string `json:"usernick"`
	Usersex  string `json:"usersex"`
}

type VisitOauth struct {
	accessToken  string `json:"access_token"`
	tokenExpires string `json:"token_expires"`
}

func TestAECCBCDecrypter(t *testing.T) {

	str := "8b78161dd4ae844629deb86143449339fb8e001b93b3b4b2861bf8a3837bf97059a7a9e49739cce002fef5e64467a77eee8a21714aa7aed77c03984353429322e71b9721ee66849b3a029224ee4d5688908022f7a776179074dcbe25ec6fa3b33b9ed9495a0000ea4586fd543f081ca3bb0f7c7da5e65f3f97806c6229658adcb3be1d4d61fcf774fbbd443cec2a30167c2b8addafae1a8cbfda7231845d14cf2d55537612b00d2ba3b7e18c034c46f38d574a8a1b4bdd5a1b962786d0b20ff7b6205ed476f30cd91aee13fc35bb221804d09c05182038536d844e03533861fcaa3f257ebd4ba8b8ac4b12897524aece"

	str = AECCBCDecrypter(str)
	fmt.Println(str)
	user := new(YBUser)
	if err := json.Unmarshal([]byte(str), user); err != nil {
		panic(err.Error())
	}

	fmt.Println(user)
}
