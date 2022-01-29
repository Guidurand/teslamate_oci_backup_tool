package tools

import (
	"context"

	"github.com/oracle/oci-go-sdk/common"
	"github.com/oracle/oci-go-sdk/example/helpers"
	"github.com/oracle/oci-go-sdk/objectstorage"
)

func InitStorageClient() objectstorage.ObjectStorageClient {
	c, clerr := objectstorage.NewObjectStorageClientWithConfigurationProvider(common.DefaultConfigProvider())
	helpers.FatalIfError(clerr)
	return c
}

func GetNamespace(ctx context.Context, c objectstorage.ObjectStorageClient) string {
	request := objectstorage.GetNamespaceRequest{}
	r, err := c.GetNamespace(ctx, request)
	helpers.FatalIfError(err)
	return *r.Value
}
