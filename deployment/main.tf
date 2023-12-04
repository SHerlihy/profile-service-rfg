terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.16"
    }
  }

  required_version = ">= 1.2.0"
}

provider "aws" {
  region = "eu-west-2"
}

resource "aws_default_vpc" "default" {
  tags = {
    Name = "Default VPC"
  }
}

resource "aws_default_subnet" "default_az1" {
  availability_zone = "eu-west-2b"

  tags = {
    Name = "Default subnet for eu-west-2b"
  }
}

resource "aws_security_group" "profile_service" {
  name   = "profile_service"
  vpc_id = aws_default_vpc.default.id
}

resource "aws_security_group_rule" "ingress_ssh" {
  type              = "ingress"
  security_group_id = aws_security_group.profile_service.id

  from_port   = "22"
  to_port     = "22"
  protocol    = "tcp"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "ingress_redis" {
  type              = "ingress"
  security_group_id = aws_security_group.profile_service.id

  from_port   = 6379
  to_port     = 6379
  protocol    = "TCP"
  cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group_rule" "egress_all_ports" {
  type              = "egress"
  security_group_id = aws_security_group.profile_service.id

  from_port   = 0
  to_port     = 0
  protocol    = "-1"
  cidr_blocks = ["0.0.0.0/0"]
}

data "aws_ami" "amazon-linux-2" {
  most_recent = true

  filter {
    name   = "owner-alias"
    values = ["amazon"]
  }

  filter {
    name   = "name"
    values = ["amzn2-ami-hvm*"]
  }
}

resource "aws_key_pair" "profile" {
  key_name   = "profile"
  public_key = file("./.ssh/id_rsa.pub")
}

resource "aws_instance" "profile" {
  ami           = data.aws_ami.amazon-linux-2.id
  instance_type = "t2.micro"

  associate_public_ip_address = true

  subnet_id = aws_default_subnet.default_az1.id

  vpc_security_group_ids = [aws_security_group.profile_service.id]

  key_name = aws_key_pair.profile.key_name
}

resource "terraform_data" "provision_server" {
  connection {
    type = "ssh"
    port = "22"

    host = aws_instance.profile.public_ip
    user = "ec2-user"

    private_key = file("./.ssh/id_rsa")

    timeout = "2m"
  }

  provisioner "file" {
    source      = var.redis_proxy_key_dest
    destination = "/home/ec2-user/.ssh/id_rsa"
  }

  provisioner "file" {
    source      = "./provision_scripts/redis-cli.sh"
    destination = "/tmp/redis-cli.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod 400 /home/ec2-user/.ssh/id_rsa",
      "chmod +x /tmp/redis-cli.sh",
      "/tmp/redis-cli.sh",
      "sudo touch /var/log/ssh_fwd_proxy.txt",
      "sudo chmod 666 /var/log/ssh_fwd_proxy.txt",
            "nohup ssh -i '/home/ec2-user/.ssh/id_rsa' -f -N -T -L 6380:${var.redis_endpoint}:6379 ${var.proxy_access} -o StrictHostKeyChecking=no ServerAliveInterval=10 ClientAliveInterval=10 ServerAliveCountMax=3 ExitOnForwardFailure=yes &>> /var/log/ssh_fwd_proxy.txt"
    ]
  }


    provisioner "file" {
        source = "./main"
        destination = "/home/ec2-user/main"
    }

    provisioner "remote-exec" {
        inline =[
      "sudo touch /var/log/server_out.txt",
      "sudo chmod 666 /var/log/server_out.txt",
            "chmod +x /home/ec2-user/main",
            "nohup /home/ec2-user/main &>> /var/log/server_out.txt &"
        ]
    }

}
