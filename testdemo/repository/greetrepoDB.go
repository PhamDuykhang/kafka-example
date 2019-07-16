package repository

import "github.com/PhamDuyKhang/kafkaexamples/testdemo/types"

type Repository struct {
	// place to define a necessary connection
}

func NewRepository() *Repository {
	return &Repository{}
}

func (r *Repository) FindWordByName(name string) types.Greeting {
	return types.Greeting{
		Mgs: "the repository need to implemnet",
	}
}
func (r *Repository) GetAllSentence() []types.Greeting {
	return []types.Greeting{
		{Mgs: "the repository need to implemnet"},
	}
}
func (r *Repository) GetOne() types.Greeting {
	return types.Greeting{
		Mgs: "the repository need to implemnet",
	}
}
