package storage

type (
	// Err type
	Err string
)

const (
	// ContentNotFoundErr is returned when the driver couldn't find the
	// content for a given ID.
	//
	// Used mostly by storage drivers to communicate with storage servers
	ContentNotFoundErr = Err("content id not found")

	// MissingDriverErr indicates that a Server was configured but it is missing
	// a storage driver
	MissingDriverErr = Err("missing storage driver")
)

// Implements error interface
func (e Err) Error() string {
	return string(e)
}
