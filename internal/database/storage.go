package database
import "github.com/ante-neh/memory-chache/types"

type Databse interface {
	GetUserById(string) *types.User 
}

type MemoryDatabase struct{
	Db map[string]*types.User 
}

func(md * MemoryDatabase) GetUserById(id string) *types.User{
	return nil
}