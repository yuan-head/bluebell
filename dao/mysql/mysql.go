package mysql

import (
	"bluebell/settings"
	"fmt"
	"github.com/spf13/viper"
	"go.uber.org/zap"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&l",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName,
	)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect DB failed, err:#{err}")
		return
	}
	db.SetMaxOpenConns(viper.GetInt("mysql.maxOpenConns"))
	db.SetMaxIdleConns(viper.GetInt("mysql.maxIdleConns"))
	return
}

func Close() {
	_ = db.Close()
}
