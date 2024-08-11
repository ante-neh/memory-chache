package database

import "github.com/ante-neh/memory-chache/types"

type MemoryStorage interface {
	GetUserById(string) types.User
}

type Chache struct {
	Chache map[string]*types.User
}

func (c *Chache) GetUserById(id string) *types.User{
	return nil 
}