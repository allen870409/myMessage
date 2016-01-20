package es

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
	"net/url"
	"time"

	"github.com/belogik/goes"
	"myMessage/Godeps/_workspace/src/github.com/astaxie/beego/logs"
)

// NewES return a LoggerInterface
func NewES() logs.Logger {
	cw := &esLogger{
		Level: logs.LevelDebug,
	}
	return cw
}

type esLogger struct {
	*goes.Connection
	DSN   string `json:"dsn"`
	Level int    `json:"level"`
}

// {"dsn":"http://localhost:9200/","level":1}
func (el *esLogger) Init(jsonconfig string) error {
	err := json.Unmarshal([]byte(jsonconfig), el)
	if err != nil {
		return err
	}
	if el.DSN == "" {
		return errors.New("empty dsn")
	} else if u, err := url.Parse(el.DSN); err != nil {
		return err
	} else if u.Path == "" {
		return errors.New("missing prefix")
	} else if host, port, err := net.SplitHostPort(u.Host); err != nil {
		return err
	} else {
		conn := goes.NewConnection(host, port)
		el.Connection = conn
	}
	return nil
}

// WriteMsg will write the msg and level into es
func (el *esLogger) WriteMsg(msg string, level int) error {
	if level > el.Level {
		return nil
	}
	t := time.Now()
	vals := make(map[string]interface{})
	vals["@timestamp"] = t.Format(time.RFC3339)
	vals["@msg"] = msg
	d := goes.Document{
		Index:  fmt.Sprintf("%04d.%02d.%02d", t.Year(), t.Month(), t.Day()),
		Type:   "logs",
		Fields: vals,
	}
	_, err := el.Index(d, nil)
	return err
}

// Destroy is a empty method
func (el *esLogger) Destroy() {

}

// Flush is a empty method
func (el *esLogger) Flush() {

}

func init() {
	logs.Register("es", NewES)
}
