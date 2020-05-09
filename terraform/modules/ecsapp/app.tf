data "aws_ecr_repository" "ecr" {
  name = "electric/${var.repo_name}"
}

locals {
  env_prefix = substr(var.environment, 0, 1)
}

module "mock_iam_role_mod_exec" {
  create_role             = "example"
  role_name               = "example"
  role_path               = "exaple"
  environment             = "example"
  role_type               = "example"
  vpc_name                = "example"
  app_name                = "example"
  policy_string           = "example"
  create_task_association = "example"
}

module "mock_iam_role_mod_task" {
  create_role             = "example"
  role_name               = "example"
  role_path               = "example"
  environment             = "example"
  role_type               = "example"
  vpc_name                = "example"
  app_name                = "example"
  policy_string           = "example"
  create_task_association = "example"
}


data "template_file" "task" {
  template = file("${path.module}/templates/task.tmpl")

  vars = {
    app_name      = "example"
    env           = "example"
    env_vars      = "example"
    secrets       = "example"
    mount_points  = "example"
    port_mappings = "example"
    command       = "example"
    repo_url      = "example"
    log_config    = "example"
    tag           = "example"
  }
}

resource "aws_ecs_task_definition" "this" {
  family                = "example"
  memory                = "example"
  cpu                   = "example"
  container_definitions = "example"
  task_role_arn         = "example"
  execution_role_arn    = "example"
  dynamic "volume" {
    for_each = var.add_volume ? [var.volumes] : []

    content {
      name      = volume.name
      host_path = volume.host_path
    }
  }
}

resource "aws_ecs_service" "app" {
  name                              = "example"
  cluster                           = "example"
  task_definition                   = "example"
  desired_count                     = "example"
  health_check_grace_period_seconds = "example"

  load_balancer {
    target_group_arn = aws_alb_target_group.app.arn
    container_name   = var.app_name
    container_port   = var.port
  }

  lifecycle {
    create_before_destroy = true
    ignore_changes        = ["task_definition"]
  }
}
