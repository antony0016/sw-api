package service

type TestService struct {
	Name string
}

func NewTestService(name string) *TestService {
	return &TestService{
		Name: name,
	}
}

func (t *TestService) Ping() string {
	return "pong"
}

func (t *TestService) SetName(name string) {
	t.Name = name
}

func (t *TestService) GetName() string {
	return t.Name
}
