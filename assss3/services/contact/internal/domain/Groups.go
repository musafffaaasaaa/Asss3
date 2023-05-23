package domain

import (
	"errors"
	"unicode/utf8"
)

type Group struct {
	id   int64
	name string
}

func (g *Group) SetName(name string) error {
	if utf8.RuneCountInString(name) > 250 {
		return errors.New("name should be less than 250 symbols")
	}
	g.name = name
	return nil
}

func (g *Group) GetId() int64 {
	return g.id
}

func (g *Group) GetName() string {
	return g.name
}

func NewGroup(id int64, name string) (*Group, error) {
	group := &Group{}
	err := group.SetName(name)
	if err != nil {
		return nil, err
	}
	group.id = id
	return group, nil
}
