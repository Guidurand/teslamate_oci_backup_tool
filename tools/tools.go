package tools

import (
	"context"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/objectstorage"
	log "github.com/sirupsen/logrus"
)

func InitStorageClient() objectstorage.ObjectStorageClient {
	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	if clerr != nil {
		log.Fatalln(clerr)
	}
	return c
}

func GetNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) string {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	if err != nil {
		log.Fatalln(err)
	}
	return *r.Value
}
