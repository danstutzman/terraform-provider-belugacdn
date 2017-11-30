package main

import (
	"fmt"
	"github.com/danielstutzman/go-belugacdn"
	"github.com/hashicorp/terraform/helper/schema"
)

func resource_belugacdn_ssl_certificate() *schema.Resource {
	return &schema.Resource{
		Create: resource_belugacdn_ssl_certificate_create,
		Read:   resource_belugacdn_ssl_certificate_read,
		Update: resource_belugacdn_ssl_certificate_update,
		Delete: resource_belugacdn_ssl_certificate_delete,

		Schema: map[string]*schema.Schema{
			"certificate": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"chain": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"key": &schema.Schema{
				Type:      schema.TypeString,
				Required:  true,
				Sensitive: true,
			},
			"site": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func convertDataToCreateSslCertificateInput(
	d *schema.ResourceData) belugacdn.CreateSslCertificateInput {

	input := belugacdn.CreateSslCertificateInput{
		Certificate: d.Get("certificate").(string),
		Chain:       d.Get("chain").(string),
		Key:         d.Get("key").(string),
		Site:        d.Get("site").(string),
	}
	return input
}

func resource_belugacdn_ssl_certificate_create(d *schema.ResourceData, m interface{}) error {
	config := m.(*belugacdn.Config)
	input := convertDataToCreateSslCertificateInput(d)

	cert, err := config.CreateSslCertificate(input)
	if err != nil {
		return fmt.Errorf("Error from CreateSslCertificate: %s", err)
	}

	d.SetId(cert.Id)

	return err
}

func resource_belugacdn_ssl_certificate_read(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_belugacdn_ssl_certificate_update(d *schema.ResourceData, m interface{}) error {
	return nil
}

func resource_belugacdn_ssl_certificate_delete(d *schema.ResourceData, m interface{}) error {
	config := m.(*belugacdn.Config)
	commonName := d.Get("site").(string)

	err := config.DeleteSslCertificate(commonName)
	if err != nil {
		return fmt.Errorf("Error from DeleteSslCertificate: %s", err)
	}

	return err
}
