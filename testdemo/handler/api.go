package handler

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/PhamDuyKhang/kafkaexamples/testdemo/handler/greetinhandler"
	"github.com/PhamDuyKhang/kafkaexamples/testdemo/repository"
	"github.com/PhamDuyKhang/kafkaexamples/testdemo/service"

	"github.com/gorilla/mux"
)

const (
	get    = http.MethodGet
	post   = http.MethodPost
	put    = http.MethodPut
	delete = http.MethodDelete
)

type KRPRouter struct {
	RouterName string
	URI        string
	Handler    http.HandlerFunc
	Method     string
}

func InitRoute() *mux.Router {
	//
	greetRepo := repository.NewRepository()
	service := service.NewGreetingService(greetRepo)
	pingHandler := greetinhandler.NewGreetingHandler(service)
	///
	routers := []KRPRouter{
		{
			RouterName: "Ping",
			URI:        "/api/v1/ping",
			Method:     post,
			Handler:    pingHandler.Ping,
		},
		{
			RouterName: "Say Hi",
			URI:        "/api/v1/hi",
			Method:     post,
			Handler:    pingHandler.SayHi,
		},
	}
	r := mux.NewRouter()
	for _, rt := range routers {
		hf := rt.Handler
		logrus.Info(r)
		r.Path(rt.URI).Methods(rt.Method).HandlerFunc(hf)
	}
	return r
}
