package model

type feedbackMessges struct {
	ID  int64 `gorm:"primary_key; AUTO_INCREMENT" json:"id"`
	UserName  string

}

