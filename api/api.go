package api

import (
	"encoding/json"
	"fmt"
	"goddb/config"
	"net/http"
	"time"
)

type Api struct {
	config  *config.Config
	handler Handler
	routes  []Route
}

func NewApi(config *config.Config, handler Handler) *Api {
	api := &Api{
		config:  config,
		handler: handler,
	}

	api.routes = append(api.routes, Route{
		Path:    "/record",
		Method:  "POST",
		Handler: api.handler.Put,
	})
	api.routes = append(api.routes, Route{
		Path:    "/record",
		Method:  "GET",
		Handler: api.handler.Retrieve,
	})
	api.routes = append(api.routes, Route{
		Path:    "/record/flush",
		Method:  "GET",
		Handler: api.handler.Flush,
	})
	api.routes = append(api.routes, Route{
		Path:    "/info",
		Method:  "GET",
		Handler: api.Info,
	})

	http.HandleFunc("/", api.routingMiddleware)

	return api
}

func (receiver *Api) Start() error {
	fmt.Printf("%v listening on port %d\n", time.Now(), receiver.config.Server.Port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", receiver.config.Server.Port), nil)
	return fmt.Errorf("error while serving http: %v", err)
}

func (receiver *Api) Info(w http.ResponseWriter, _ *http.Request) {
	_ = json.NewEncoder(w).Encode(map[string]string{
		"version":     "0.0.1",
		"description": "I am storing",
		"name":        "goddb storage api",
	})
}

func (receiver *Api) routingMiddleware(w http.ResponseWriter, r *http.Request) {
	contentTypeMiddleware(w, r)

	path := r.URL.Path
	method := r.Method
	fmt.Printf("incomming request path: %s, method: %s \n", path, method)
	for _, route := range receiver.routes {
		if route.Path == path && route.Method == method {
			loggingMiddleware(route, w, r)
		}
	}
}

func contentTypeMiddleware(w http.ResponseWriter, _ *http.Request) {
	w.Header().Add("Content-Type", "application/json")
}

func loggingMiddleware(route Route, w http.ResponseWriter, r *http.Request) {
	startTime := time.Now()
	fmt.Printf("%v request started path: %s, method: %s\n", time.Now(), route.Path, route.Method)
	route.Handler(w, r)
	elapsed := time.Since(startTime)
	fmt.Printf("%v request finished execution time: %s path: %s, method: %s\n", time.Now(), elapsed, route.Path, route.Method)
}
