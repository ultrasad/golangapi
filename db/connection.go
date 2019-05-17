package db

import (
	"database/sql"
	"fmt"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

/*CreateCon Create mysql connection*/
func CreateCon() *sql.DB {
	//db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/golang_restful_api_sample_dev?charset=utf8&parseTime=True&loc=Local")
	db, err := sql.Open("mysql", viper.GetString("mysql.user")+":"+viper.GetString("mysql.pass")+"@"+viper.GetString("mysql.protocal")+"/"+viper.GetString("mysql.db")+"?charset=utf8&parseTime=True&loc=Local")
	db.Query("SET NAMES utf8")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("db is connected")
	}
	//defer db.Close()
	// make sure connection is available
	err = db.Ping()
	fmt.Println(err)
	if err != nil {
		fmt.Println("db is not connected")
		fmt.Println(err.Error())
	}
	return db
}

/*end mysql connection*/
