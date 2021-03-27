package mapper

import (
	"github.com/Masterminds/squirrel"
)

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (um *UserMapper) GetUserList(id int) *SqlMapper {
	return Mapper(squirrel.Select("*").From("users").Where("id = ?", id).Limit(10).ToSql())
}
