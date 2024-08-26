package config

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabase(viper *viper.Viper, zapLog *zap.Logger) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	gormConfig := &gorm.Config{
		Logger: logger.New(&zapLogWriter{Logger: zapLog}, logger.Config{
			SlowThreshold:             time.Second * 5,
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
	default:
		if dbDsn == "" {
			dbDsn = "db/gorm.db"
		}
		db, err = gorm.Open(sqlite.Open(dbDsn), gormConfig)
	}
	if err != nil {
		zapLog.Fatal("failed to connect database:", zap.Error(err))
	}

	connection, err := db.DB()
	if err != nil {
		zapLog.Fatal("failed to connect database:", zap.Error(err))
	}

	idleConnection := viper.GetInt("database.pool.idle")
	maxConnection := viper.GetInt("database.pool.max")
	maxLifeTimeConnection := viper.GetInt("database.pool.lifetime")
	connection.SetMaxIdleConns(idleConnection)
	connection.SetMaxOpenConns(maxConnection)
	connection.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeConnection))

	return db
}

type zapLogWriter struct {
	Logger *zap.Logger
}

func (l *zapLogWriter) Printf(message string, args ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(fmt.Sprintf("%v", message[3:]), args[1:]...))
}
