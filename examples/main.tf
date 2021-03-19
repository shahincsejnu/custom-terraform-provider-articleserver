terraform {
  required_providers {
    articleserver = {
      version = "0.2"
      source = "hashicorp.com/edu/articleserver"
    }
  }
}

variable "article_name" {
  type    = string
  default = "demo"
}

data "articleserver_articles" "all" {}

# Returns all articles
output "all_articles" {
  value = data.articleserver_articles.all
}
