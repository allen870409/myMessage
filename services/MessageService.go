package services

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	. "myMessage/models"
	. "myMessage/dtos"
	"time"
)

func ListMessages(limit int) (*[]Message, error) {
	var messages []Message
	o := orm.NewOrm()
	_, err := o.QueryTable(new(Message)).RelatedSel().Limit(limit).All(&messages)
	return &messages, err
}

func GetMessage(id int) (*Message, error) {
	o := orm.NewOrm()
	message := &Message{Id: id}
	err := o.Read(message)
	if err != nil {
		return nil, err
	}
	if message.User != nil {
		o.Read(message.User)
	}
	return message, err
}

func UpdateMessage(id int, body []byte) (*Message, error)  {
	o := orm.NewOrm()
	o.Begin()
	messageDto := &MessageDto{Id: id}
	message := &Message{Id: id}
	err := o.Read(message)
	if err == nil {
		json.Unmarshal(body, messageDto)
		message.Id = int(id)
		message.Content = messageDto.Content
		if _, err := o.Update(message); err != nil {
			return nil, err
		}
		o.Commit()
	} else {
		o.Rollback()
		return nil, err
	}
	return &Message{Id: int(id)}, nil
}

func CreateMessage(body []byte) (*Message, error) {
	o := orm.NewOrm()
	o.Begin()
	messageDto := &MessageDto{}
	json.Unmarshal(body, messageDto)
	message := &Message{Id:-1, Content:messageDto.Content, Created:time.Now(), User:&User{Id:messageDto.User}}
	id, err := o.Insert(message)
	if err != nil{
		o.Rollback()
		return nil, err
	}else{
		o.Commit()
		return &Message{Id: int(id)}, nil
	}
}

func DeleteMessage(id int) (*Message, error) {
	o := orm.NewOrm()
	o.Begin()
	message := &Message{Id: id}
	if _, err := o.Delete(message); err != nil {
		o.Rollback()
		return nil, err
	} else {
		o.Commit()
		return message, nil
	}
}

