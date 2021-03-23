package articleserver

func flattenArticles(articles []Article) []map[string]interface{} {
	specs := make([]map[string]interface{}, len(articles))
	for i := range articles {
		specs[i] = map[string]interface{}{
			"id":     articles[i].ID,
			"title":  articles[i].Title,
			"body":   articles[i].Body,
			"author": flattenAuthor(articles[i]),
		}
	}

	return specs
}

func flattenAuthor(article Article) []map[string]interface{} {
	specs := make([]map[string]interface{}, 1)

	specs[0] = map[string]interface{}{
		"id":     article.Author.ID,
		"name":   article.Author.Name,
		"rating": article.Author.Rating,
	}

	return specs
}

var tkn = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTY0ODczNzd9.HRshxnME992RbuMeB97FpBjckkmJFmjNhNnenLpK3zw"
