package app3;

import "github.com/viant/dsc"

type QueryRequest struct {
	SQL string
}


type QueryResponse struct {
	Status string
	Data []map[string]interface{}
}


type Service interface {

	Query(request *QueryRequest) *QueryResponse

}


type service struct {
	manager dsc.Manager
}


func (s *service) Query(request *QueryRequest) *QueryResponse {
	var respone = &QueryResponse{
		Data:make([]map[string]interface{}, 0),
	}

	err := s.manager.ReadAll(&respone.Data, request.SQL, nil, nil)
	if err != nil {
		respone.Status = err.Error()
	} else {
		respone.Status ="ok"
	}
	return respone
}


func NewService(config *dsc.Config) (Service, error) {
	manager, err := dsc.NewManagerFactory().Create(config)
	if err != nil {
		return nil, err
	}
	return &service{manager:manager}, nil
}