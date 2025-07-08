package bucket

import (
	"go_backend/internal/config"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type AwsBucketUploader struct {
	uploader *s3manager.Uploader
	bucket   string
}

func NewAwsUploader(cfg *config.CloudConfig) (*s3manager.Uploader, error) {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(cfg.Region),
	})

	return s3manager.NewUploader(sess), err
}

/*
func UploadContent() {

	if len(os.Args) != 3 {
		exitErrorf("bucket and file name required\nUsage: %s bucket_name filename",
			os.Args[0])
	}

	bucket := os.Args[1]
	filename := os.Args[2]

	file, err := os.Open(filename)
	if err != nil {
		exitErrorf("Unable to open file %q, %v", err)
	}

	defer file.Close()

	client := s3manager.

	// This is for
	sess, err := session.NewSession(s3manager.NewUploader(sess)
	)

	uploader := s3manager.NewUploader(sess)

	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(filename),
		Body:   file,
	})

	if err != nil {
		// Print the error and exit.
		exitErrorf("Unable to upload %q to %q, %v", filename, bucket, err)
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)

}

// For Now
func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func PreSignUrl_Upload() {

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	svc := s3.New(sess)

	req, _ := svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String("amzn-s3-demo-bucket"),
		Key:    aws.String("myKey"),
		Body:   strings.NewReader("EXPECTED CONTENTS"),
	})
	str, err := req.Presign(15 * time.Minute)

	log.Println("The URL is:", str, " err:", err)

}
*/
