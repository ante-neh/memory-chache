package server

import (
	"log"
	"net/http"
	"github.com/ante-neh/memory-chache/internal/database"
	"github.com/ante-neh/memory-chache/types"
)

type Server struct {
	InfoLogger *log.Logger
	ErrorLogger *log.Logger
	Address string
	Db *database.MemoryDatabase
	Chache *database.Chache
}


func NewServer(infoLogger, errorLogger *log.Logger, address string) *Server{
	return &Server{
		InfoLogger:infoLogger,
		ErrorLogger:errorLogger,
		Address: address,
		Db: &database.MemoryDatabase{Db:make(map[string] types.User)},
		Chache: &database.Chache{Chache: make(map[string] types.User)},
	}
}


func (s *Server) Start() *http.Server{
	return &http.Server{
		Addr:s.Address,
		ErrorLog: s.ErrorLogger,
		Handler:s.Router(),
	}
}