package monitors

import "github.com/crossplane/terrajet/pkg/config"

// Configure the terraform resource
func Configure(p *config.Provider) {
	p.AddResourceConfigurator("datadog_monitor", func(r *config.Resource) {
		r.ExternalName = config.IdentifierFromProvider
		r.ShortGroup = "monitor"
	})
}
