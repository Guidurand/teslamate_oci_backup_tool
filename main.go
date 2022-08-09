package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/akamensky/argparse"
	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	"guidurand.go/teslamatebackup/cleanup"
	"guidurand.go/teslamatebackup/upload"
)

func main() {
	config := common.DefaultConfigProvider()
	sclient, err := objectstorage.NewObjectStorageClientWithConfigurationProvider(config)
	if err != nil {
		log.Fatalln(err)
	}

	parser := argparse.NewParser("teslamatebackup", "Upload and manage backup from teslamate")

	bucketOCI := parser.String("b", "bucket", &argparse.Options{Required: true, Help: "Backup OCI bucket"})

	uploadCommand := parser.NewCommand("upload", "Upload file to a OCI bucket")
	backupFilepath := uploadCommand.String("f", "file", &argparse.Options{Required: true, Help: "Backup file path to upload"})

	cleanupCommand := parser.NewCommand("cleanup", "Remove files depending retention duration")
	cleanupRetention := cleanupCommand.Int("r", "retention", &argparse.Options{Required: true, Help: "Retention duration in days"})

	err = parser.Parse(os.Args)
	if err != nil {
		log.Fatalln(parser.Usage(err))
	}

	if uploadCommand.Happened() {
		upload.UploadFile(*backupFilepath, *bucketOCI, sclient)
	}

	if cleanupCommand.Happened() {
		files := cleanup.ListFiles(*bucketOCI, sclient)
		filesToDel := cleanup.CheckDate(files, *cleanupRetention)
		cleanup.DeleteObject(filesToDel, *bucketOCI, sclient)
	}
}
