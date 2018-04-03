package db

//	"fmt"
//	"stocknew/lotterytools/model"

func (dc *DBClient) GetLotterData(cur int, size int) ([][]string, error) {
	datas := make([][]string, 0)
	if cur == 0 {
		//查询数据
		rows, err := dc.Conn.Query("select `period`,`number1`,`number2`,`number3`,`number4`,`number5`,`number6`,`number7`,`number8`,`number9`,`number10` from PK10 order by period desc limit ?;", size)
		if err != nil {
			return datas, err
		}
		for rows.Next() {
			var period string
			var number1 string
			var number2 string
			var number3 string
			var number4 string
			var number5 string
			var number6 string
			var number7 string
			var number8 string
			var number9 string
			var number10 string

			err := rows.Scan(&period, &number1, &number2, &number3, &number4, &number5, &number6, &number7, &number8, &number9, &number10)
			if err != nil {
				return datas, err
			}
			data := make([]string, 0)
			data = append(data, period)
			data = append(data, number1)
			data = append(data, number2)
			data = append(data, number3)
			data = append(data, number4)
			data = append(data, number5)
			data = append(data, number6)
			data = append(data, number7)
			data = append(data, number8)
			data = append(data, number9)
			data = append(data, number10)

			datas = append(datas, data)
		}
	} else {
		rows, err := dc.Conn.Query("select `period`,`number1`,`number2`,`number3`,`number4`,`number5`,`number6`,`number7`,`number8`,`number9`,`number10` from PK10 where period <=? order by period desc limit ?;", cur, size)
		if err != nil {
			return datas, err
		}
		for rows.Next() {
			var period string
			var number1 string
			var number2 string
			var number3 string
			var number4 string
			var number5 string
			var number6 string
			var number7 string
			var number8 string
			var number9 string
			var number10 string

			err := rows.Scan(&period, &number1, &number2, &number3, &number4, &number5, &number6, &number7, &number8, &number9, &number10)
			if err != nil {
				return datas, err
			}
			data := make([]string, 0)
			data = append(data, period)
			data = append(data, number1)
			data = append(data, number2)
			data = append(data, number3)
			data = append(data, number4)
			data = append(data, number5)
			data = append(data, number6)
			data = append(data, number7)
			data = append(data, number8)
			data = append(data, number9)
			data = append(data, number10)

			datas = append(datas, data)
		}
	}

	return datas, nil
}

func (dc *DBClient) GetMissData(cur int, size int) ([][]string, error) {
	datas := make([][]string, 0)
	if cur == 0 {
		//查询数据
		rows, err := dc.Conn.Query("select `period`,`number1`,`number2`,`number3`,`number4`,`number5`,`number6`,`number7`,`number8`,`number9`,`number10` from PK10Miss order by period desc limit ?;", size)
		if err != nil {
			return datas, err
		}
		for rows.Next() {
			var period string
			var number1 string
			var number2 string
			var number3 string
			var number4 string
			var number5 string
			var number6 string
			var number7 string
			var number8 string
			var number9 string
			var number10 string

			err := rows.Scan(&period, &number1, &number2, &number3, &number4, &number5, &number6, &number7, &number8, &number9, &number10)
			if err != nil {
				return datas, err
			}
			data := make([]string, 0)
			data = append(data, period)
			data = append(data, number1)
			data = append(data, number2)
			data = append(data, number3)
			data = append(data, number4)
			data = append(data, number5)
			data = append(data, number6)
			data = append(data, number7)
			data = append(data, number8)
			data = append(data, number9)
			data = append(data, number10)

			datas = append(datas, data)
		}
	} else {
		rows, err := dc.Conn.Query("select `period`,`number1`,`number2`,`number3`,`number4`,`number5`,`number6`,`number7`,`number8`,`number9`,`number10` from PK10Miss where period <=? order by period desc limit ?;", cur, size)
		if err != nil {
			return datas, err
		}
		for rows.Next() {
			var period string
			var number1 string
			var number2 string
			var number3 string
			var number4 string
			var number5 string
			var number6 string
			var number7 string
			var number8 string
			var number9 string
			var number10 string

			err := rows.Scan(&period, &number1, &number2, &number3, &number4, &number5, &number6, &number7, &number8, &number9, &number10)
			if err != nil {
				return datas, err
			}
			data := make([]string, 0)
			data = append(data, period)
			data = append(data, number1)
			data = append(data, number2)
			data = append(data, number3)
			data = append(data, number4)
			data = append(data, number5)
			data = append(data, number6)
			data = append(data, number7)
			data = append(data, number8)
			data = append(data, number9)
			data = append(data, number10)

			datas = append(datas, data)
		}
	}

	return datas, nil
}
