package server

import "net/http"

func (s *Server) Router() *http.ServeMux{
	mux := http.NewServeMux() 
	mux.Handle("GET /api/v1/users/", http.HandlerFunc(s.handleGetUser))
	
	return mux 
}