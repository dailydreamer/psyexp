package config

import (
  "log"
  "flag"
)

var (
  Port string
)

// InitConfig init the config parameters.
func InitConfig() {
  flag.StringVar(&Port, "port", "8000", "port service listen at")
  flag.Parse()
  log.Println("Config loaded")
}