package api

import (
	"encoding/json"
	"fmt"
	"goddb/goddb"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	service goddb.Service
}

func NewHandler(service goddb.Service) Handler {
	return Handler{service: service}
}

// TODO: fix error hanling
func (receiver Handler) Put(w http.ResponseWriter, req *http.Request) {
	fmt.Println("save")
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var saveRequest goddb.SaveValue
	err = json.Unmarshal(requestBody, &saveRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = receiver.service.Put(saveRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": err.Error(),
		})
		return
	}
	w.WriteHeader(201)
	err = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"result":  "xd",
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
