package main

import (
	"github.com/gordalina/terraform-provider-mongodbatlas/mongodbatlas"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: mongodbatlas.Provider})
}
