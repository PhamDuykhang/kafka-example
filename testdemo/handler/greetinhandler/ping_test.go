package greetinhandler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/PhamDuyKhang/kafkaexamples/testdemo/handler/greetinhandler/mock"
	"github.com/PhamDuyKhang/kafkaexamples/testdemo/types"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestSayHi(t *testing.T) {
	ctl := gomock.NewController(t)

	srvMock := mock.NewMockGateKeeper(ctl)

	srvMock.EXPECT().GetGreeting().Return(types.Greeting{
		ID:  "1234",
		Mgs: "Hello i'm Mocking",
	})
	expectObj := types.Greeting{
		ID:  "requestID:32303-545-321234",
		Mgs: "Hello i'm Mocking",
	}

	//
	expectJson, err := json.MarshalIndent(expectObj, "", " ")
	if err != nil {
		t.Fail()
	}
	//
	handler := NewGreetingHandler(srvMock)
	req, err := http.NewRequest("POST", "/api/v1/hi", nil)
	if err != nil {
		t.Fail()
	}
	rr := httptest.NewRecorder()
	http.HandlerFunc(handler.SayHi).ServeHTTP(rr, req)

	assert.JSONEq(t, string(expectJson), rr.Body.String(), "two data isn't same")

}
