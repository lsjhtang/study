package service

import (
	"ddd/application/assembler"
	"ddd/application/dto"
	"ddd/domain/models"
)

type UserService struct {
	assUserReq *assembler.UserD2M
	assUserRsp *assembler.UserM2d
}

func (us *UserService) GetSimpleUserInfo(dto *dto.SimpleUserReq) *models.User {
	user := us.assUserReq.D2M_SimpleUserInfo(dto)
	return user
}
