package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
)

func main() {
	var err error
	err = godotenv.Load()
	failOnErr(err)
	err = downloadS3Object(
		os.Getenv("BUCKET_NAME"),
		os.Getenv("OBJECT_KEY"),
		os.Getenv("FILENAME"),
	)
	failOnErr(err)
}

// downloadS3Object stored in a AWS S3 bucket to a local file
func downloadS3Object(bucketName string, objectKey string, filename string) (err error) {
	// Create a file to download to
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY, AWS_REGION environment variables are used by LoadDefaultConfig
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return
	}

	// Create a S3 client using our configuration
	s3Client := s3.NewFromConfig(cfg)

	// Download the S3 object using the S3 manager object downloader
	downloader := manager.NewDownloader(s3Client)
	_, err = downloader.Download(context.TODO(), file, &s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})
	return
}

func failOnErr(err error) {
	if err != nil {
		log.Fatal(err.Error())
	}
}
