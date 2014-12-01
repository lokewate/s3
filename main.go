package main

import (
	"flag"
	"github.com/goamz/goamz/aws"
	"github.com/goamz/goamz/s3"
	"log"
)

func main() {
	filename := flag.String("uploadfile", "", "File to be uploaded")
	//	accessKey := flag.String("accesskey", "", "S3 access key")
	//	secretKey := flag.String("secretkey", "", "S3 secret key")
	flag.Parse()

	auth, err := aws.EnvAuth()
	s3Inst := s3.New(auth, aws.SAEast)
	result, err := s3Inst.Bucket("nayana-dev-sin").List("", "", "", 100)
	log.Println(result)
}
