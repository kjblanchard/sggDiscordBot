terraform {
  backend "s3" {
    bucket         = "supergoon-terraform-plans"
    key            = "sggDiscordBot/terraform.tfstate"
    region         = "us-east-2"
    dynamodb_table = "supergoon-terraform-plans-lock"
  }
}