resource "aws_sns_topic" "topic" {
  name = local.name

  fifo_topic = var.fifo_topic

  tags = local.tags
}
