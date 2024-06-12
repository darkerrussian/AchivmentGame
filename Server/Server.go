package Server

import "net/http"

func Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("its work !"))
}
