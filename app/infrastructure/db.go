package infrastructure

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DB struct {
	Host       string
	Username   string
	Password   string
	DBName     string
	Port       string
	URI        string
	Connection *gorm.DB
}

func NewDB(c *Config) *DB {
	return newDB(&DB{
		Host:     c.DB.Host,
		Username: c.DB.Username,
		Password: c.DB.Password,
		DBName:   c.DB.DBName,
		Port:     c.DB.Port,
		URI:      c.DB.URI,
	})
}

func NewTestDB(c *Config) *DB {
	return newDB(&DB{
		Host:     c.DB.Host,
		Username: c.DB.Username,
		Password: c.DB.Password,
		DBName:   c.DB.DBName,
		Port:     c.DB.Port,
		URI:      c.DB.URI,
	})
}

func newDB(d *DB) *DB {
	db, err := gorm.Open(
		postgres.Open(
			"host=" + d.Host + " user=" + d.Username + " password=" + d.Password + " dbname=" + d.DBName + " port=" + d.Port,
		),
	)
	if err != nil {
		panic(err.Error())
	}
	d.Connection = db
	return d
}

func (db *DB) Begin() *gorm.DB {
	return db.Connection.Begin()
}

func (db *DB) Connect() *gorm.DB {
	return db.Connection
}
