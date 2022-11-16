package main

import (
	"fmt"

	"github.com/tabaproc/testify-study/ex3/data"
)

// ビジネスロジック
func execute(id int, repo data.RepositoryOperator) data.Staff {
	staff := repo.Staff(id)

	log("昇給前", staff)

	repo.UpdateIncome(&staff)
	log("昇給1回目", staff)

	repo.UpdateIncome(&staff)
	log("昇給2回目", staff)

	return staff
}

// log関数
func log(title string, staff data.Staff) {
	fmt.Printf("[%s]StaffID: %d, 名前: %s, ポジション: %s, 年収: %s\n",
		title, staff.ID, staff.Name, staff.Post, staff.AnnualIncome())
}

// main関数
func main() {
	repo := data.NewRepository()
	_ = execute(1, repo)
}
