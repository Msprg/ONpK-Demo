# --- root/main.tf ---

data "http" "my_public_ip" {
  url = "http://ipv4.icanhazip.com"
}

data "cloudinit_config" "user_data" {
  gzip          = false
  base64_encode = false

  part {
    content_type = "text/x-shellscript"
    filename     = "userdata_base"
    content      = file("${path.module}/scripts/base.sh")
  }

  part {
    content_type = "text/x-shellscript"
    filename     = "userdata_docker"
    content      = file("${path.module}/scripts/docker.sh")
  }

  part {
    content_type = "text/x-shellscript"
    filename     = "userdata_minikube"
    content      = file("${path.module}/scripts/minikube.sh")
  }
}

resource "openstack_compute_keypair_v2" "keypair" {
  name = "${local.project}-${var.environment}-keypair"
}

resource "local_file" "private_key" {
  content         = openstack_compute_keypair_v2.keypair.private_key
  filename        = "${path.module}/${openstack_compute_keypair_v2.keypair.name}.pem"
  file_permission = "0400"
}

module "instance" {
  source        = "github.com/drajcan-nephthys/kis-onpk/terraform/modules/compute"
  project       = local.project
  environment   = var.environment
  my_public_ip  = data.http.my_public_ip.response_body
  key_pair_name = openstack_compute_keypair_v2.keypair.name
  flavor_name   = var.flavor_name
  user_data     = data.cloudinit_config.user_data.rendered
}
