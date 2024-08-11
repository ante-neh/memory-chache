package database

import "github.com/ante-neh/memory-chache/types"

type MemoryStorage interface {
	GetUserById(string) types.User
}

type chache struct {
	chache map[string]string
}

func (c *chache) GetUserById(id string) *types.User{
	return nil 
}