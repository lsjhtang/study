package assembler

import (
	"ddd/application/dto"
	"ddd/domain/aggregates"
)

type UserM2d struct {
}

func (u2d *UserM2d) M2D_SimpleUserInfo(member *aggregates.Member) *dto.SimpleUserInfo {
	SimpleUserInfo := &dto.SimpleUserInfo{}
	SimpleUserInfo.ID = member.User.ID
	SimpleUserInfo.UserId = member.User.UserId
	SimpleUserInfo.NickName = member.User.NickName
	SimpleUserInfo.Money = member.User.Extra.Money

	return SimpleUserInfo
}
