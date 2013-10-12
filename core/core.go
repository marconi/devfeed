package core

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/boj/redistore"
	log "github.com/cihub/seelog"
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

// merge settings from envars
func mergeEnv(s interface{}) {
	rType := reflect.TypeOf(s)
	rVal := reflect.ValueOf(s)

	numFields := rType.Elem().NumField()
	for i := 0; i < numFields; i++ {
		fType := rType.Elem().Field(i)
		fVal := rVal.Elem().Field(i)

		fNumFields := fType.Type.NumField()
		for j := 0; j < fNumFields; j++ {
			ffField := fType.Type.Field(j)
			ffVal := fVal.Field(j)

			envName := strings.ToUpper(fmt.Sprintf("%s_%s", fType.Name, ffField.Name))
			switch ffField.Type.Kind() {
			case reflect.Int:
				envVal := os.Getenv(envName)
				if envVal != "" {
					envVal, err := strconv.ParseInt(envVal, 10, 0)
					if err != nil {
						log.Info("Unable to parse settings: ", envName)
					} else {
						ffVal.SetInt(envVal)
					}
				}
			case reflect.String:
				envVal := os.Getenv(envName)
				if envVal != "" {
					ffVal.SetString(envVal)
				}
			case reflect.Bool:
				envVal := os.Getenv(envName)
				if envVal != "" {
					envVal, err := strconv.ParseBool(envVal)
					if err != nil {
						log.Info("Unable to parse settings: ", envName)
					} else {
						ffVal.SetBool(envVal)
					}
				}
			}
		}
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

	// merge settings from envars
	mergeEnv(Config)
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
