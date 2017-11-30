package main

import (
	"github.com/danielstutzman/go-belugacdn"
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BELUGACDN_USERNAME", nil),
				Description: "BelugaCDN username",
			},
			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("BELUGACDN_PASSWORD", nil),
				Description: "BelugaCDN password",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"belugacdn_site": resource_belugacdn_site(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	username := d.Get("username").(string)
	password := d.Get("password").(string)

	return &belugacdn.Config{
		Username: username,
		Password: password,
	}, nil
}
