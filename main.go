package main

import (
	"flag"

	"github.com/oracle/oci-go-sdk/objectstorage"
	"guidurand.go/teslamatebackup/cleanup"
	"guidurand.go/teslamatebackup/tools"
	"guidurand.go/teslamatebackup/upload"
)

func main() {
	clean := flag.Bool("cleanup", false, "remove files depending retention duration")
	retention := flag.Int("retention", -1, "retention duration in days")
	up := flag.Bool("upload", false, "upload a file")
	file := flag.String("file", "", "path to the file to upload")
	bucket := flag.String("bucket", "", "storage bucket name")

	flag.Parse()

	var sclient objectstorage.ObjectStorageClient

	if *clean || *up {
		sclient = tools.InitStorageClient()
	} else {
		flag.PrintDefaults()
	}

	if *clean {
		if *retention > 0 && len(*bucket) > 0 {
			files := cleanup.ListFiles(*bucket, sclient)
			filesToDel := cleanup.CheckDate(files, *retention)
			cleanup.DeleteObject(filesToDel, *bucket, sclient)
		} else {
			flag.PrintDefaults()
		}
	}

	if *up {
		if len(*file) > 0 && len(*bucket) > 0 {
			upload.UploadFile(*file, *bucket, sclient)
		} else {
			flag.PrintDefaults()
		}
	}
}
