package main

import (

	_ "github.com/viant/bgc"
	"flag"
	"log"
	"github.com/viant/toolbox"
	"github.com/viant/dsc"
	"github.com/adranwit/go-kickstart-from-java/app3"
)


//go run server.go --configUrl=file:///Users/awitas/go/src/github.com/adranwit/go-kickstart-from-java/app3/main/config.json
var configUrl = flag.String("configUrl", "configUrl", "URL for cce.Config json resource")



func main() {
	flag.Parse()

	*configUrl = "file:///Users/awitas/go/src/github.com/adranwit/go-kickstart-from-java/app3/main/config.json"

	config := &dsc.Config{}
	err := toolbox.LoadConfigFromUrl(*configUrl, config)
	if err != nil {
		log.Fatal(err)
	}
	service, err := app3.NewService(config)
	if err != nil {
		log.Fatal(err)
	}

	server :=  app3.NewServer("8083", service)
	server.Start()
}

/*

{

           "SQL":"SELECT request, userProfile FROM `viant-adelphic.adlogs_us.ad_impressions_20170808` LIMIT 10"
}


 */