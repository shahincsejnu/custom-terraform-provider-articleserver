terraform {
  required_providers {
    articleserver = {
      version = "0.2"
      source = "hashicorp.com/edu/articleserver"
    }
  }
}

data "articleserver_article" "single" {}

output "single_article" {
  value = data.articleserver_article.single.article
}