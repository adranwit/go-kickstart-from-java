package app2

import (
	"github.com/viant/dsc"
	"github.com/viant/toolbox"
	"strings"
)


const selectSites = "SELECT id, url, name, category, status FROM site"


type ReadRequest struct {
	Where string
	Parameters []interface{}
	OrderBy []string
}


type ReadResponse struct {
	Status string
	Data []*Site
}


type PersistRequest struct {
	Data *Site
}



type Service interface {

	Read(request *ReadRequest) *ReadResponse

	Persist(requet *PersistRequest) *ReadResponse

}


type service struct {
	manager dsc.Manager
}


func (s *service) Read(request *ReadRequest) *ReadResponse {

	var response = &ReadResponse{
		Data:make([]*Site, 0),
	}

	var sql = selectSites
	if request.Where != "" {
		sql += " WHERE " + request.Where
	}

	if len(request.OrderBy) > 0 {
		sql += " ORDER  BY  " + strings.Join(request.OrderBy, ",")
	}


 	err := s.manager.ReadAll(&response.Data, sql, request.Parameters, nil)
	if err != nil {
		response.Status = err.Error()
	} else {
		response.Status = "ok"
	}

	return response;
}



func (s *service) Persist(request *PersistRequest) *ReadResponse {
	var response = &ReadResponse{
		Data:make([]*Site, 0),
	}

	_, _, err := s.manager.PersistSingle(request.Data, "site", nil)
	if err != nil {
		response.Status  = err.Error()
	} else {
		response.Status  = "ok"
	}

	response.Data = append(response.Data, request.Data)
	return response
}



func NewService(config *dsc.Config) (Service, error) {
	manager, err := dsc.NewManagerFactory().Create(config)
	if err != nil {
		return  nil, err
	}
	return &service{
		manager:manager,
	}, nil
}


//NewServiceFromURL creates a new service
func NewServiceFromURL(configURL string) (Service, error) {
	config := &dsc.Config{}
	err := toolbox.LoadConfigFromUrl(configURL, config);
	if err != nil {
		return nil, err
	}
	return NewService(config)
}