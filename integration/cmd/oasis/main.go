package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/pprof"
	"strconv"

	"github.com/Telenav/osrm-backend/integration/service/oasis/entrypoint"
	"github.com/Telenav/osrm-backend/integration/util/appversion"

	"github.com/golang/glog"
)

// main.go for oasis api
func main() {
	flag.Parse()
	appversion.PrintExit()
	defer glog.Flush()
	fmt.Print(flags.listenPort)

	if flags.cpuProfileFile != "" {
		f, err := os.Create(flags.cpuProfileFile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	mux := http.NewServeMux()

	oasisService, err := entrypoint.NewHttpHandler(flags.osrmBackendEndpoint, flags.finderType, flags.tnSearchEndpoint,
		flags.tnSearchAPIKey, flags.tnSearchAPISignature, flags.localDataPath)
	if err != nil {
		glog.Errorf("Failed to create oasis handler due to err %+v.\n", err)
		return
	}

	mux.Handle("/oasis/v1/earliest/", oasisService)

	// listen
	listening := ":" + strconv.Itoa(flags.listenPort)
	glog.Infof("Listening on %s", listening)
	glog.Fatal(http.ListenAndServe(listening, mux))
}
