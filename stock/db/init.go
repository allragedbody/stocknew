package db

import (
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"takistan/stock/model"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const InitDate = "20130101"

var (
	dbhostsip  = "127.0.0.1:3306" //IP地址
	dbusername = "root"           //用户名
	dbpassword = ""               //密码
	dbname     = "stock"          //表名
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

func (dc *DBClient) GetDateRange(code string) (string, string, error) {

	//查询数据
	rows, err := dc.Conn.Query("select max(Date) from stockdayinfos where code=?", code)
	if err != nil {
		return "notupdate", "notupdate", err
	}
	today := time.Now().Format("20060102")
	for rows.Next() {
		var c sql.NullString
		err := rows.Scan(&c)
		if err != nil {
			return "error", "error", err
		}
		if !c.Valid {
			return InitDate, today, nil
		}

		if c.String == today {
			return "notupdate", "notupdate", errors.New("notupdate")
		} else {
			t, _ := time.Parse("20060102", c.String)
			return t.Add(time.Duration(86400 * time.Second)).Format("20060102"), today, nil
		}
	}
	return "notupdate", "notupdate", errors.New("notupdate")
}

func (dc *DBClient) InsertData(dinfos []*model.DayInfo) error {
	tx, err := dc.Conn.Begin()
	if err != nil {
		tx.Rollback()
		return err
	}
	//	Code            string
	//	Date            string
	//	OpenPx          float64
	//	ClosePx         float64
	//	HighPx          float64
	//	LowPx           float64
	//	BusinessAmount  float64
	//	BusinessBalance float64
	if len(dinfos) == 0 {
		fmt.Printf("插入股票数据为零。")
		return nil
	}

	benchstr := ""
	for index, dinfo := range dinfos {
		if index == 0 {
			benchstr += "('" + dinfo.Code + "','" + dinfo.Date + "'," + strconv.FormatFloat(dinfo.OpenPx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.ClosePx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.HighPx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.LowPx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.BusinessAmount, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.BusinessBalance, 'G', -2, 32) + ")"
		} else {
			benchstr += ",('" + dinfo.Code + "','" + dinfo.Date + "'," + strconv.FormatFloat(dinfo.OpenPx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.ClosePx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.HighPx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.LowPx, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.BusinessAmount, 'G', -2, 32) + "," + strconv.FormatFloat(dinfo.BusinessBalance, 'G', -2, 32) + ")"

			//			benchstr += ",(" + dinfo.Code + "," + dinfo.Date + "," + dinfo.OpenPx + "," + dinfo.ClosePx + "," + dinfo.HighPx + "," + dinfo.LowPx + "," + dinfo.BusinessAmount + "," + dinfo.BusinessBalance + "," + ")"
		}
	}

	insertstr := fmt.Sprintf("insert into stockdayinfos(Code,Date,OpenPx,ClosePx,HighPx,LowPx,BusinessAmount,BusinessBalance) values %v", benchstr)
	//	fmt.Println(insertstr)

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
	return nil
}
