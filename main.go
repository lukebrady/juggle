package main

import (
	"fmt"
	"hash/fnv"
)

// Endpoint is a node endpoint object representing a machine
// within a production envrionment. Juggle will load-balance
// between an array of these endpoints.
type Endpoint struct {
	ipAddr  string
	dnsName string
}

// Endpoints holds a slice of Endpoint objects.
type Endpoints []Endpoint

// Juggler is the main load-balancing object that will distribute
// traffic using LRU and Hash based algorithms.
type Juggler struct {
	endpoints *Endpoints
}

// NewJuggler creates a new Juggler object.
func NewJuggler(endpoints *Endpoints) *Juggler {
	return &Juggler{
		endpoints: endpoints,
	}
}

func (j *Juggler) hashEndpointLocation(endpoint *Endpoint) uint32 {
	// Create new hash for the endpoint.
	hash := fnv.New32()
	// Write the endpoints data to the hash.
	hash.Write([]byte(endpoint.ipAddr))
	hash.Write([]byte(endpoint.dnsName))
	// Sum the hash and return the modulus value.
	return (hash.Sum32() % uint32(len(*j.endpoints)))
}

func main() {
	var ends Endpoints
	end1 := Endpoint{ipAddr: "192.168.1.193", dnsName: "hello.lukebrains.com"}
	end2 := Endpoint{ipAddr: "192.168.1.195", dnsName: "testsite123.com"}
	ends = append(ends, end1)
	ends = append(ends, end2)
	j := NewJuggler(&ends)
	fmt.Println(ends[j.hashEndpointLocation(&end2)])
	fmt.Println(ends[j.hashEndpointLocation(&end1)])
}
