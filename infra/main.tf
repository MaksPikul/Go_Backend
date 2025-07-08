provider "aws" {
  region = var.aws_region
}

resource "aws_db_instance" "main_RDB" {
  identifier         = "auth-db-instance"
  allocated_storage  = 20
  engine             = "postgres"
  engine_version     = "15.2"
  instance_class     = "db.t3.micro"
  #name               = var.db_name
  username           = var.db_username
  password           = var.db_password
  skip_final_snapshot = true
  publicly_accessible = false
}
