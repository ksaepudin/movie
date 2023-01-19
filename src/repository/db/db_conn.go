package db

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DriverMySQL struct {
	config Database
	db     *gorm.DB
}

var (
	ConnMovieDb DbDriver
)

// DbDriver is object DB
type DbDriver interface {
	Db() interface{}
}

func init() {
	var err error
	ConnMovieDb, err = NewMySQLDriver(Database{
		Host:         "localhost",
		Port:         "3306",
		Username:     "root",
		Password:     "",
		Dbname:       "db_movie",
		Adapter:      "mysql",
		MaxOpen:      10,
		MaxIdle:      7,
		IdleLifeTime: 60,
		OpenLifeTime: 60,
	})
	if err != nil {
		log.Fatalf("unable to connect to Movie DB")
	}
}

// NewMySQLDriver new object SQL Driver
func NewMySQLDriver(config Database) (DbDriver, error) {
	dbConn, err := connect(config)
	// Disable table name's pluralization, if set to true, `User`'s table name will be `user`

	if err != nil {
		return nil, err
	}

	return &DriverMySQL{
		config: config,
		db:     dbConn,
	}, nil
}

func connect(config Database) (*gorm.DB, error) {
	user := config.Username
	password := config.Password
	host := config.Host
	port := config.Port
	dbname := config.Dbname

	dsn := user + ":" + password + "@(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=True&loc=Local"
	dbConn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return nil, err
	}
	conn, err := dbConn.DB()
	if err != nil {
		return nil, err
	}

	conn.SetMaxIdleConns(config.MaxIdle)
	conn.SetMaxOpenConns(config.MaxOpen)
	conn.SetConnMaxLifetime(time.Minute * time.Duration(config.OpenLifeTime))
	conn.SetConnMaxIdleTime(time.Minute * time.Duration(config.IdleLifeTime))

	return dbConn, err
}

// Db get db instance of gorm
func (m *DriverMySQL) Db() interface{} {
	return m.db
}
