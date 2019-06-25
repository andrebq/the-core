package storage

import (
	"reflect"
	"testing"
)

type (
	inMemoryStorage map[ContentID]Content
)

func (d inMemoryStorage) Put(id ContentID, c Content) error {
	d[id] = c
	return nil
}

func (d inMemoryStorage) Get(id ContentID) (Content, error) {
	c, ok := d[id]
	if !ok {
		return nil, ContentNotFoundErr
	}
	return c, nil
}

func (d inMemoryStorage) Exists(id ContentID) (bool, error) {
	_, ok := d[id]
	return ok, nil
}

func TestServer(t *testing.T) {
	server, err := New(WithDriver(make(inMemoryStorage)))
	if err != nil {
		t.Fatal(err)
	}
	content := Content([]byte("hello world"))
	id, err := server.Put(content)
	if err != nil {
		t.Fatal(err)
	}

	if id != content.GetID() {
		t.Fatal("server id differs from expected content id")
	}

	contentFromServer, err := server.Get(id)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(content, contentFromServer) {
		t.Fatal("stored content differs from retreived content")
	}
	if id != contentFromServer.GetID() {
		t.Fatal("computed ids differs between stored/retreived content")
	}
}
