locals {
  config_map_name = "discord-bot-data"
}
module "sg_discord_bot_deployment" {
  source          = "./modules/deployment"
  image_name      = "enf3rno/supergoon-discord-bot"
  image_tag       = "2"
  deployment_name = "discord-bot"
  env_config_maps = [local.config_map_name]
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
resource "kubernetes_config_map" "main" {
  metadata {
    name = local.config_map_name
  }

  data = {
    DISCORD_BOT_TOKEN          = var.DISCORD_BOT_TOKEN
    DISCORD_APP_ID             = var.DISCORD_APP_ID
    DISCORD_SUPERGOON_GUILD_ID = var.DISCORD_SUPERGOON_GUILD_ID
    DISCORD_CLIENT_SECRET      = var.DISCORD_CLIENT_SECRET
    GITHUB_ACCESS_TOKEN        = var.GITHUB_ACCESS_TOKEN
    DISCORD_CLIENT_ID          = var.DISCORD_CLIENT_ID
  }

}
