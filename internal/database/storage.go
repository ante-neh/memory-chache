package database
import "github.com/ante-neh/memory-chache/types"

type Databse interface {
	GetUserById(string) *types.User 
}

type database struct{
	db map[string]*types.User 
}

func(d * database) GetUserById(id string) *types.User{
	return nil
}