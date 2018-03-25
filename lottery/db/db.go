package db

import (
	"database/sql"
	"errors"
	"fmt"
	//	"stocknew/lottery/model"
	//	"strconv"
	//	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "pkpk"           //用户名
	dbpassword = "@Dyf19840218@"  //密码
	dbname     = "pkpk"           //数据库名
)

type DBClient struct {
	Conn *sql.DB
}

var dbconn *DBClient

func Init() error {
	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostsip+")/"+dbname+"?charset=utf8")
	if err != nil {
		return err
	}
	dbconn = &DBClient{}
	dbconn.Conn = db

	return nil
}

func GetDB() *DBClient {
	return dbconn
}

func (dc *DBClient) InsertData(k string, v []string) error {
	var period sql.NullString

	tx, err := dc.Conn.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query("select period from PK10 where period=?", k)

	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		tx.Rollback()
		return err
	}
	for rows.Next() {
		err := rows.Scan(&period)
		if err != nil {
			fmt.Printf(err.Error())
			rows.Close()
			tx.Rollback()
			return err
		}
		if period.Valid {
			rows.Close()
			tx.Rollback()
			return errors.New("数据已经录入。")
		}
	}
	err = rows.Err()
	if err != nil {
		fmt.Printf(err.Error())
		rows.Close()
		tx.Rollback()
		dc.Conn.Close()
		return err
	}
	rows.Close()
	insertstr := fmt.Sprintf("insert into PK10(period,number1,number2,number3,number4,number5,number6,number7,number8,number9,number10) values (%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v)", k, v[0], v[1], v[2], v[3], v[4], v[5], v[6], v[7], v[8], v[9])
	//	fmt.Println(insertstr)
	stm, err := tx.Prepare(insertstr)
	if err != nil {
		tx.Rollback()
		dc.Conn.Close()
		return err
	}
	_, err = stm.Exec()
	if err != nil {
		tx.Rollback()
		dc.Conn.Close()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		dc.Conn.Close()
		return err
	}
	//	dc.Conn.Close()
	return nil

}
