package data

import "strconv"

// Staff スタッフデータ
type Staff struct {
	ID   int
	Name string
	Post string
	Rate float64
}

func (r *Staff) AnnualIncome() string {
	annualIncome := strconv.FormatFloat(r.Rate, 'f', 1, 64)
	return annualIncome + "BTC"
}

var StaffData = map[int]map[string]interface{}{
	1: {"name": "山田", "post": "Leader", "rate": 3.1},
	2: {"name": "佐藤", "post": "Member", "rate": 2.8},
	3: {"name": "鈴木", "post": "Member", "rate": 2.9},
	4: {"name": "おじさん", "post": "Manager", "rate": 5.0},
}

var UpRate = map[string]float64{
	"Member":  0.1,
	"Leader":  0.2,
	"Manager": 0.3,
}
