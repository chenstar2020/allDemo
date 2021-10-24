/*
 * 连接数据库
 */

package geeorm

import (
	"database/sql"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/session"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type Engine struct {
	db *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string)(e *Engine, err error){
	db, err := sql.Open(driver, source)
	if err != nil{
		log.ERRORF("sql.Open:[%v %v] fail, err:", driver, source, err)
		return
	}

	if err = db.Ping(); err != nil{
		log.ERRORF("db.Ping fail, err:", err)
		return
	}

	db.SetConnMaxLifetime(time.Minute)  //设置连接存活时间  从连接创建开始计算，到期强制关闭
	db.SetMaxOpenConns(10)           //设置最大打开连接  可以在程序启动时初始化多个连接, 如果打开连接数超过此值，将需要等待连接释放
	db.SetMaxIdleConns(5)            //设置最大空闲连接  空闲连接数不能超过maxOpenConn  达到最大空闲连接上限时，将会释放空闲连接

	dial, ok := dialect.GetDialect(driver)
	if !ok{
		log.ERRORF("dialect %s Not Found", driver)
		return
	}
	e = &Engine{
		db: db,
		dialect: dial,
	}

	return
}

func (e *Engine) Close(){
	if err := e.db.Close(); err != nil{
		log.ERRORF("db.Close failed, err:", err)
	}
}

func (e *Engine) NewSession() *session.Session{
	return session.New(e.db, e.dialect)
}