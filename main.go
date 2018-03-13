package main

import(
	"flag"

	"github.com/juju/errors"
	"github.com/double-chosen-will-server/config"
	log "github.com/sirupsen/logrus"

	"github.com/double-chosen-will-server/logutil"
)

var(
	configPath *string = flag.String("config", "", "configuration file path")
)

var(
	cfg *config.Config
)

func main(){
	flag.Parse()

	setupConfig()
	setupLog()
	log.Info("Server started")
	log.Info("shutting down")
}

func setupConfig() {
	cfg = config.NewConfig()
	if *configPath == ""{
		log.Infof("no configuration file found, use default configuration")
		return;
	}
	if err := cfg.LoadConfig(*configPath); err != nil {
		log.Warnf(errors.ErrorStack(err))
	}

}

func setupLog(){
	if err := logutil.InitLogger(cfg.ToLogConfig()); err != nil {
		log.Infof("an error happend when setting up logger: %s", errors.Trace(err))
	}
}

func init(){
}