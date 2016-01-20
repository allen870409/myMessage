package models

import (
	"time"
	"fmt"
)

type Message struct {
	Id      int    `orm:"auto"`
	Content string `orm:"size(200)"`
	User  *User  `orm:"rel(fk)"`
	Created time.Time
}

func (message Message) String() string {
	return fmt.Sprintf("%d\t%s", message.Id, message.Content)
}

