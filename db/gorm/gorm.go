package gorm

import (
	//mysql driver

	//_ "github.com/go-sql-driver/mysql"

	"github.com/jinzhu/gorm"

	//mysql dialects
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"fmt"

	"github.com/spf13/viper"
)

var (
	//Db ...
	Db  *gorm.DB
	err error
)

//ConnectMySQL is connect mysql db with gorm
func ConnectMySQL() *gorm.DB {
	/*
		DBMS := "mysql"
		USER := "interfacedev"
		PASS := "Interface@APIDev"
		PROTOCOL := "tcp(172.22.228.224:3306)"
		DBNAME := "apishopdev"
	*/

	/* viper.SetConfigType("yaml")
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.ReadInConfig()
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_")) */

	DBMS := "mysql"
	USER := viper.GetString("mysql.user")
	PASS := viper.GetString("mysql.pass")
	PROTOCOL := viper.GetString("mysql.protocal")
	DBNAME := viper.GetString("mysql.db")

	CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?charset=utf8&parseTime=True"
	/*db,err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}*/

	fmt.Println("MySQL Connect:: ", CONNECT)

	if Db, err = gorm.Open(DBMS, CONNECT); err != nil {
		panic(err.Error())
	}

	//defer db.Close()
	// make sure connection is available
	if err = Db.DB().Ping(); err != nil {
		panic(err)
	}

	//db.LogMode(true)

	return Db
}

//DBManager return gorm db
func DBManager() *gorm.DB {
	return Db
}
