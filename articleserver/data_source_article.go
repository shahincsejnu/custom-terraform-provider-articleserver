package articleserver

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"net/http"
	"strconv"
	"time"
)

func dataSourceArticle() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceArticleRead,
		Schema: map[string]*schema.Schema{
			"article": &schema.Schema{
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
}

func dataSourceArticleRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}

	// warnings or errors can be collected in a slice type
	var diags diag.Diagnostics

	// articles := make([]map[string]interface{}, 0)
	//article := Article{
	//	ID:     "201",
	//	Title:  "terraform post",
	//	Body:   "just test",
	//	Author: Author{
	//		ID:     "101",
	//		Name:   "terraform oka",
	//		Rating: 10,
	//	},
	//}
	//
	//byteArticle, err := json.Marshal(article);
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//req, err := http.NewRequest("POST", fmt.Sprintf("%s/article", "http://localhost:8080/api"), bytes.NewBuffer(byteArticle))
	//if err != nil {
	//	return diag.FromErr(err)
	//}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/1", "http://localhost:8080/api/article"), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	//req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Token", tkn)

	r, err := client.Do(req)

	if err != nil {
		fmt.Println("error occured")
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// articles := make([]map[string]interface{}, 0)
	var articles []Article
	var demoArticle Article
	err = json.NewDecoder(r.Body).Decode(&demoArticle)
	articles = append(articles, demoArticle)

	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("article", flattenArticles(articles)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
