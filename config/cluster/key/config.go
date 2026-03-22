package key

import (
	"github.com/crossplane/upjet/v2/pkg/config"
)

func Configure(p *config.Provider) {
	p.AddResourceConfigurator("garage_key", func(r *config.Resource) {
		r.ShortGroup = "key"
		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]any) (map[string][]byte, error) {
			conn := map[string][]byte{}
			if a, ok := attr["id"].(string); ok {
				conn["access_key_id"] = []byte(a)
			}
			if a, ok := attr["secret_access_key"].(string); ok {
				conn["secret_access_key"] = []byte(a)
			}
			return conn, nil
		}
	})
}
