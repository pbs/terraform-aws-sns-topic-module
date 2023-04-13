variable "name" {
  description = "Name of the SNS Topic"
  default     = null
  type        = string
}

variable "fifo_topic" {
  description = "Boolean indicating whether or not to create a FIFO (first-in-first-out) topic."
  default     = false
  type        = bool
}
