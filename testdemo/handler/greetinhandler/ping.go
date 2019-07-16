package greetinhandler

import (
	"encoding/json"
	"net/http"

	"github.com/PhamDuyKhang/kafkaexamples/testdemo/types"
	"github.com/sirupsen/logrus"
)

type GateKeeper interface {
	GetGreeting() types.Greeting
	SaySimply() types.Greeting
	SayWithName(name string) types.Greeting
}
type GreetingHandler struct {
	service GateKeeper
}

func NewGreetingHandler(anypersion GateKeeper) *GreetingHandler {
	return &GreetingHandler{
		service: anypersion,
	}
}

func (handler *GreetingHandler) Ping(w http.ResponseWriter, r *http.Request) {
	logrus.Info("ping ping")
	data, err := json.MarshalIndent(map[string]string{
		"uri":    r.RequestURI,
		"status": "200",
		"mgs":    "ping ping"}, "", " ")
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
func (handler *GreetingHandler) SayHi(w http.ResponseWriter, r *http.Request) {
	logrus.Info("say hi")
	mgs := handler.service.GetGreeting()
	data, err := json.MarshalIndent(mgs, "", " ")
	if err != nil {
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.Write(data)
}
