package config

import (
	"os"
	"time"
)

var (
	defaultAmqpURI                   = "amqp://"
	defaultPoolSize                  = 1
	defaultProviderName              = "docker"
	defaultHardTimeout, _            = time.ParseDuration("50m")
	defaultLogTimeout, _             = time.ParseDuration("10m")
	defaultBuildCacheFetchTimeout, _ = time.ParseDuration("5m")
	defaultBuildCachePushTimeout, _  = time.ParseDuration("5m")
	defaultHostname, _               = os.Hostname()
)