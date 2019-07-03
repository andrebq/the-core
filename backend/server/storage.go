package server

import (
	"context"
	"github.com/andrebq/the-core/backend/api"
	"github.com/andrebq/the-core/backend/storage"
)

type (
	// Storage exposes the blob storage api
	Storage struct {
		server *storage.Server
	}
)

// NewStorage returns a new Storage server object
func NewStorage(actual *storage.Server) *Storage {
	return &Storage{server: actual}
}

// Put proxy
func (s *Storage) Put(ctx context.Context, content *api.Content) (*api.ContentID, error) {
	id, err := s.server.Put(content.Body)
	if err != nil {
		return nil, err
	}
	return &api.ContentID{
		Alg:     id.Alg,
		HexHash: id.ID,
	}, nil
}

// Get proxy
func (s *Storage) Get(ctx context.Context, contentID *api.ContentID) (*api.Content, error) {
	content, err := s.server.Get(storage.ContentID{
		Alg: contentID.Alg,
		ID:  contentID.HexHash,
	})
	if err != nil {
		return nil, err
	}
	return &api.Content{
		Body: content,
		Tags: nil,
	}, nil
}
