package upload

import (
	"context"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	log "github.com/sirupsen/logrus"
	"guidurand.go/teslamatebackup/tools"
)

func UploadFile(filepath string, bucket string, sc objectstorage.ObjectStorageClient) {

	ctx := context.Background()
	bname := bucket
	namespace := tools.GetNamespace(ctx, sc)

	filename := path.Base(filepath)
	filesize, e := getObjectSize(filepath)
	if e != nil {
		log.Fatalln(e)
	}
	file, e := os.Open(filepath)
	if e != nil {
		log.Fatalln(e)
	}
	defer file.Close()

	e = putObject(ctx, sc, namespace, bname, filename, filesize, file, nil)
	if e != nil {
		log.Fatalln(e)
	}

}

func putObject(ctx context.Context, c objectstorage.ObjectStorageClient, namespace, bucketname, objectname string, contentLen int64, content io.ReadCloser, metadata map[string]string) error {
	request := objectstorage.PutObjectRequest{
		NamespaceName: &namespace,
		BucketName:    &bucketname,
		ObjectName:    &objectname,
		ContentLength: &contentLen,
		PutObjectBody: content,
		OpcMeta:       metadata,
	}
	_, err := c.PutObject(ctx, request)
	fmt.Printf("Upload object %s to bucket %s\n", objectname, bucketname)
	return err
}

func getObjectSize(filepath string) (int64, error) {
	fi, err := os.Stat(filepath)
	if err != nil {
		return 0, err
	}
	return fi.Size(), nil
}
