package data

import (
	"taa-ex/internal/conf"
    "database/sql"
    "fmt"
    "github.com/aws/aws-sdk-go/service/ssm"
    "github.com/google/wire"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	_ "github.com/aws/aws-sdk-go/service/ssm"
	"github.com/go-kratos/kratos/v2/log"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTaaEngineRepository)

// Data .
type Data struct {
	db    *gorm.DB
	sqlDB *sql.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	l := log.NewHelper(logger)
	var db *gorm.DB
	var sqlDB *sql.DB
	var err error

	//TODO: Fetch the config from Common library
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-south-1")},
	)

	svc := ssm.New(sess)

	dbParameter, err := svc.GetParameter(&ssm.GetParameterInput{Name: aws.String("testUserDb"),
		WithDecryption: aws.Bool(true)})

	dbConn := dbParameter.Parameter.Value
	fmt.Printf("Conn - %s \n", *dbConn)

	// Open connection to the database
	if c.Database.Driver == "mysql" {
		dsn := fmt.Sprintf("%s?charset=utf8mb4&parseTime=True&loc=Local", *dbConn)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{PrepareStmt: true})
		if err != nil {
			l.Errorf("Unable to open db connection: %v", err)
		}

		sqlDB, err = db.DB()
		if err != nil {
			l.Errorf("Unable to get generic db interface: %v", err)
		}

		// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxIdleConns(int(c.Database.MaxIdleConns))

		// SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetMaxOpenConns(int(c.Database.MaxOpenConns))

		// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
		sqlDB.SetConnMaxLifetime(time.Duration(c.Database.GetMaxConnLifetimeInMins()) * time.Minute)
	}

	d := &Data{db, sqlDB}
	cleanup := func() {
		l.Info("closing the data resources")
		sqlDB.Close()
	}
	return d, cleanup, nil
}
