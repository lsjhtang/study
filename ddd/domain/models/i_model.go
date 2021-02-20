package models

import "fmt"

type IModel interface {
	ToString() string
}

type Model struct {
	EntityId   int
	EntityName string
}

func (m *Model) SetName(name string) {
	m.EntityName = name
}

func (m *Model) SetId(id int) {
	m.EntityId = id
}

func (m *Model) ToString() string {
	return fmt.Sprintf("model is:%s, id is %d", m.EntityName, m.EntityId)
}
