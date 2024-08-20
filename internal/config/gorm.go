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

	dbType := viper.GetString("database.type")
	switch dbType {
	case "pg":
		db, err = postgreSQLConnection(viper, gormConfig)
	case "sqlite":
		db, err = sqliteConnection(viper, gormConfig)
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

func postgreSQLConnection(viper *viper.Viper, gormConfig *gorm.Config) (db *gorm.DB, err error) {
	url := viper.GetString("database.pg.url")
	return gorm.Open(postgres.Open(url), gormConfig)
}

func sqliteConnection(viper *viper.Viper, gormConfig *gorm.Config) (db *gorm.DB, err error) {
	dbFile := viper.GetString("database.sqlite.file")
	return gorm.Open(sqlite.Open(dbFile), gormConfig)
}

type zapLogWriter struct {
	Logger *zap.Logger
}

func (l *zapLogWriter) Printf(message string, args ...interface{}) {
	l.Logger.Debug(fmt.Sprintf(fmt.Sprintf("%v", message[3:]), args[1:]...))
}
