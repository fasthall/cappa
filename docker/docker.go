package docker

import (
	"bufio"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/reference"
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"io/ioutil"
	"os"
)

var cli *client.Client

func init() {
	host := "unix:///var/run/docker.sock"
	fmt.Printf("Trying to connect to Docker daemon at %s...\n", host)
	defaultHeaders := map[string]string{"User-Agent": "engine-api-cli-1.0"}
	var err error
	cli, err = client.NewClient(host, "v1.22", nil, defaultHeaders)
	if err != nil {
		panic(err)
	}
	fmt.Println("Docker daemon connected.")
}

func List() []string {
	options := types.ContainerListOptions{All: true}
	containers, err := cli.ContainerList(context.Background(), options)
	if err != nil {
		panic(err)
	}
	list := make([]string, len(containers))
	for i, c := range containers {
		list[i] = c.ID
	}
	return list
}

func Pull(image string) {
	ref, tag, err := reference.Parse(image)
	fmt.Println("Pull", ref, tag, err)
	resp, err := cli.ImagePull(context.Background(), ref, types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	body, err := ioutil.ReadAll(resp)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(body))
}

func Remove(image string) {
	ref, tag, err := reference.Parse(image)
	fmt.Println("Remove", ref, tag, err)
	cli.ImageRemove(context.Background(), ref, types.ImageRemoveOptions{})
}

func Exist(image string) bool {
	ref, tag, err := reference.Parse(image)
	fmt.Println("Exist", ref, tag, err)
	_, _, err = cli.ImageInspectWithRaw(context.Background(), ref)
	if err == nil || !client.IsErrContainerNotFound(err) {
		return true
	} else {
		return false
	}
}

func Create(image string) string {
	ref, tag, err := reference.Parse(image)
	fmt.Println("Create", ref, tag, err)
	// AuroRemove doesn't seem to work
	res, err := cli.ContainerCreate(context.Background(), &container.Config{Image: ref}, &container.HostConfig{AutoRemove: true}, nil, "")
	fmt.Println(err)
	fmt.Printf("Created a container %s\n", res.ID)
	return res.ID
}

func Start(cid string) {
	err := cli.ContainerStart(context.Background(), cid, types.ContainerStartOptions{})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Started container %s\n", cid)
}

func Logs(cid string) string {
	body, err := cli.ContainerLogs(context.Background(), cid, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Timestamps: false,
		Details:    false,
		Follow:     false,
	})
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer body.Close()
	content, err := ioutil.ReadAll(body)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", content)
	return string(content)
}

func Copy(cid string, filename string, path string) {
	file, err := os.Open(filename)
	fmt.Println(file, err)
	err = cli.CopyToContainer(context.Background(), cid, path, bufio.NewReader(file), types.CopyToContainerOptions{
		AllowOverwriteDirWithFile: true,
	})
	fmt.Println(err)
	if err != nil {
		panic(err)
	}
}
