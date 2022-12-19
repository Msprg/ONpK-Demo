# --- root/locals.tf ---

locals {
  kis_os_region   = "RegionOne"
  kis_os_auth_url = "http://158.193.138.33:5000/v3/"
  kis_os_endpoint_overrides = {
    compute = "http://158.193.138.33:8774/v2.1/"
    image   = "http://158.193.138.33:9292/v2.0/"
    network = "http://158.193.138.33:9696/v2.0/"
  }
  my_public_ip = "${data.http.my_public_ip.response_body}/32"
  project      = lower("${var.tenant_name}-minikube")
}
