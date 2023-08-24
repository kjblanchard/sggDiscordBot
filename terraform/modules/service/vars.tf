variable "service_name" {
  type        = string
  description = "The name of the service - metadata tag"
}

variable "service_selector" {
  type        = string
  description = "The actual app service selector that should match the pods, app = "
}

variable "service_type" {
  type        = string
  description = "This should be ClusterIP or NodePort probably for minikube stuff"
}

variable "ports" {
  description = "(Optional) The ports that we should listen on, useful for clusterip ports"
  default     = []
}
variable "node_ports" {
  description = "(Optional) The nodeport to listen on."
  default     = []
}