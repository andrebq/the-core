package storage

import (
	"github.com/pkg/errors"
)

type (
	// Server holds the API to read/write to the blob storage
	Server struct {
		driver Driver
	}

	// Driver implements the interface used by a Server to read/write to the underlying storage
	Driver interface {
		// Put the given content under the given ID
		Put(id ContentID, content Content) error

		// Get the content under the given ID
		Get(id ContentID) (Content, error)

		// Exists returns true if the id is present within the driver
		Exists(id ContentID) (bool, error)
	}

	// Option for a given server
	Option interface {
		// Set the given option to the given server
		Set(s *Server) error
	}

	optionFn func(s *Server) error
)

// New returns a server with the given options
func New(opts ...Option) (*Server, error) {
	s := &Server{}
	for _, o := range opts {
		err := o.Set(s)
		if err != nil {
			return nil, err
		}
	}
	if s.driver == nil {
		return nil, MissingDriverErr
	}
	return s, nil
}

// WithDriver configures the server to use the given driver
func WithDriver(d Driver) Option {
	if d == nil {
		panic("driver cannot be nil")
	}
	return optionFn(func(s *Server) error {
		s.driver = d
		return nil
	})
}

// Put the given content in the server and returns the calculated ID
func (s *Server) Put(c Content) (ContentID, error) {
	id := c.GetID()
	ok, err := s.driver.Exists(id)
	if err != nil {
		return ContentID{}, err
	}
	if ok {
		return id, nil
	}
	return id, s.driver.Put(id, c)
}

// Get the content under the given ID or returns an error
func (s *Server) Get(id ContentID) (Content, error) {
	c, err := s.driver.Get(id)
	if err != nil && err == ContentNotFoundErr {
		return nil, err
	} else if err != nil {
		return nil, errors.Wrap(err, "storage: unable to find content by id")
	}
	return c, nil
}

func (f optionFn) Set(s *Server) error {
	return f(s)
}
