package controllers

import (
	"io/ioutil"
	"github.com/astaxie/beego"
	. "myMessage/dtos"
	. "myMessage/models"
	"net/http"
	. "myMessage/services"
	"github.com/astaxie/beego/orm"
)

type MessageController struct {
	beego.Controller
}

func (this *MessageController) List() {
	limit, err := this.GetInt("limit", 0)
	if checkErr(this, err) {
		return
	}
	messages, err := ListMessages(limit)
	if checkErr(this, err){
		return
	}else{
		responseWithStatus(this, http.StatusOK, messages)
	}

}

func (this *MessageController) Get() {
	id, err := this.GetInt(":id")
	if checkErr(this, err) {
		return
	}
	message, err := GetMessage(id)
	if err == orm.ErrNoRows {
		responseWithStatus(this, http.StatusNotFound, &Message{Id: id})
	}else if checkErr(this, err){
		return
	}else{
		responseWithStatus(this, http.StatusOK, message)
	}
}

func (this *MessageController) Put() {
	id, err := this.GetInt(":id")
	if checkErr(this, err) {
		return
	}
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	if checkErr(this, err) {
		return
	}
	message, err := UpdateMessage(id, body)
	if checkErr(this, err){
		return
	}else{
		responseWithStatus(this, http.StatusOK, message)
	}
}

func (this *MessageController) Post() {
	body, err := ioutil.ReadAll(this.Ctx.Request.Body)
	if checkErr(this, err) {
		return
	}
	message, err := CreateMessage( body)
	if checkErr(this, err){
		return
	}else{
		responseWithStatus(this, http.StatusOK, message)
	}
}

func (this *MessageController) Delete() {
	id, err := this.GetInt(":id")
	if checkErr(this, err) {
		return
	}
	message, err := DeleteMessage(id)
	if checkErr(this, err){
		return
	}else{
		responseWithStatus(this, http.StatusOK, message)
	}
}

func checkErr(c *MessageController, err error) bool {
	if err != nil {
		result := &ResponseJson{Status: http.StatusInternalServerError, Data: err.Error()}
		c.Data["json"] = result
		c.ServeJSON()
		return true
	}
	return false
}

func responseWithStatus(c *MessageController, status int, data interface{}) {
	result := &ResponseJson{Status: status, Data: data}
	c.Data["json"] = result
	c.ServeJSON()
}
