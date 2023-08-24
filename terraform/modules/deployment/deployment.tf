resource "kubernetes_deployment" "k8s_deployment" {
  metadata {
    name = "${var.deployment_name}-deployment"
    labels = {
      app = "${var.deployment_name}-deployment"
    }
  }
  spec {
    replicas = 1
    selector {
      match_labels = {
        app = "${var.deployment_name}"
      }
    }
    template {
      metadata {
        labels = {
          app = "${var.deployment_name}"
        }
      }
      spec {
        container {
          image = "${var.image_name}:${var.image_tag}"
          name  = var.deployment_name
          dynamic "port" {
            for_each = var.ports
            content {
              container_port = port.value.container_port
              name           = port.value.name
            }
          }
          dynamic "volume_mount" {
            for_each = var.volume_mount
            content {
              mount_path = volume_mount.value.mount_path
              sub_path   = lookup(volume_mount.value, "sub_path", null)
              name       = volume_mount.value.volume_name
              read_only  = lookup(volume_mount.value, "read_only", false)
            }
          }
          dynamic "env_from" {
            for_each = var.env_config_maps
            content {
              config_map_ref {
                name = env_from.value
              }
            }
          }
        }
        dynamic "volume" {
          for_each = var.volume_host_path
          content {
            host_path {
              path = volume.value.path_on_node
              type = lookup(volume.value, "type", null)
            }
            name = volume.value.volume_name
          }
        }
      }
    }
  }
}
