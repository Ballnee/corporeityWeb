package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

var db *sqlx.DB

func Init() (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("mysql.user"), viper.GetString("mysql.password"), viper.GetString("mysql.host"),
		viper.GetInt("mysql.port"), viper.GetString("mysql.dbname"))
	// 也可以使用MustConnect连接不成功就panic
	println(dsn)
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}

	return
}

func Close() {
	_ = db.Close()
}
