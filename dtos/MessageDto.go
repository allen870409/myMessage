package dtos
import "time"

type MessageDto struct {
	Id      int
	Content string
	User  int
	Created time.Time
}