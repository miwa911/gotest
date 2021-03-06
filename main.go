package main

import (
  "os"
  log "github.com/sirupsen/logrus"
  "gotest/app/routes"
  "gotest/config"
  "gotest/app"
)

func init() {
  // Log as JSON instead of the default ASCII formatter.
  log.SetFormatter(&log.JSONFormatter{})

  // Output to stdout instead of the default stderr
  // Can be any io.Writer, see below for File example
  log.SetOutput(os.Stdout)

  // Only log the warning severity or above.
  //log.SetLevel(log.WarnLevel)
}

func main() {
	cfg, err := config.New("config/config.json")
	if err != nil {
		log.Fatal(err)
	}

	app := app.New(cfg)
	router := routes.NewRouter(app)
	app.Run(router)

}
