// Now that you have created the provider, add the article data resource.
// the article data source will pull information on all articles by articleserver
// As a general convention, Terraform providers put each data source in their own file, named after the resource, prefixed with data_source_.

package articleserver

// The libraries imported here will be used in dataSourceArticlesRead
import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)


type Article struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author Author `json:"author"`
}

type Author struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Rating float64 `json:"rating"`
}


// The dataSourceArticles function returns a schema.Resource which defines the schema and CRUD operations for the resource.
// Since Terraform data resources should only read information (not create, update or delete), only read (ReadContext) is defined.
func dataSourceArticles() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceArticlesRead,
		// All Terraform resources must have a schema. This allows the provider to map the JSON response to the schema.
		// /articles endpoint returns an array of articles
		// since the response returns a list of articles, the articles schema should reflect that
		Schema: map[string]*schema.Schema{
			"articles": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"title": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"body": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"author": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"rating": &schema.Schema{
										Type:     schema.TypeFloat,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}

	// Now that you’ve defined your data source, you can add it to your provider.
	// In your articleserver/provider.go file, add the articles data source to the DataSourcesMap
}

// Now that you defined the articles schema, we can implement the dataSourceArticlesRead function.
// This function creates a new GET request to localhost:8080/api/articles
// Then, it decodes the response into a []map[string]interface{}
// The d.Set("articles", articles) function sets the response body (list of articles object)
// to Terraform articles data source, assigning each value to its respective schema position.
// Finally, it uses SetID to set the resource ID.
// Notice that this function returns a diag.Diagnostics type, which can return multiple errors and warnings to Terraform,
// giving users more robust error and warning messages.
// used the diag.FromErr() helper function to convert a Go error to a diag.Diagnostics type.
// The existence of a non-blank ID tells Terraform that a resource was created.
// This ID can be any string value, but should be a value that Terraform can use to read the resource again.
// Since this data resource doesn't have a unique ID, you set the ID to the current UNIX time, which will force this
// resource to refresh during every Terraform apply.
func dataSourceArticlesRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// warnings or errors can be collected in a slice type
	var diags diag.Diagnostics

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/articles", "http://localhost:8080/api"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTYxNzE3MDB9.lt8XHO4bnB7Y5cnQyWRm-qIqCvTzS_7EgXcEHQNmpu8")

	r, err := client.Do(req)

	if err != nil {
		fmt.Println("error occured")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// articles := make([]map[string]interface{}, 0)
	var articles []Article
	err = json.NewDecoder(r.Body).Decode(&articles)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("articles", flattenArticles(articles)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func flattenArticles(articles []Article) []map[string]interface{} {
	specs := make([]map[string]interface{}, len(articles))
	for i := range articles {
		specs[i] = map[string]interface{}{
			"id": articles[i].ID,
			"title": articles[i].Title,
			"body": articles[i].Body,
			"author": flattenAuthor(articles[i]),
		}
	}

	return specs
}

func flattenAuthor(article Article) []map[string]interface{} {
	specs := make([]map[string]interface{}, 1)

	specs[0] = map[string]interface{}{
		"id": article.Author.ID,
		"name": article.Author.Name,
		"rating": article.Author.Rating,
	}

	return specs;
}


