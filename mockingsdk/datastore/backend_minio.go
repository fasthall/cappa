package datastore

import (
	"io"
	"time"

	"github.com/fasthall/cappa/mq"

	"github.com/minio/minio-go"
)

type MinioAdapter struct {
	client *minio.Client
	mq     *mq.MQ
}

func NewMinio() (MinioAdapter, error) {
	endpoint := "127.0.0.1:9000"
	accessKeyID := "AVVRYS8CO23QMVH07O55"
	secretAccessKey := "2r5/xWcQIyopEX6BsM9Cwx7h4iHpvK7fOX7qgbpa"
	useSSL := false
	cli, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		return MinioAdapter{}, err
	}
	mqc, err := mq.NewMQ()
	if err != nil {
		return MinioAdapter{}, err
	}
	return MinioAdapter{client: cli, mq: mqc}, nil
}

func (cli MinioAdapter) Put(bucket string, object string, content io.Reader, objType string) error {
	found, err := cli.client.BucketExists(bucket)
	if !found {
		err = cli.client.MakeBucket(bucket, "us-east-1")
		if err != nil {
			return err
		}
	}
	_, err = cli.client.PutObject(bucket, object, content, objType)
	event := mq.Event{
		Time:   time.Now(),
		Type:   "datastore",
		Action: "put",
		Bucket: bucket,
		Object: object,
	}
	cli.mq.Send(event)
	return err
}

func (cli MinioAdapter) Get(bucket string, object string) (io.Reader, error) {
	obj, err := cli.client.GetObject(bucket, object)
	cli.client.RemoveBucket("testbucket")
	return obj, err
}
