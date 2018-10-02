package model

import (
	"testing"
)

func TestFeedbackMessages_Create(t *testing.T) {
	mesg := new(FeedbackMessages)
	mesg.ImgUrl="asdfasd"
	mesg.Create()
}
