package mapper

import "github.com/Masterminds/squirrel"

type UserMapper struct {
}

func NewUserMapper() *UserMapper {
	return &UserMapper{}
}

func (um *UserMapper) GetUserList() *SqlMapper {
	return Mapper(squirrel.Select("*").From("users").Limit(10).ToSql())
}
