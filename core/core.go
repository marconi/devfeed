package core

import (
	"encoding/gob"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/boj/redistore"
	"github.com/gorilla/sessions"
	"labix.org/v2/mgo"
)

var (
	Config  *AppConfig
	Session *redistore.RediStore
	Mongo   *mgo.Session
	Db      *mgo.Database
)

type AppConfig struct {
	App struct {
		Port         int
		Templates    string
		SecretKey    string
		SessConnPool int
		SessName     string
		BaseUrl      string
		Debug        bool
	}
	Redis struct {
		Url string
	}
	Mongo struct {
		Name string
		Url  string
	}
	Email struct {
		DefaultFrom string
		Url         string
	}
	Sentry struct {
		Dns string
	}
}

func LoadConfig() {
	file, err := ioutil.ReadFile("config.json")
	if err != nil {
		panic("error opening config: " + err.Error())
	}

	Config = new(AppConfig)
	if err = json.Unmarshal(file, Config); err != nil {
		panic("error parsing config: " + err.Error())
	}
}

func InitMongo() {
	session, err := mgo.Dial(Config.Mongo.Url)
	if err != nil {
		panic(err)
	}

	// switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	Db = session.DB(Config.Mongo.Name)
	Mongo = session
}

func InitSessionStore(vals ...interface{}) {
	Session = redistore.NewRediStore(
		Config.App.SessConnPool,
		"tcp",
		Config.Redis.Url,
		"",
		[]byte(Config.App.SecretKey),
	)

	for _, val := range vals {
		gob.Register(val)
	}
}

func GetSession(r *http.Request) (*sessions.Session, error) {
	return Session.Get(r, Config.App.SessName)
}
