package goddb

import (
	"fmt"
	"net/http"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return Handler{service: service}
}

func (receiver Handler) Save(w http.ResponseWriter, req *http.Request) {
	fmt.Println("save")
	/*
		err := receiver.service.Save(SaveValue{})
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
