package config

import (
	// Note(turkenh): we are importing this to embed provider schema document
	_ "embed"

	garageProvider "github.com/jkossis/terraform-provider-garage/garage/provider"

	ujconfig "github.com/crossplane/upjet/v2/pkg/config"

	bucketCluster "github.com/eaglesemanation/provider-garage/config/cluster/bucket"
	keyCluster "github.com/eaglesemanation/provider-garage/config/cluster/key"
	bucketNamespaced "github.com/eaglesemanation/provider-garage/config/namespaced/bucket"
	keyNamespaced "github.com/eaglesemanation/provider-garage/config/namespaced/key"
)

const (
	resourcePrefix = "garage"
	modulePath     = "github.com/eaglesemanation/provider-garage"
)

//go:embed schema.json
var providerSchema string

//go:embed provider-metadata.yaml
var providerMetadata string

// GetProvider returns provider configuration
func GetProvider() *ujconfig.Provider {
	gp := &garageProvider.GarageProvider{}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("garage.crossplane.io"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginFrameworkProvider(gp),
		ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		bucketCluster.Configure,
		keyCluster.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}

// GetProviderNamespaced returns the namespaced provider configuration
func GetProviderNamespaced() *ujconfig.Provider {
	gp := &garageProvider.GarageProvider{}

	pc := ujconfig.NewProvider([]byte(providerSchema), resourcePrefix, modulePath, []byte(providerMetadata),
		ujconfig.WithRootGroup("garage.m.crossplane.io"),
		ujconfig.WithIncludeList([]string{}),
		ujconfig.WithTerraformPluginFrameworkProvider(gp),
		ujconfig.WithTerraformPluginFrameworkIncludeList(ExternalNameConfigured()),
		ujconfig.WithFeaturesPackage("internal/features"),
		ujconfig.WithDefaultResourceOptions(
			ExternalNameConfigurations(),
		),
		ujconfig.WithExampleManifestConfiguration(ujconfig.ExampleManifestConfiguration{
			ManagedResourceNamespace: "crossplane-system",
		}))

	for _, configure := range []func(provider *ujconfig.Provider){
		// add custom config functions
		bucketNamespaced.Configure,
		keyNamespaced.Configure,
	} {
		configure(pc)
	}

	pc.ConfigureResources()
	return pc
}
