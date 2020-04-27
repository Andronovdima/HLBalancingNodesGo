package main

import "github.com/prometheus/client_golang/prometheus"

var (
	RequestCount = prometheus.NewCounter(prometheus.CounterOpts{Name: "RPC",})
	//Cpu = prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{})
	//Cpu = prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{})
)