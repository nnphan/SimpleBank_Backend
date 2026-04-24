package initialize

import (
	"log"
	"simplebank/global"
	"simplebank/internal/util"
	"time"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgreSQL()  {
	m := global.Config.Postgresql
	dsn := util.BuildPostgreSQLDSN(m.Host, m.Port, m.User, m.Password, m.DBName, m.SSLMode)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
	})
	if err != nil {
		log.Fatal("cannot connect to DB: ", err)
	}
	util.CheckPaniceError(err, "cannot connect to DB: ")
	global.Logger.Info("Successfully connected to PostgreSQL database")
	global.Mdb = db

	// Set connection pool settings
	SetPool()
}

func SetPool() {
	p := global.Config.Postgresql
	sqlDB, err := global.Mdb.DB()
	if err != nil {
		log.Fatal("cannot get sql.DB from gorm.DB: ", err)
	}
	sqlDB.SetConnMaxIdleTime(time.Duration(p.MaxIdleConns)) // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxOpenConns(p.MaxOpenConns) // SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetConnMaxLifetime(time.Duration(p.ConnMaxLifetime)) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.

}


func migrateTable() {


}


