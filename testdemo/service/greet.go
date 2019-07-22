package service

import (
	"fmt"
	"time"

	"github.com/PhamDuyKhang/kafkaexamples/testdemo/types"
)

type GreetRepository interface {
	FindWordByName(name string) types.Greeting
	GetAllSentence() []types.Greeting
	GetOne() types.Greeting
}
type GreetingService struct {
	repo GreetRepository
}

func NewGreetingService(r GreetRepository) *GreetingService {
	return &GreetingService{
		repo: r,
	}
}
func (srv GreetingService) GetGreeting() types.Greeting {
	decorationString := "requestID:32303-545-32"
	st := srv.repo.GetOne()
	st.ID = st.ID + decorationString
	t := time.Now()
	st.Time = t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	return st
}
func (srv GreetingService) SaySimply() types.Greeting {
	st := srv.repo.GetOne()
	t := time.Now()
	st.Time = t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	return st
}
func (srv GreetingService) SayWithName(name string) types.Greeting {
	//some action to transform the name
	nameWithUUID := fmt.Sprintf("cis-id:2335334-232-53-%s", name)
	g := srv.repo.FindWordByName(nameWithUUID)
	t := time.Now()
	g.Time = t.Format("Mon Jan 2 15:04:05 -0700 MST 2006")
	return g
}
