package db

import (
	//	"fmt"
	"stocknew/fortune/models"
)

func (dc *DBClient) GetStockDateData(code string, datesize int) ([]*models.DayInfo, error) {
	datas := make([]*models.DayInfo, 0)

	//查询数据
	rows, err := dc.Conn.Query("select `code`,`date`,`OpenPx`,`ClosePx`,`HighPx`,`LowPx`,`BusinessAmount`,`BusinessBalance` from stockdayinfos where Code=? order by id desc limit ?;", code, datesize)
	if err != nil {
		return datas, err
	}

	for rows.Next() {
		var code string
		var date string
		var oppx float64
		var clpx float64
		var hipx float64
		var lopx float64
		var ba float64
		var bb float64
		err := rows.Scan(&code, &date, &oppx, &clpx, &hipx, &lopx, &ba, &bb)
		if err != nil {
			return datas, err
		}
		datas = append(datas, &models.DayInfo{code, date, oppx, clpx, hipx, lopx, ba, bb})
	}

	return datas, nil
}
