package main

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		f, err := NewBlobFolder(ctx, "resume-container", "./website", &FolderArgs{})
		if err != nil {
		}
		ctx.Export("bucketName", f.containerName)
		return nil
	})
}
