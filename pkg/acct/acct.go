package acct

// Service ...
type Service struct {
	repo Repository
}

// CreateRequest ...
type CreateRequest struct {
	Name     string
	Password string
}

// Repository ...
type Repository interface {
	CreateAccount(req *CreateRequest) error
	ResetPassword(name string) error
}

// New ...
func New(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// Register ..
func (s *Service) Register(req *CreateRequest) error {
	return s.repo.CreateAccount(req)
}

// ResetPassword ...
func (s *Service) ResetPassword(name string) error {
	return s.repo.ResetPassword(name)
}
