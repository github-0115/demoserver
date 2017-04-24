package db

import (
	c "CoreSystemBase/config"
	"database/sql"
	"fmt"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	log "github.com/inconshreveable/log15"
	"gopkg.in/mgo.v2"
)

var (
	Session   *mgo.Session
	DBUserUri string
	User      *MongoSession
	Lables    *sql.DB
)

func init() {
	DBUserUri = c.DBCfg.UserCenterMongoTask.DB
	mongoUrl := c.DBCfg.UserCenterMongoTask.String()

	var err error
	Session, err = mgo.Dial(mongoUrl)
	if err != nil {
		log.Error(fmt.Sprintf("connect Mongo error", err))
		panic(err)
	}

	User = new(MongoSession).Init(
		c.DBCfg.UserCenterMongoTask,
		"user",
	)

	Lables, err = sql.Open("mysql", c.DBCfg.LablesSqlTask.User+":"+c.DBCfg.LablesSqlTask.Password+"@tcp("+c.DBCfg.LablesSqlTask.Host+":"+strconv.Itoa(int(c.DBCfg.LablesSqlTask.Port))+")/"+c.DBCfg.LablesSqlTask.DB+"?charset=utf8")
	if err != nil {
		log.Error(fmt.Sprintf("connect mysql error", err))
		panic(err)
	}
	Lables.SetMaxOpenConns(2000)
	Lables.SetMaxIdleConns(1000)
	Lables.Ping()
}
