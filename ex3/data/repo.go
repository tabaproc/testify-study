package data

// RepositoryOperator インターフェイス
type RepositoryOperator interface {
	Staff(id int) Staff
	UpdateIncome(staff *Staff)
}

// repository
type repository struct {
}

// NewRepository repositoryコンストラクタ
func NewRepository() RepositoryOperator {
	return &repository{}
}

// Staff 顧客情報を取得
func (r *repository) Staff(id int) Staff {
	return Staff{
		ID:   id,
		Name: StaffData[id]["name"].(string),
		Post: StaffData[id]["post"].(string),
		Rate: StaffData[id]["rate"].(float64),
	}
}

func (r *repository) UpdateIncome(staff *Staff) {
	staff.Rate += UpRate[staff.Post]
}
