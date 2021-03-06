package main

import (
	"fmt"
	"github.com/Cepave/open-falcon-backend/common/logruslog"
	"github.com/Cepave/open-falcon-backend/common/vipercfg"
	"os"

	"github.com/Cepave/open-falcon-backend/modules/nodata/collector"
	"github.com/Cepave/open-falcon-backend/modules/nodata/config"
	"github.com/Cepave/open-falcon-backend/modules/nodata/g"
	"github.com/Cepave/open-falcon-backend/modules/nodata/http"
	"github.com/Cepave/open-falcon-backend/modules/nodata/judge"
	"github.com/Cepave/open-falcon-backend/modules/nodata/sender"
)

func main() {
	vipercfg.Parse()
	vipercfg.Bind()

	if vipercfg.Config().GetBool("version") {
		fmt.Println(g.VERSION)
		os.Exit(0)
	}
	if vipercfg.Config().GetBool("vg") {
		fmt.Println(g.VERSION, g.COMMIT)
		os.Exit(0)
	}

	// global config
	vipercfg.Load()
	g.ParseConfig(vipercfg.Config().GetString("config"))
	logruslog.Init()
	// proc
	g.StartProc()

	// config
	config.Start()
	// collector
	collector.Start()
	// judge
	judge.Start()
	// sender
	sender.Start()

	// http
	http.Start()

	select {}
}
