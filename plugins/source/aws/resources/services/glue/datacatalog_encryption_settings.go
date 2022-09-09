package glue

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/service/glue"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func DatacatalogEncryptionSettings() *schema.Table {
	return &schema.Table{
		Name:        "aws_glue_datacatalog_encryption_settings",
		Description: "Contains configuration information for maintaining Data Catalog security",
		Resolver:    fetchGlueDatacatalogEncryptionSettings,
		Multiplex:   client.ServiceAccountRegionMultiplexer("glue"),
		Columns: []schema.Column{
			{
				Name:            "account_id",
				Description:     "The AWS Account ID of the resource.",
				Type:            schema.TypeString,
				Resolver:        client.ResolveAWSAccount,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:        "return_connection_password_encrypted",
				Description: "When the ReturnConnectionPasswordEncrypted flag is set to \"true\", passwords remain encrypted in the responses of GetConnection and GetConnections",
				Type:        schema.TypeBool,
				Resolver:    schema.PathResolver("ConnectionPasswordEncryption.ReturnConnectionPasswordEncrypted"),
			},
			{
				Name:     "connection_password_encryption",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ConnectionPasswordEncryption"),
			},
			{
				Name:        "encryption_at_rest_catalog_encryption_mode",
				Description: "The encryption-at-rest mode for encrypting Data Catalog data",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("EncryptionAtRest.CatalogEncryptionMode"),
			},
			{
				Name:     "encryption_at_rest",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("EncryptionAtRest"),
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchGlueDatacatalogEncryptionSettings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Glue
	result, err := svc.GetDataCatalogEncryptionSettings(ctx, &glue.GetDataCatalogEncryptionSettingsInput{})
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	res <- result.DataCatalogEncryptionSettings
	return nil
}
