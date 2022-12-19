# --- root/outputs.tf

output "instance" {
  value = [for x in module.instance[*] : x]
}
