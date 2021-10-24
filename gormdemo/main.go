package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type User struct{
	Name string
	Age int
	Birthday time.Time
}

func main(){
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:root@tcp(127.0.0.1:3306)/bw_conductor?charset=utf8mb4&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256, // string 类型字段的默认长度
		DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil{
		fmt.Println("gorm.Open fail:", err)
	}

/*	mysqlDb, err := db.DB()
	mysqlDb.SetMaxOpenConns(10)  //最大打开连接
	mysqlDb.SetMaxIdleConns(5)   //最大空闲连接
	mysqlDb.SetConnMaxLifetime(100 * time.Second)*/


/*	user := User{Name: "Jinzhu", Age: 25, Birthday: time.Now()}
	db.Create(user)

	users := []User{{"tom", 12, time.Now()}, {"jack", 15, time.Now()}}
	db.Create(users)*/
	user := User{}
	db.First(&user)
	fmt.Println(user.Name, user.Age, user.Birthday)
}