resource "aws_alb_target_group" "app" {
  port     = var.port
  protocol = "HTTP"
  vpc_id   = var.vpc_id

  lifecycle {
    create_before_destroy = true
  }

  health_check {
    interval            = 45
    path                = var.health_check_path
    unhealthy_threshold = 3
    healthy_threshold   = 3
    timeout             = 30
  }

  tags = {
    Name        = "${var.vpc_name}-${var.app_name}"
    terraform   = "true"
    environment = var.environment
  }
}

resource "aws_alb_listener_rule" "app" {
  listener_arn = var.alb_listener_arn

  action {
    type             = "forward"
    target_group_arn = aws_alb_target_group.app.arn
  }

  condition {
    field  = "host-header"
    values = [var.dns_name]
  }

}

