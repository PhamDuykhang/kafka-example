package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	repo_mock "github.com/PhamDuyKhang/kafkaexamples/testdemo/service/mock"
	"github.com/PhamDuyKhang/kafkaexamples/testdemo/types"

	"github.com/golang/mock/gomock"
)

func TestGetGreeting(T *testing.T) {
	// create controller to track the test result
	controller := gomock.NewController(T)
	greetingRepoMock := repo_mock.NewMockGreetRepository(controller)
	greetingRepoMock.EXPECT().GetOne().Return(types.Greeting{
		ID:  "123",
		Mgs: "Hello I'm mocking",
	}).Times(1)
	expectID := "requestID:32303-545-32"
	greetingSvr := NewGreetingService(greetingRepoMock)
	gotData := greetingSvr.GetGreeting()

	assert.Equal(T, expectID, gotData.ID, "two id isn't same")
}
