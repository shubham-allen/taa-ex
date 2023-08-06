package data

import (
	"database/sql"
	"fmt"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"taa-ex/internal/conf"
	"time"

	_ "github.com/aws/aws-sdk-go/service/ssm"
	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTaaEngineRepository)

// Data .
type Data struct {
	db    *gorm.DB
	sqlDb *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(logger)
	var db *gorm.DB
	var sqlDb *sql.DB
	var err error
	// Open connection to the database
	if c.Database.Driver == "mysql" {
		dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", c.Database.Source)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
		if err != nil {
			l.Errorf("Unable to open db connection: %v", err)
		}
		sqlDb, err = db.DB()
		if err != nil {
			l.Errorf("Unable to get generic db interface: %v", err)
		}
		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDb.SetMaxIdleConns(int(c.Database.MaxIdleConns))
		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDb.SetMaxOpenConns(int(c.Database.MaxOpenConns))
		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDb.SetConnMaxLifetime(time.Duration(c.Database.GetMaxConnLifetimeInMins()) * time.Minute)
	}
	d := &Data{db, sqlDb}
	cleanup := func() {
		l.Info("closing the data resources")
		sqlDb.Close()
	}
	return d, cleanup, nil
}
