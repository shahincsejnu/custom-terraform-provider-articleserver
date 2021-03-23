package articleserver

import (
	"encoding/json"
	"fmt"

	"net/http"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	client := &http.Client{Timeout: 10 * time.Second}
	// Warning or errors can be collected in a slice type
	//var diags diag.Diagnostics

	ID := "10"

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/article/%s", "http://localhost:8080/api", ID), nil)
	oka := fmt.Sprintf("%s/article/%s", "http://localhost:8080/api", ID)
	fmt.Println(oka)
	fmt.Println("request")
	if err != nil {
		return
	}
	req.Header.Set("Token", tkn)

	r, err := client.Do(req)
	fmt.Println("response")
	if err != nil {
		return
	}
	defer r.Body.Close()

	var article []Article
	var demoArticle Article
	err = json.NewDecoder(r.Body).Decode(&demoArticle)
	fmt.Println(demoArticle)
	article = append(article, demoArticle)

	fmt.Println(article)

	//if err := d.Set("article", flattenArticles(article)); err != nil {
	//	return
	//}

	return
}
