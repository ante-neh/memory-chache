package server

import (
	"net/http"
	"github.com/ante-neh/memory-chache/util"
)

func (s *Server) handleGetUser(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id") 
	
	user, ok := s.Chache.GetUserById(id) 

	if !ok {
		util.ResponseWithJson(w, 200, user)
	}

	user, ok = s.Db.GetUserById(id)

	if ok{
		util.ResponseWithError(w, 404, "user not found")
	}

	s.Chache.AddUser(id, *user)
	util.ResponseWithJson(w, 200, user)
}