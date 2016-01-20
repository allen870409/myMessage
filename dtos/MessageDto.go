package dtos
import "time"

type MessageDto struct {
	Id      int    `orm:"auto"`
	Content string `orm:"size(200)"`
	User  int
	Created time.Time
}