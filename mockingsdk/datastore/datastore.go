package datastore

import (
	"fmt"
	"io"
)

type DatastoreAdapter interface {
	Put(bucket string, object string, content io.Reader, objType string) error
	Get(bucket string, object string) (io.Reader, error)
	//Remove()
}

type BackendNotSupportedError struct{}

func (e *BackendNotSupportedError) Error() string {
	return fmt.Sprintf("Backend not supported")
}

func NewDatastore(backend string) (DatastoreAdapter, error) {
	if backend == "minio" {
		return NewMinio()
	} else {
		return nil, &BackendNotSupportedError{}
	}
}
