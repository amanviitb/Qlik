package handlers

import "net/http"

func (s *server) handleGetSingle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Messages Exist"))
	}
}
