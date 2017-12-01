# Terraform BelugaCDN provider

## How to install
- `go get -u github.com/danielstutzman/terraform-provider-belugacdn`
- Edit `~/.terraformrc` to add the following (replacing $GOPATH with your actual GOPATH):
  ```
  providers {
    belugacdn = "$GOPATH/bin/terraform-provider-belugacdn"
  }
  ```
- `terraform init`

## Example usage

The following configuration creates a site at `www.danstutzman.com`
that serves from the AWS S3 bucket named `danstutzman-www-danstutzman-com`.

```
variable "belugacdn_username" {}
variable "belugacdn_password" {}

provider "belugacdn" {
  username = "${var.belugacdn_username}"
  password = "${var.belugacdn_password}"
}

resource "belugacdn_site" "www-danstutzman-com" {
  name    = "www.danstutzman.com"
  origin  = "danstutzman-www-danstutzman-com.s3-website-us-east-1.amazonaws.com"
  hostnames = ["www.danstutzman.com"]
  redirect_http_to_https = true
}

resource "belugacdn_ssl_certificate" "www-danstutzman-com" {
  certificate = "${acme_certificate.www-danstutzman-com.certificate_pem}"
  chain       = "${acme_certificate.www-danstutzman-com.issuer_pem}"
  key         = "${acme_certificate.www-danstutzman-com.private_key_pem}"
  site        = "${belugacdn_site.www-danstutzman-com.name}"
}

resource "aws_route53_record" "www-danstutzman-com" {
  zone_id = "${aws_route53_zone.danstutzman-com.zone_id}"
  name    = "www.danstutzman.com"
  type    = "CNAME"
  ttl     = "3600"
  records = ["${belugacdn_site.www-danstutzman-com.cname}"]
}
```
