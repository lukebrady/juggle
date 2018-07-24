package main

// RRLinkedList is the base data structure that will store
// information such as which host is currently receiving traffic.
type RRLinkedList struct {
	hosts   []string
	current int
	size    int
}

// NewRRLinkedList takes a slice of hostnames and creates the circular
// linked list that will be used to load-balance web traffic.
func NewRRLinkedList(hosts []string) *RRLinkedList {
	// Return a RRLinkedList object to be used by juggle.
	return &RRLinkedList{
		current: 0,
		hosts:   hosts,
		size:    len(hosts),
	}
}

// NextHost increments the host for the round robin.
func (l *RRLinkedList) NextHost() {
	// Check to see if the list has been iterated through.
	if l.current < l.size-1 {
		// Increment the current index to the next host in the slice.
		l.current++
	} else {
		// If the list has been iterated through, go back to 0.
		l.current = 0
	}
}
