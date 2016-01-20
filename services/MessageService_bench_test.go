package services
import (
	"testing"
	"time"
	. "myMessage/models"
	"math/rand"
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"log"
	"github.com/astaxie/beego/config"
	"myMessage/models"
	_ "github.com/go-sql-driver/mysql"
	"strconv"
)
func init() {
	myConfig, err := config.NewConfig("ini", configPath)
	if err != nil {
		log.Panicf("Can't find config file at '%v'.", configPath)
	} else {
		MyConfig = myConfig
	}

	orm.RegisterModel(new(models.User))
	orm.RegisterModel(new(models.Message))

	dbUser := myConfig.String("db_user")
	dbPwd := myConfig.String("db_pwd")
	dbPort := myConfig.String("db_port")
	dbHost := myConfig.String("db_host")
	dbName := myConfig.String("db_name")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dbUser+":"+dbPwd+"@tcp("+dbHost+":"+dbPort+")/"+dbName+"?charset=utf8")
	orm.RunSyncdb("default", false, true)
}
func Benchmark_CreateMessage(b *testing.B){

	for i := 0; i < b.N; i++ {
		message := &Message{Id: -1, Content: "Benchmark_CreateMessage " + strconv.Itoa(i), Created:time.Now(), User:&User{Id:user_ids[rand.Int()%4]} }
		body, _ := json.Marshal(message)
		CreateMessage(body)
	}
}

func Benchmark_ListMessages(b *testing.B){
	limit := rand.Int() / 10
	for i := 0; i < b.N; i++ {
		ListMessages(limit)
	}
}


func Benchmark_UpdateMessage(b *testing.B){
	for i := 0; i < b.N; i++ {
		message := &Message{Id: -1, Content: "Benchmark_UpdateMessage " + strconv.Itoa(i), Created:time.Now(), User:&User{Id:user_ids[rand.Int()%4]} }
		body, _ := json.Marshal(message)
		CreateMessage(body)
	}
}

func Benchmark_GetMessage(b *testing.B){
	for i := 0; i < b.N; i++ {
		GetMessage(i)
	}
}

func Benchmark_DeleteMessage(b *testing.B){
	for i := 0; i < b.N; i++ {
		DeleteMessage(i)
	}
}





