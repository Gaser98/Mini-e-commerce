variable "aws_region" {
  description = "AWS region"
  type        = string
  default     = "us-east-1"
}

variable "azs" {
  description = "Availability zones"
  type        = list(string)
  default     = ["us-east-1a", "us-east-1b"]
}

variable "db_name" {
  description = "PostgreSQL database name"
  type        = string
  default     = "demo"
}

variable "db_username" {
  description = "PostgreSQL username"
  type        = string
  default     = "demo"
}

variable "db_password" {
  description = "PostgreSQL password"
  type        = string
  sensitive   = true
}
