package mockkkk

import (
	"testing"

	"go-backoffice/src/github.com/stretchr/testify/assert"

	"github.com/PhamDuyKhang/kafkaexamples/testdemo/mockkkk/mock"

	"github.com/golang/mock/gomock"
)

func TestSayHi(t *testing.T) {
	ctl := gomock.NewController(t)
	db_mocking := mock.NewMockDB(ctl)

	db_mocking.EXPECT().Hii().Return("Hi3")

	svr := NewService(db_mocking)

	expectString := "Hi"

	gotString := svr.SayHi()

	assert.Equal(t, expectString, gotString, "not same")
}
