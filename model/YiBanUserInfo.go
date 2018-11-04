package model

type YBUserInfo struct {
	YbUserid     string `json:"yb_userid"`
	YbUsername   string `json:"yb_username"`
	YbUsernick   string `json:"yb_usernick"`
	YbSex        string `json:"yb_sex"`
	YbMoney      string `json:"yb_money"`
	YbExp        string `json:"yb_exp"`
	YbUserhead   string `json:"yb_userhead"`
	YbSchoolid   string `json:"yb_schoolid"`
	YbSchoolname string `json:"yb_schoolname"`
	YbRealname   string `json:"yb_realname"`
	YbStudentid  string `json:"yb_studentid"`
	YbIdentity   string `json:"yb_identity"`
}

type GetInfo struct {
	AppName string
	AppVq   string
}
