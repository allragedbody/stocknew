package db

import (
	"database/sql"
	"errors"
	"fmt"
	//	"stocknew/lotterytools/model"
	//	"strconv"
	//	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
        dbhostsip  = "127.0.0.1:3306" //IP地址
        dbusername = "pkpk"           //用户名
        dbpassword = "@Dyf19840218@"               //密码
        dbname     = "pkpk"        //数据库名

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

func (dc *DBClient) RestorePlanToDB(cp string, nl []int, pt int, rp int, sts string, gr bool, ct string) error {
	var sqlcp sql.NullString
	var grint int
	if gr == true {
		grint = 1
	} else {
		grint = 0
	}

	nll := fmt.Sprintf("%v", nl)

	tx, err := dc.Conn.Begin()
	if err != nil {
		return err
	}

	rows, err := tx.Query("select currentPierod from NewHistoryPush where currentPierod=?", cp)

	if err != nil {
		fmt.Printf("insert data error: %v\n", err)
		tx.Rollback()
		return err
	}
	for rows.Next() {
		err := rows.Scan(&sqlcp)
		if err != nil {
			fmt.Printf(err.Error())
			rows.Close()
			tx.Rollback()
			return err
		}
		if sqlcp.Valid {
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
		return err
	}

	rows.Close()
	insertstr := fmt.Sprintf("insert into NewHistoryPush(currentPierod,numberList,putTime,realPutTime,status,getReward,createTime) values ('%v','%v',%v,%v,'%v',%v,'%v')", cp, nll, pt, rp, sts, grint, ct)
	fmt.Printf(insertstr)
	stm, err := tx.Prepare(insertstr)
	if err != nil {
		tx.Rollback()
		return err
	}
	_, err = stm.Exec()
	if err != nil {
		tx.Rollback()
		return err
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return err
	}
	//	dc.Conn.Close()
	return nil

}
