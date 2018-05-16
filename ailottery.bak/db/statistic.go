package db

import (
	"time"

	"github.com/astaxie/beego/logs"
)


func (dc *DBClient) GetPutHistoryData() (map[int]int, error) {
	datas := make(map[int]int, 0)

	datas[1] = 0
	datas[2] = 0
	datas[3] = 0
	datas[4] = 0
	datas[5] = 0
	datas[6] = 0
	datas[7] = 0
	datas[8] = 0
	datas[9] = 0
	datas[10] = 0
	datas[11] = 0
	datas[12] = 0
	datas[13] = 0
	datas[14] = 0
	datas[15] = 0

	tmplist := ""
	tmptimes := 0
	//查询数据
	ndate := time.Now().Format("2006-01-02")
        sqlstr := "select currentPierod,numberList,putTime from HistoryPush where  createtime like '"+ndate+"%' and id< (select max(id) from HistoryPush where status='等开' and createtime like '"+ndate+"%');" 
        logs.Debug("Sql %v",sqlstr)
 
	rows, err := dc.Conn.Query(sqlstr)
	if err != nil {
		return datas, err
	}
	for rows.Next() {
		var currentPierod string
		var numberList string
		var putTime int

		err := rows.Scan(&currentPierod, &numberList, &putTime)
		if err != nil {
			return datas, err
		}
		if tmplist == "" {
			tmplist = numberList
			tmptimes = putTime
		} else {
			if tmplist == numberList {
				tmptimes = putTime
			} else {
				//中了，需要重新拿到puttime ，存储之前的puttime
				datas[tmptimes] += 1
				tmptimes = putTime
				tmplist = numberList
			}
		}

	}

	return datas, nil
}
