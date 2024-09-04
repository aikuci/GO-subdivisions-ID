package config

import (
	"log"
	"os"
	"time"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	gormConfig := &gorm.Config{
		Logger: logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold:             5 * time.Second,
			Colorful:                  false,
			IgnoreRecordNotFoundError: true,
			ParameterizedQueries:      true,
			LogLevel:                  logger.Info,
		}),
	}

	dbDialect := viper.GetString("database.dialect")
	dbDsn := viper.GetString("database.dsn")
	switch dbDialect {
	case "pg":
		if dbDsn == "" {
			dbDsn = "postgresql://postgres:postgres@localhost:5432"
		}
		db, err = gorm.Open(postgres.Open(dbDsn), gormConfig)
	case "sqlite":
		if dbDsn == "" {
			dbDsn = "db/gorm.db"
		}
		db, err = gorm.Open(sqlite.Open(dbDsn), gormConfig)
	default:
		log.Fatalf("unsupported database dialect: %s", dbDialect)
	}

	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	connection, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}

	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")
	if maxLifeTimeConnection <= 0 {
		maxLifeTimeConnection = 3600 // default to 1 hour if not configured
	}
	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}
