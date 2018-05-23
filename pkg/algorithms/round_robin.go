package algorithms

// CircLinkedList is the base data structure that will store
// information such as which host is currently receiving traffic.
type CircLinkedList struct {
	current *HostObject
	size    uint
}
