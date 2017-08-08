package app2_test

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/viant/dsc"

	"github.com/viant/dsunit"
	"github.com/adranwit/go-kickstart-from-java/app2"
	"github.com/stretchr/testify/assert"
	"github.com/viant/toolbox"
	"time"
)





func TestRead(t *testing.T) {
	dsunit.InitDatastoreFromURL(t, "test://test/datastore_init.json")
	dsunit.ExecuteScriptFromURL(t, "test://test/script_request.json")
	dsunit.PrepareDatastoreFor(t, "mytestdb2", "test://test/", "Read")

	configUrl := dsunit.ExpandTestProtocolAsURLIfNeeded("test://test/config/store.json")
	service, err := app2.NewServiceFromURL(configUrl)
	assert.Nil(t, err)


	response := service.Read(&app2.ReadRequest{
		Where:"id = ?",
		Parameters:[]interface{} {
			1,
		},
		OrderBy:[]string{
			"id",
		},
	})


	assert.Equal(t, response.Status, "ok")
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "Abc", response.Data[0].Name)


	server := app2.NewServer("8882", service)
	go server.Start()
	time.Sleep(2 * time.Second)

	var result = &app2.ReadResponse{}


	{

		err := toolbox.RouteToService("post", "http://127.0.0.1:8882/v1/site/read/", &app2.PersistRequest{

		}, &result)
		if err != nil {
			t.Errorf("Failed to send get request  %v", err)
		}
		assert.Equal(t, response.Status, "ok")
		assert.Equal(t, 3, len(result.Data))
		assert.Equal(t, "Abc", result.Data[0].Name)

	}

}




func TestPersist(t *testing.T) {

	dsunit.InitDatastoreFromURL(t, "test://test/datastore_init.json")

	dsunit.ExecuteScriptFromURL(t, "test://test/script_request.json")
	dsunit.PrepareDatastoreFor(t, "mytestdb2", "test://test/", "Persist")

	configUrl := dsunit.ExpandTestProtocolAsURLIfNeeded("test://test/config/store.json")
	service, err := app2.NewServiceFromURL(configUrl)
	assert.Nil(t, err)



	var category = "Travel";
	response := service.Persist(&app2.PersistRequest{
		Data: &app2.Site{
			Name:     "AZZZZZZ",
			Url:      "bbc.com/dferewrew",
			Category: &category,
		},
	})


	assert.Equal(t, response.Status, "ok")
	assert.Equal(t, 1, len(response.Data))
	assert.Equal(t, "AZZZZZZ", response.Data[0].Name)
	assert.Equal(t, 3, response.Data[0].Id)

	dsunit.ExpectDatasetFor(t, "mytestdb2", dsunit.SnapshotDatasetCheckPolicy, "test://test/", "Persist")





	server := app2.NewServer("8882", service)
	go server.Start()
	time.Sleep(2 * time.Second)

	var result = &app2.ReadResponse{}


	{


		var category = "Finance"
		err := toolbox.RouteToService("post", "http://127.0.0.1:8882/v1/site/persist/", &app2.PersistRequest{
			Data:&app2.Site{
				Id:       2,
				Name:     "Cde------",
				Url:      "http://eee.com",
				Category: &category,
			},
		}, &result)
		if err != nil {
			t.Errorf("Failed to send get request  %v", err)
		}
		assert.Equal(t, response.Status, "ok")
		assert.Equal(t, 1, len(result.Data))
		assert.Equal(t, "Cde------", result.Data[0].Name)

	}

}


