package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi-azure-native-sdk/storage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"io/fs"
	"mime"
	"path"
	"path/filepath"
	"reflect"
)

type BlobFolder struct {
	pulumi.ResourceState

	containerName pulumi.StringOutput `pulumi:"containerName"`
}

func NewBlobFolder(ctx *pulumi.Context, containerName string, siteDir string, args *FolderArgs) (*BlobFolder, error) {

	var resource BlobFolder
	// Stack exports
	err := ctx.RegisterComponentResource("pulumi:example:BlobFolder", containerName, &resource)
	if err != nil {
		return nil, err
	}
	//creating a resource group
	resourceGroup, err := resources.NewResourceGroup(ctx, "new-resource-group", nil)
	if err != nil {
		return nil, err
	}
	//creating an azure storage account
	account, err := storage.NewStorageAccount(ctx, "newstorage", &storage.StorageAccountArgs{
		ResourceGroupName: resourceGroup.Name,
		Sku: &storage.SkuArgs{
			Name: pulumi.String("Standard_LRS"),
		},
		Kind: pulumi.String("StorageV2"),
	})
	if err != nil {
		return nil, err
	}

	staticWebsite, err := storage.NewStorageAccountStaticWebsite(ctx, "staticWebsite", &storage.StorageAccountStaticWebsiteArgs{
		AccountName:       account.Name,
		ResourceGroupName: resourceGroup.Name,
		IndexDocument:     pulumi.String("index.html"),
	})
	if err != nil {
		return nil, err
	}

	// For each file in the directory, create a blob object
	err = filepath.Walk(siteDir, func(name string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			rel, err := filepath.Rel(siteDir, name)
			if err != nil {
				return err
			}

			_, err = storage.NewBlob(ctx, rel, &storage.BlobArgs{
				ResourceGroupName: resourceGroup.Name,
				AccountName:       account.Name,
				ContainerName:     staticWebsite.ContainerName,
				Source:            pulumi.NewFileAsset(name),
				ContentType:       pulumi.String(mime.TypeByExtension(path.Ext(name))),
			})
			if err != nil {
				return err
			}
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	ctx.Export("staticWebsite", account.PrimaryEndpoints)

	return &resource, nil
}

type folderArgs struct {
}

type FolderArgs struct {
}

func (FolderArgs) ElementType() reflect.Type {

	return reflect.TypeOf((*folderArgs)(nil)).Elem()
}
