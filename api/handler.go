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

func (receiver Handler) Put(w http.ResponseWriter, req *http.Request) {
	requestBody, err := ioutil.ReadAll(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": fmt.Sprintf("error while reading request body %v", err.Error()),
		})
		return
	}

	var saveRequest goddb.SaveValue
	err = json.Unmarshal(requestBody, &saveRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": fmt.Sprintf("error while parsing request body %v", err.Error()),
		})
		return
	}

	err = receiver.service.Put(saveRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": fmt.Sprintf("error while saving key:value %v", err.Error()),
		})
		return
	}
	w.WriteHeader(201)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": "key:value saved",
	})
}

func (receiver Handler) Retrieve(w http.ResponseWriter, req *http.Request) {
	retrieveRequest := goddb.RetrieveValue{Key: req.URL.Query().Get("key")}
	err, value := receiver.service.Retrieve(retrieveRequest)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": fmt.Sprintf("error while retriving key:value %v", err.Error()),
		})
		return
	}
	if len(value) < 1 {
		w.WriteHeader(http.StatusInternalServerError)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"success": false,
			"message": fmt.Sprintf("empty value returned while retriving key:value"),
		})
		return
	}
	w.WriteHeader(200)
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"message": value,
	})
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
