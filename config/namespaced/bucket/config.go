package bucket

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("garage_bucket", func(r *config.Resource) {
		r.ShortGroup = "bucket"
	})
	p.AddResourceConfigurator("garage_bucket_permission", func(r *config.Resource) {
		r.ShortGroup = "bucket"
		r.References["access_key_id"] = config.Reference{
			TerraformName: "garage_key",
		}
		r.References["bucket_id"] = config.Reference{
			TerraformName: "garage_bucket",
		}
	})
}
