package oss

import (
	"context"
	"github.com/minio/minio-go/v7"
	"log"
	"os"
	"testing"

	"github.com/minio/minio-go/v7/pkg/credentials"
)

func TestMinIO(t *testing.T) {
	endpoint := "82.157.160.117:10005"
	accessKeyID := "root"
	secretAccessKey := "openIM123"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	// Make a new bucket called testbucket.
	bucketName := "im-mediums"
	location := "us-east-1"

	found, err := minioClient.BucketExists(context.Background(), bucketName)
	if err != nil {
		log.Fatalln(err)
	}
	// 没有找到这个 bucket 就创建
	if !found {
		if err = minioClient.MakeBucket(context.Background(), bucketName, minio.MakeBucketOptions{Region: location}); err != nil {
			log.Fatalln(err)
			return
		}
	}

	objectName := "111/testdata.png"
	contentType := "application/octet-stream"

	openFile, err := os.Open("golang.png")
	if err != nil {
		log.Fatalln(err)
	}

	objectStat, err := openFile.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	// Upload the test file with FPutObject
	info, err := minioClient.PutObject(context.Background(), bucketName, objectName, openFile, objectStat.Size(), minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Uploaded", "my-objectname", " of size: ", info.Size, "Successfully.")
	log.Println(info)
}

type User struct {
	Name string `json:"name"`
}

func TestAnge(t *testing.T) {
	arr := []User{
		{Name: "1"},
		{Name: "2"},
		{Name: "3"},
		{Name: "4"},
		{Name: "5"},
		{Name: "6"},
	}
	temp := []*User{}
	for _, user := range arr {
		temp = append(temp, &user)
	}
	for _, user := range temp {
		t.Log(user.Name)
	}
}
