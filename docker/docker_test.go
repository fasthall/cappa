package docker

import "testing"

func TestCopy(t *testing.T) {
	filename := "./docker_test.go"
	cid := "f1dce161899c8f619875a000cd21a119132e5c0c3bf9dc61d2066f58d7adca4b"
	path := "/docker_test.go"
	Copy(cid, filename, path)
}
