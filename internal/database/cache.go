package database

import (
	"sync"
	"github.com/ante-neh/memory-chache/types"
)

type MemoryStorage interface {
	GetUserById(string) (types.User, bool)
	AddUser(string, types.User) error
}

var mu sync.RWMutex
type Chache struct {
	Chache map[string]types.User
}

func (c *Chache) GetUserById(id string) (*types.User, bool){
	//make it concurrent safe
	mu.RLock()
	defer mu.RUnlock()

	//get the user
	user, ok := c.Chache[id] 
	return &user, ok
}

func (c *Chache) AddUser(id string, user types.User) error{
	//make it concurrent safe 
	mu.Lock()
	defer mu.Unlock() 

	//store the user
	c.Chache[id] = user 

	return nil 
}