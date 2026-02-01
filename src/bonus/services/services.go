package services

import (
	model "bonus/model"
	repository "bonus/repository"
)

type Bonus interface {
	GetInfoAboutUserPrivilege(username string) (model.PrivilegeResponse, error)
	UpdateBonus(username, ticketUID string, price int) (model.PrivilegeInfo, error)
	UpdateBonusBonus(username, ticketUid string, price int) (int, error)
	UpdateBonusDelete(username string, price int) error
}

type Services struct {
	Bonus
}

func NewServices(repo *repository.Repository) *Services {
	return &Services{
		Bonus: NewBonusService(repo.RepoBonus),
	}
}
