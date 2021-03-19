# custom-terraform-provider-

- Followed this [tutorial](https://learn.hashicorp.com/tutorials/terraform/provider-setup?in=terraform/providers)

## Intuitions

- Terraform searches for plugins in the format of `terraform-<TYPE>-<NAME>`  // remember this one, your binary should be in this convention always
- Here, the plugin is of type "provider" and of name "articleserver".
- To verify things are working correctly, execute the recently created binary, by:
    - `go build -o terraform-provider-articleserver`
    - `./terraform-provider-articleserver`
- All Terraform resources must have a schema. This allows the provider to map the JSON response to the schema.
- Format go code by `go fmt ./...`
- Now that you’ve implemented read and created the articles data source, verify that it works.
    - First, confirm that you are in the root directory of this project
    - `go build .`
    - Create the appropriate subdirectory within the user plugins directory for the articleserver provider if it doesn't exist already:
        - `mkdir -p ~/.terraform.d/plugins/hashicorp.com/edu/articleserver/0.2/linux_amd64`
    - Next, move the binary to the appropriate subdirectory within your user plugins directory:
        - `mv terraform-provider-articleserver ~/.terraform.d/plugins/hashicorp.com/edu/articleserver/0.2/linux_amd64`
- Navigate to the `custom-terraform-provider-articleserver/examples` directory. Here you can use your custom provider by terraform
         

# Resources

- [ ] [Source Addresses](https://www.terraform.io/docs/language/providers/requirements.html#source-addresses)
- [ ] [Provider Requirements](https://www.terraform.io/docs/language/providers/requirements.html)


     