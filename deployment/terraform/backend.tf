terraform {
  backend "s3" {
    bucket         = "patrykbalcerowski-portfolio"
    key            = "cloudflare/terraform.tfstate"
    region         = "eu-central-1"
    use_lockfile   = true
    encrypt        = true 
  }
}
