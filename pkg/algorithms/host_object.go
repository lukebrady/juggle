package algorithms

// HostObject stores data about a host and it's connections
// and other useful meta-data.
type HostObject struct {
	hostname    string
	connections uint
	algorithm   string
	neighbor    *HostObject
}
