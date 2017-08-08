package app2

import (
	"testing"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/viant/dsc"

	"github.com/viant/dsunit"
)

func TestRead(t *testing.T) {
	dsunit.InitDatastoreFromURL(t, "test://test/datastore_init.json")
	dsunit.ExecuteScriptFromURL(t, "test://test/script_request.json")
	dsunit.PrepareDatastoreFor(t, "mytestdb", "test://test/", "Read")




	dsunit.ExpectDatasetFor(t, "mytestdb", dsunit.SnapshotDatasetCheckPolicy, "test://test/", "Read")
}


