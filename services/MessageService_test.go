package services
import (
	"testing"
	"time"
	. "myMessage/models"
	"math/rand"
	"encoding/json"
	"github.com/astaxie/beego/config"
	_ "github.com/go-sql-driver/mysql"
)

const (
	configPath = "../conf/app.conf"
	db         = "mysql"
)

var user_ids  = [4]int{1, 2, 3, 4}
var index int = 0
var testTime = time.Now().Format("2006_01_02_15_04_05")
var MyConfig config.Configer


func TestCreateMessage(t *testing.T){
	message := &Message{Id: -1, Content: "TestCreateMessage " + testTime, Created:time.Now(), User:&User{Id:user_ids[rand.Int() % 4]} }
	body, _ := json.Marshal(message)
	if result, err := CreateMessage(body); err != nil {
		t.Error("TestCreateMessageNormal Error!" + err.Error())
	}else {
		t.Log("TestCreateMessageNormal successful!\n" + result.String())
		index = result.Id
	}
}

func TestListMessages(t *testing.T){
	if _, err := ListMessages(0); err != nil {
		t.Error("TestListMessages Error!" + err.Error())
	}else {
		t.Log("TestListMessages successful!\n")
	}
}


func TestUpdateMessage(t *testing.T){
	message := &Message{Id: index, Content: "TestUpdateMessage " + testTime, Created:time.Now(), User:&User{Id:user_ids[rand.Int()%4]} }
	body, _ := json.Marshal(message)
	if result, err := CreateMessage(body); err != nil {
		t.Error("TestUpdateMessage Error!" + err.Error())
	}else {
		t.Log("TestUpdateMessage successful!\n" + result.String())
	}
}


func TestGetMessage(t *testing.T){
	if result, err := GetMessage(index); err != nil {
		t.Error("TestGetMessage Error!" + err.Error())
	}else {
		t.Log("TestGetMessage successful!\n" + result.String())
	}
}

func TestDeleteMessage(t *testing.T){
	if result, err := DeleteMessage(index); err != nil {
		t.Error("TestDeleteMessage Error!" + err.Error())
	}else {
		t.Log("TestDeleteMessage successful!\n" + result.String())
	}
}







