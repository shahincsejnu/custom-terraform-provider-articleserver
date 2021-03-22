package articleserver

// The helper/schema library is part of Terraform Core. It abstracts many of the complexities and ensures consistency between providers.
import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider that will be returned by the ProviderFunc in main.go
//  The *schema.Provider type can accept:
// the resources it supports (ResourcesMap and DataSourcesMap)
// configuration keys (properties in *schema.Schema{})
// any callbacks to configure (ConfigureContextFunc)
func Provider() *schema.Provider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"articleserver_resourceCRUD": resourceCRUD(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"articleserver_articles": dataSourceArticles(),
			"articleserver_article":  dataSourceArticle(),
		},
	}
}
