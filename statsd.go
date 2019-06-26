package main

import (
	"github.com/etsy/statsd/examples/go"
)

type Statsd struct{
	Client *statsd.StatsdClient
}

func (s Statsd) Connect(host string, port int){
	client := statsd.New(host, port)
	s.Client = client
}

func (s Statsd) Close(){
	s.Client.Close()
}

var Client *statsd.StatsdClient 

func connectToStatsD(host string, port int){
	client := statsd.New(host, port)
	Client = client
}
