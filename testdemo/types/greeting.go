package types

type Greeting struct {
	ID   string `json:"id"`
	Time string `json:"create_at"`
	Mgs  string `json:"message"`
}
type Heller interface {
}
