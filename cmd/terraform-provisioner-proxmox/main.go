package main

import (
	"github.com/rhysm/terraform-provider-proxmox/proxmox"
	"github.com/hashicorp/terraform/plugin"
	"github.com/hashicorp/terraform/terraform"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProvisionerFunc: func() terraform.ResourceProvisioner {
			return proxmox.Provisioner()
		},
	})
}
