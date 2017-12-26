//package test

//type data struct {
//	date  string
//	value float64
//}

//var datas []data

//func main() {
//	datas = make([]data, 0)

//	datas = append(datas, data{"20170926", 10.32})
//	datas = append(datas, data{"20170927", 10.35})
//	datas = append(datas, data{"20170928", 10.38})
//	datas = append(datas, data{"20170929", 10.49})
//	datas = append(datas, data{"20170930", 10.72})
//	datas = append(datas, data{"20171010", 10.69})
//	datas = append(datas, data{"20171011", 10.75})
//	datas = append(datas, data{"20171012", 10.72})
//	datas = append(datas, data{"20171013", 10.69})
//	datas = append(datas, data{"20171016", 10.7})
//	datas = append(datas, data{"20171017", 10.45})
//	datas = append(datas, data{"20171018", 10.27})
//	datas = append(datas, data{"20171019", 10.11})
//	datas = append(datas, data{"20171020", 10.13})
//	datas = append(datas, data{"20171023", 10.29})
//	datas = append(datas, data{"20171024", 10.12})
//	datas = append(datas, data{"20171025", 10.15})
//	datas = append(datas, data{"20171026", 10.05})
//	datas = append(datas, data{"20171027", 10.18})
//	datas = append(datas, data{"20171030", 9.77})
//	datas = append(datas, data{"20171031", 9.53})
//	datas = append(datas, data{"20171101", 9.63})
//	datas = append(datas, data{"20171102", 9.67})
//	datas = append(datas, data{"20171103", 9.53})
//	datas = append(datas, data{"20171106", 9.23})
//	datas = append(datas, data{"20171107", 9.42})
//	datas = append(datas, data{"20171108", 9.42})
//	datas = append(datas, data{"20171109", 9.5})
//	datas = append(datas, data{"20171110", 9.48})
//	datas = append(datas, data{"20171113", 9.39})
//	datas = append(datas, data{"20171114", 9.35})
//	datas = append(datas, data{"20171115", 9.29})
//	datas = append(datas, data{"20171116", 9.31})
//	datas = append(datas, data{"20171117", 9.14})
//	datas = append(datas, data{"20171120", 8.47})
//	datas = append(datas, data{"20171121", 8.69})
//	datas = append(datas, data{"20171122", 8.54})
//	datas = append(datas, data{"20171123", 8.26})
//	datas = append(datas, data{"20171124", 8.15})
//	datas = append(datas, data{"20171127", 8.24})
//	datas = append(datas, data{"20171128", 8.37})
//	datas = append(datas, data{"20171129", 8.25})
//	datas = append(datas, data{"20171130", 8.23})
//	datas = append(datas, data{"20171201", 8.21})
//	datas = append(datas, data{"20171204", 8.2})
//	datas = append(datas, data{"20171205", 7.77})
//	datas = append(datas, data{"20171206", 7.81})
//	datas = append(datas, data{"20171207", 7.91})
//	datas = append(datas, data{"20171208", 8.42})
//	datas = append(datas, data{"20171211", 8.66})
//	datas = append(datas, data{"20171212", 8.55})
//	datas = append(datas, data{"20171213", 8.56})
//	datas = append(datas, data{"20171214", 8.48})
//	datas = append(datas, data{"20171215", 8.34})
//	datas = append(datas, data{"20171218", 8.08})
//	datas = append(datas, data{"20171219", 8.07})
//	datas = append(datas, data{"20171220", 8.16})
//	datas = append(datas, data{"20171221", 7.95})
//	datas = append(datas, data{"20171222", 8.03})
//	datas = append(datas, data{"20171225", 7.85})

//}

package test

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(r.Intn(100))
	}
}
