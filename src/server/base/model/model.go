package model
import (
	"github.com/jinzhu/gorm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
var (
	MysqlConn    *gorm.DB
	err	error
)

type Mysql struct {
	Name     string `json:"name"`
	User     string `json:"user" default:"root"`
	Password string `json:"password" required:"true" env:"db_password"`
	Host     string `json:"host"`
	Port     string `default:"3306" json:"port"`
}

func init()  {
	var mysql = Mysql{
		Name: "wjecho",
		User: "root",
		Host: "localhost",
		Port: "3306",
		Password: "159874"}

	setupMysqlConn(&mysql)
}

func setupMysqlConn(mysql *Mysql) {
	var connectionString string
	if mysql.Password == ""{
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.Name)
	}else{
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", mysql.User, mysql.Password, mysql.Host, mysql.Port, mysql.Name)
	}
	MysqlConn, err = gorm.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}
	err = MysqlConn.DB().Ping()
	if err != nil {
		panic(err)
	}
	MysqlConn.LogMode(true)

	MysqlConn.DB().SetMaxIdleConns(10)
	MysqlConn.DB().SetMaxOpenConns(100)

	MysqlConn.Set("gorm:table_options", "ENGINE=InnoDB")
}
