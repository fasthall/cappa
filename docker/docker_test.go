package docker

import "testing"

func testCopy(t *testing.T) {
	filename := "./docker_test.go"
	cid := "f1dce161899c8f619875a000cd21a119132e5c0c3bf9dc61d2066f58d7adca4b"
	path := "/docker_test.go"
	Copy(cid, filename, path)
}

func TestCreateWithBinds(t *testing.T) {
	image := "payload"
	binds := []string{"/Users/fasthall:/payload"}
	env := []string{"PAYLOAD=/payload/test"}
	Create(image, binds, env)
}
