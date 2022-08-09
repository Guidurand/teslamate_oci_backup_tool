package cleanup

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	"guidurand.go/teslamatebackup/tools"
)

func ListFiles(bucket string, sc objectstorage.ObjectStorageClient) (f []string) {
	ctx := context.Background()
	namespace := tools.GetNamespace(ctx, sc)
	prefix := "teslamate_backup"

	req := objectstorage.ListObjectsRequest{
		NamespaceName: &namespace,
		Prefix:        &prefix,
		BucketName:    &bucket}

	resp, err := sc.ListObjects(ctx, req)
	if err != nil {
		log.Fatalln(err)
	}

	files := make([]string, 0)

	for _, o := range resp.ListObjects.Objects {
		name := *o.Name
		files = append(files, name)
	}
	return files
}

func CheckDate(f []string, r int) (d []string) {
	toDel := make([]string, 0)
	for _, name := range f {
		fileTS, err := strconv.Atoi(strings.Split(strings.Split(name, "-")[1], ".")[0])
		if err != nil {
			log.Fatalln(err)
		}
		actualTS := int(time.Now().Unix())
		if actualTS-fileTS > r*24*60*60 {
			toDel = append(toDel, name)
		}
	}
	return toDel
}

func DeleteObject(f []string, b string, sc objectstorage.ObjectStorageClient) {
	ctx := context.Background()
	namespace := tools.GetNamespace(ctx, sc)

	for _, file := range f {
		fmt.Println(file)
		req := objectstorage.DeleteObjectRequest{
			NamespaceName: &namespace,
			BucketName:    &b,
			ObjectName:    &file,
		}
		_, err := sc.DeleteObject(ctx, req)
		if err != nil {
			log.Fatalln(err)
		}
	}

}
