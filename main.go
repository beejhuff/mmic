package main

import (
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/Sirupsen/logrus"
	"github.com/eternnoir/mmic/mmic/config"
)

func main() {
	configPtr := flag.String("c", "config.toml", "Config file path.")
	debugPtr := flag.Bool("d", false, "Enable Debug level.")
	flag.Parse()

	if *debugPtr {
		log.Info("Enable debug level.")
		log.SetLevel(log.DebugLevel)
	}
	if len(*configPtr) < 1 {
		log.Error("Config file path must set.Use -h to get some help.")
	}

	var config config.MmicConfig
	if _, err := toml.DecodeFile(*configPtr, &config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%#v \n", config)
}
