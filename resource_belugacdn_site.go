package main

import (
	"fmt"
	"github.com/danielstutzman/go-belugacdn"
	"github.com/hashicorp/terraform/helper/schema"
	"strconv"
)

func resource_belugacdn_site() *schema.Resource {
	return &schema.Resource{
		Create: resource_belugacdn_site_create,
		Read:   resource_belugacdn_site_read,
		Update: resource_belugacdn_site_update,
		Delete: resource_belugacdn_site_delete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"origin": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"hostnames": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"redirect_http_to_https": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func convertDataToSiteConfiguration(d *schema.ResourceData) belugacdn.SiteConfiguration {
	input := belugacdn.SiteConfiguration{
		Origin: d.Get("origin").(string),
	}

	if d.Get("redirect_http_to_https").(bool) {
		input.Rules = []belugacdn.SiteRule{
			{
				Paths: []string{"/"},
				Actions: []belugacdn.SiteAction{
					{
						Action: "canonicalizeScheme",
						Scheme: "https",
					},
				},
			},
		}
	}

	if attr, ok := d.GetOk("hostnames"); ok {
		hostnames := []string{}
		for _, s := range attr.([]interface{}) {
			hostnames = append(hostnames, s.(string))
		}
		input.Hostnames = hostnames
	}

	return input
}

func resource_belugacdn_site_create(d *schema.ResourceData, m interface{}) error {
	config := m.(*belugacdn.Config)
	siteName := d.Get("name").(string)
	input := convertDataToSiteConfiguration(d)

	site, err := config.CreateSite(siteName, input)
	if err != nil {
		return fmt.Errorf("Error from CreateSite: %s", err)
	}

	d.SetId(strconv.Itoa(site.DomainId))

	return err
}

func resource_belugacdn_site_read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_belugacdn_site_update(d *schema.ResourceData, m interface{}) error {
	config := m.(*belugacdn.Config)
	siteName := d.Get("name").(string)
	input := convertDataToSiteConfiguration(d)

	_, err := config.UpdateSite(siteName, input)

	return err
}

func resource_belugacdn_site_delete(d *schema.ResourceData, m interface{}) error {
	config := m.(*belugacdn.Config)
	siteName := d.Get("name").(string)

	err := config.DeleteSite(siteName)

	return err
}
