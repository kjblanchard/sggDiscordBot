resource "kubernetes_service" "kubernetes_service_module" {
  metadata {
    name = var.service_name
  }
  spec {
    selector = {
      app = "${var.service_selector}"
    }
    session_affinity = "ClientIP"

    dynamic "port" {
      for_each = var.ports
      content {
        port = port.value.port
        name = port.value.name
        dynamic "target_port" {
          for_each = port.value.target_port
          content {
            target_port = port.value.target_port
          }
        }
      }
    }
    dynamic "port" {
      for_each = var.node_ports
      content {
        port      = port.value.port
        name      = port.value.name
        node_port = port.value.node_port
      }
    }

    type = var.service_type
  }
}