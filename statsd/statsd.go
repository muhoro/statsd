package statsd

import (
	"github.com/etsy/statsd/examples/go"
)

var Client *statsd.StatsdClient

func Connect(host string, port int) {
	Client = statsd.New(host, port)
}

func Close() {
	Client.Close()
}
