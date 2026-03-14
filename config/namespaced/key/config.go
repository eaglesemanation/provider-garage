package key

import "github.com/crossplane/upjet/v2/pkg/config"

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("garage_key", func(r *config.Resource) {
		r.ShortGroup = "key"
	})
}
