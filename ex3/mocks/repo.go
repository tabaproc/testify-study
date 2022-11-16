package mocks

import (
	"github.com/stretchr/testify/mock"
	"github.com/tabaproc/testify-study/ex3/data"
)

type RepoMock struct {
	mock.Mock
}

func (r *RepoMock) Staff(id int) data.Staff {
	args := r.Called(id)
	return args.Get(0).(data.Staff)
}

func (r *RepoMock) UpdateIncome(staff *data.Staff) {
	_ = r.Called(staff)
	staff.Rate += data.UpRate[staff.Post]
}
