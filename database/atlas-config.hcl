variable "url" {
  type = string
}

variable "dev_url" {
  type = string
}

env "default" {
  // Declare where the schema definition resides.
  // Also supported: ["multi.hcl", "file.hcl"].
  src = "./database/atlas-schema.hcl"

  // Define the URL of the database which is managed
  // in this environment.
  url = var.url
  
  // Define the URL of the Dev Database for this environment
  // See: https://atlasgo.io/concepts/dev-database
  dev = var.dev_url

  migration {
    dir = "file://database/migrations?format=golang-migrate"
    format = atlas
  }
}