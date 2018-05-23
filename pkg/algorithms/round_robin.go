package algorithms

// RRLinkedList is the base data structure that will store
// information such as which host is currently receiving traffic.
type RRLinkedList struct {
	hosts   []HostObject
	current HostObject
	size    uint
}

// NewRRLinkedList takes a slice of hostnames and creates the circular
// linked list that will be used to load-balance web traffic.
func NewRRLinkedList(hosts []string) *RRLinkedList {
	var hostObjs []HostObject
	// Iterate through the supplied list of hosts and create host objects.
	for _, host := range hosts {
		hostObjs = append(hostObjs, HostObject{
			algorithm:   "rr",
			connections: 0,
			hostname:    host,
			neighbor:    nil,
		})
	}
	// Return a RRLinkedList object to be used by juggle.
	return &RRLinkedList{
		current: hostObjs[0],
		hosts:   hostObjs,
		size:    uint(len(hosts)),
	}
}
