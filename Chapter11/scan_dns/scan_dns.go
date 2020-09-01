package main

import (
	"log"
	"net"
	"strconv"
	"strings"
)

var subnetToScan = "10.0.0" // First three octets

func main() {
	activeThreads := 0
	doneChannel := make(chan bool)

	for ip := 0; ip <= 255; ip++ {
		fullIp := subnetToScan + "." + strconv.Itoa(ip)
		go resolve(fullIp, doneChannel)
		activeThreads++
	}

	// Wait for all threads to finish
	for activeThreads > 0 {
		<-doneChannel
		activeThreads--
	}
}

func resolve(ip string, doneChannel chan bool) {
	addresses, err := net.LookupAddr(ip)
	if err == nil {
		log.Printf("%s - %s\n", ip, strings.Join(addresses, ", "))
	}
	doneChannel <- true
}
