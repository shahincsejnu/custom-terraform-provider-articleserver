// main.go is in the root of the repository
package main

// main function consume the Plugin SDK's plugin library which facilitates the RPC communication between Terraform Core and the plugin.
import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"

	"github.com/shahincsejnu/terraform_projects/custom-terraform-provider-articleserver/articleserver"
)

func main() {
	// Notice the ProviderFunc returns a *schema.Provider from the articleserver package.
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return articleserver.Provider()
		},
	})
}