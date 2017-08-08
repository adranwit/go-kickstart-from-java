package app1

type Service interface {
	Reverse(values []int) []int
}

type service struct{}

func (s *service) Reverse(values []int) []int {
	var result = make([]int, 0)
	for i := len(values) - 1; i >= 0; i-- {
		result = append(result, values[i])
	}
	return result
}

func NewService() Service {
	return &service{}
}
