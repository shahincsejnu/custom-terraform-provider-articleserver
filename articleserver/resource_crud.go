package articleserver

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func resourceCRUD() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceCreate,
		ReadContext:   resourceRead,
		UpdateContext: resourceUpdate,
		DeleteContext: resourceDelete,
		Schema: map[string]*schema.Schema{
			"article": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"title": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"body": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
						},
						"author": &schema.Schema{
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Required: true,
									},
									"rating": &schema.Schema{
										Type:     schema.TypeFloat,
										Required: true,
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

func resourceCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	var id, title, body, aid, name string
	var rating float64

	articleSchema := d.Get("article").([]interface{})
	i := articleSchema[0].(map[string]interface{})

	id = i["id"].(interface{}).(string)
	title = i["title"].(interface{}).(string)
	body = i["body"].(interface{}).(string)
	authorSchema := i["author"].([]interface{})
	author := authorSchema[0].(map[string]interface{})
	aid = author["id"].(interface{}).(string)
	name = author["name"].(interface{}).(string)
	rating = author["rating"].(interface{}).(float64)

	// articles := make([]map[string]interface{}, 0)
	article := Article{
		ID:    id,
		Title: title,
		Body:  body,
		Author: Author{
			ID:     aid,
			Name:   name,
			Rating: rating,
		},
	}

	byteArticle, err := json.Marshal(article)
	if err != nil {
		log.Println(err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/article", "http://localhost:8080/api"), bytes.NewBuffer(byteArticle))
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY0Mjc2MzJ9.iqosnVH8UIEfTosCPqjSCwLhcm88zFoD1H7g4Bg1Ecs")

	r, err := client.Do(req)
	if err != nil {
		fmt.Println("error occured")
		return diag.FromErr(err)
	}

	var demoArticle Article
	err = json.NewDecoder(r.Body).Decode(&demoArticle)

	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	//d.SetId(demoArticle.ID)
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}

func resourceRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	//ID := d.Id()

	articleSchema := d.Get("article").([]interface{})
	i := articleSchema[0].(map[string]interface{})

	resID := i["id"].(interface{}).(string)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/article/%s", "http://localhost:8080/api", resID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY0Mjc2MzJ9.iqosnVH8UIEfTosCPqjSCwLhcm88zFoD1H7g4Bg1Ecs")

	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	var article []Article
	var demoArticle Article
	body, err := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(body, &demoArticle)
	if err != nil {
		return diag.FromErr(err)
	}
	article = append(article, demoArticle)

	if err := d.Set("article", flattenArticles(article)); err != nil {
		return diag.FromErr(err)
	}

	// always run
	//d.SetId(demoArticle.ID)

	return diags
}

func resourceUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	if d.HasChange("article") {
		client := &http.Client{Timeout: 10 * time.Second}
		// Warning or errors can be collected in a slice type

		var id, title, body, aid, name string
		var rating float64

		articleSchema := d.Get("article").([]interface{})
		i := articleSchema[0].(map[string]interface{})

		id = i["id"].(interface{}).(string)
		title = i["title"].(interface{}).(string)
		body = i["body"].(interface{}).(string)
		authorSchema := i["author"].([]interface{})
		author := authorSchema[0].(map[string]interface{})
		aid = author["id"].(interface{}).(string)
		name = author["name"].(interface{}).(string)
		rating = author["rating"].(interface{}).(float64)

		// articles := make([]map[string]interface{}, 0)
		article := Article{
			ID:    id,
			Title: title,
			Body:  body,
			Author: Author{
				ID:     aid,
				Name:   name,
				Rating: rating,
			},
		}

		byteArticle, err := json.Marshal(article)
		if err != nil {
			log.Println(err)
		}

		ID := article.ID

		req, err := http.NewRequest("PUT", fmt.Sprintf("%s/article%s", "http://localhost:8080/api", ID), bytes.NewBuffer(byteArticle))
		if err != nil {
			return diag.FromErr(err)
		}

		req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY0Mjc2MzJ9.iqosnVH8UIEfTosCPqjSCwLhcm88zFoD1H7g4Bg1Ecs")

		_, err = client.Do(req)
		if err != nil {
			fmt.Println("error occured")
			return diag.FromErr(err)
		}

	}

	//return resourceRead(ctx, d, m)
	return diags
}

func resourceDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	client := &http.Client{Timeout: 10 * time.Second}
	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	ID := d.Id()

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/article/%s", "http://localhost:8080/api", ID), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	req.Header.Set("Token", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY0Mjc2MzJ9.iqosnVH8UIEfTosCPqjSCwLhcm88zFoD1H7g4Bg1Ecs")

	_, err = client.Do(req)
	if err != nil {
		fmt.Println("error occured")
		return diag.FromErr(err)
	}

	// d.SetId("") is automatically called assuming delete returns no errors, but
  	// it is added here for explicitness.
  	d.SetId("")

	return diags
}
