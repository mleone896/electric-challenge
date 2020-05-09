variable "app_name" { type = string }
variable "environment" { type = string }
variable "repo_name" { type = string }
variable "add_volume" { type = bool }
variable "vpc_name" { type = string }
variable "vpc_id" { type = string }
variable "execution_role_policy" { type = string }
variable "command" { type = list }
variable "task_role_policy" { type = string }
variable "mount_points" { type = list }
variable "environment_vars" { type = list }
variable "application_secrets" { type = list }
variable "image_tag" { type = string }
variable "desired_count" { type = number }
variable "volumes" { type = map }
variable "ecs_cluster_id" { type = string }
variable "grace_period" { type = string }
variable "port_mappings" { type = list }
variable "port" { type = string }
variable "log_configuration" {}
variable "memory" { type = string }
variable "cpu" { type = string }

variable "alb_listener_arn" {
  type    = string
  default = ""
}

variable "dns_name" {
  description = "dns name for alb host header mapping"
  type        = string
  default     = ""
}

variable "health_check_path" {
  type    = string
  default = "/v1/status"
}

variable "app_alb_needed" {
  type    = bool
  default = false
}

