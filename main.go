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
		//// Create an Azure Resource Group
		//resourceGroup, err := resources.NewResourceGroup(ctx, "new-resource-group", nil)
		//if err != nil {
		//	return err
		//}
		//
		//// Create an Azure resource (Storage Account)
		//account, err := storage.NewStorageAccount(ctx, "newstorage", &storage.StorageAccountArgs{
		//	ResourceGroupName: resourceGroup.Name,
		//	Sku: &storage.SkuArgs{
		//		Name: pulumi.String("Standard_LRS"),
		//	},
		//	Kind: pulumi.String("StorageV2"),
		//})
		//if err != nil {
		//	return err
		//}
		//
		//staticWebsite, err := storage.NewStorageAccountStaticWebsite(ctx, "staticWebsite", &storage.StorageAccountStaticWebsiteArgs{
		//	AccountName:       account.Name,
		//	ResourceGroupName: resourceGroup.Name,
		//	IndexDocument:     pulumi.String("index.html"),
		//})
		//if err != nil {
		//	return err
		//}
		//
		//_, err = storage.NewBlob(ctx, "index.html", &storage.BlobArgs{
		//	ResourceGroupName: resourceGroup.Name,
		//	AccountName:       account.Name,
		//	ContainerName:     staticWebsite.ContainerName,
		//	Source:            pulumi.NewFileAsset("./website/index.html"),
		//	ContentType:       pulumi.String("text/html"),
		//})
		//if err != nil {
		//	return err
		//}
		//
		//// Export the primary key of the Storage Account
		//ctx.Export("primaryStorageKey", pulumi.All(resourceGroup.Name, account.Name).ApplyT(
		//	func(args []interface{}) (string, error) {
		//		resourceGroupName := args[0].(string)
		//		accountName := args[1].(string)
		//		accountKeys, err := storage.ListStorageAccountKeys(ctx, &storage.ListStorageAccountKeysArgs{
		//			ResourceGroupName: resourceGroupName,
		//			AccountName:       accountName,
		//		})
		//		if err != nil {
		//			return "", err
		//		}
		//
		//		return accountKeys.Keys[0].Value, nil
		//	},
		//))
		//
		//ctx.Export("staticEndpoint", account.PrimaryEndpoints.Web())

		return nil
	})
}
