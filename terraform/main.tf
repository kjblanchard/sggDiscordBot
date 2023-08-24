module "sg_discord_bot_deployment" {
  source          = "./modules/deployment"
  image_name      = "enf3rno/sg-discord-bot"
  image_tag       = "latest"
  deployment_name = "discord-bot"
  ports = [
    {
      container_port = 80
      name           = "http"
    }
  ]
}
module "sg_discord_bot_service" {
  source           = "./modules/service"
  service_name     = "discord-bot-service"
  service_selector = "discord-bot"
  service_type     = "ClusterIP"
  ports = [
    {
      name        = "http"
      port        = 8090
      target_port = 80
    }
  ]
}