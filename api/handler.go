package api

import (
	"fmt"
	"goddb/goddb"
	"net/http"
)

type Handler struct {
	service goddb.Service
}

func NewHandler(service goddb.Service) Handler {
	return Handler{service: service}
}

func (receiver Handler) Put(w http.ResponseWriter, req *http.Request) {
	fmt.Println("save")
	/*
		err := receiver.service.Put(SaveValue{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
}

func (receiver Handler) Retrieve(w http.ResponseWriter, req *http.Request) {
	fmt.Println("retrieve")
	/*
		err, _ := receiver.service.Retrieve(RetrieveValue{})
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
}

func (receiver Handler) Flush(w http.ResponseWriter, req *http.Request) {
	fmt.Println("flush")
	/*
		err := receiver.service.Flush()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	*/
}
