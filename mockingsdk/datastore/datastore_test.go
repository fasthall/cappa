package datastore

import (
	"bytes"
	"fmt"
	"testing"
)

func TestMinioPut(t *testing.T) {
	cli, err := NewDatastore("minio")
	if err != nil {
		t.Errorf("%s", err)
	}
	content := bytes.NewReader([]byte("testcontent"))
	err = cli.Put("go-test", "TestMinioPut", content, "application/octet-stream")
	if err != nil {
		t.Errorf("%s", err)
	}
}

func TestMinioGet(t *testing.T) {
	cli, err := NewDatastore("minio")
	if err != nil {
		t.Errorf("%s", err)
	}
	obj, err := cli.Get("go-test", "TestMinioPut")
	if err != nil {
		t.Errorf("%s", err)
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(obj)
	fmt.Println(buf)
}
