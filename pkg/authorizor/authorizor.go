package authorizor

// Service ...
type Service struct {
	repo Repository

}

type Repository interface{}

// New ...
func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Register ..
func (s *Service) Register() {}
