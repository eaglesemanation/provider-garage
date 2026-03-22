package config

import (
	"fmt"
	"strings"

	"github.com/crossplane/upjet/v2/pkg/config"
	"github.com/hashicorp/terraform-plugin-go/tfprotov6"
)

const (
	BucketErrorMsg = "Bad request: Either id, globalAlias or search must be provided (but not several of them)"
	KeyErrorMsg    = "Bad request: Either id or search must be provided (but not both)"
)

const (
	ErrFmtNoAttribute    = "required attribute %s not found"
	ErrFmtUnexpectedType = "unexpected type for attribute %s, expected string"
)

// ExternalNameConfigs contains all external name configurations for this
// provider.
var ExternalNameConfigs = map[string]config.ExternalName{
	"garage_bucket":            idWithIgnore(BucketErrorMsg),
	"garage_bucket_permission": permissionExternalName(),
	"garage_key":               idWithIgnore(KeyErrorMsg),
}

func idWithIgnore(substr string) config.ExternalName {
	e := config.IdentifierFromProvider
	e.IsNotFoundDiagnosticFn = func(diags []*tfprotov6.Diagnostic) bool {
		for _, diag := range diags {
			if strings.Contains(diag.Detail, substr) {
				return true
			}
		}
		return false
	}
	return e
}

func permissionExternalName() config.ExternalName {
	e := config.IdentifierFromProvider
	e.GetExternalNameFn = func(tfstate map[string]any) (string, error) {
		bucketId, ok := tfstate["bucket_id"]
		if !ok {
			return "", fmt.Errorf(ErrFmtNoAttribute, "bucket_id")
		}
		bucketIdStr, ok := bucketId.(string)
		if !ok {
			return "", fmt.Errorf(ErrFmtUnexpectedType, "bucket_id")
		}
		accessKeyId, ok := tfstate["access_key_id"]
		if !ok {
			return "", fmt.Errorf(ErrFmtNoAttribute, "access_key_id")
		}
		accessKeyIdStr, ok := accessKeyId.(string)
		if !ok {
			return "", fmt.Errorf(ErrFmtUnexpectedType, "access_key_id")
		}
		return fmt.Sprintf("%s/%s", bucketIdStr, accessKeyIdStr), nil
	}
	return e
}

// ExternalNameConfigurations applies all external name configs listed in the
// table ExternalNameConfigs and sets the version of those resources to v1beta1
// assuming they will be tested.
func ExternalNameConfigurations() config.ResourceOption {
	return func(r *config.Resource) {
		if e, ok := ExternalNameConfigs[r.Name]; ok {
			r.ExternalName = e
		}
	}
}

// ExternalNameConfigured returns the list of all resources whose external name
// is configured manually.
func ExternalNameConfigured() []string {
	l := make([]string, len(ExternalNameConfigs))
	i := 0
	for name := range ExternalNameConfigs {
		// $ is added to match the exact string since the format is regex.
		l[i] = name + "$"
		i++
	}
	return l
}
