output "name" {
  description = "Name of the SNS Topic"
  value       = module.topic.name
}

output "arn" {
  description = "ARN of the SNS Topic"
  value       = module.topic.arn
}
