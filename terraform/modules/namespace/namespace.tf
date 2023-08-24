resource "kubernetes_namespace" "namespace_module" {
  metadata {
    name = var.namespace_name
  }
}