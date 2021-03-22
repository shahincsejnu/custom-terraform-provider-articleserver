terraform {
  required_providers {
    articleserver = {
      version = "0.2"
      source = "hashicorp.com/edu/articleserver"
    }
  }
}

resource "articleserver_resourceCRUD" "first_resource" {
  article {
    id = "16"
    title = "terraform post"
    body = "just demo changed"
    author {
      id = "20"
      name = "terraform"
      rating = 20
    }
  }
}