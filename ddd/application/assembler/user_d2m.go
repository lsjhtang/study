package assembler

import (
	"ddd/application/dto"
	Models "ddd/domain/models"
	"github.com/go-playground/validator/v10"
)

type UserD2M struct {
	V *validator.Validate
}

func (u2d *UserD2M) D2M_SimpleUserInfo(dto *dto.SimpleUserReq) *Models.User {
	err := u2d.V.Struct(dto)
	if err != nil {
		panic(err)
	}
	return Models.NewUser().Set(map[string]interface{}{"ID": dto.ID})
}
