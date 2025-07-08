variable "aws_region" {
  default = "eu-west-2"
}

variable "db_name" {
  default = "rdb"
}

variable "db_username" {
  default = "postgres"
}

variable "db_password" {
  type      = string
  sensitive = true
}
