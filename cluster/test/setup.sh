#!/usr/bin/env bash
set -aeuo pipefail

echo "Running setup.sh"

SCRIPT_DIR=$( cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )
PROVIDER_DIR=$( readlink -f "$SCRIPT_DIR/../.." )

echo "Creating garage deployment"
${KUBECTL} -n crossplane-system apply -f "$PROVIDER_DIR/examples/namespaced/providerconfig/garage.yaml"

echo "Creating default provider"
${KUBECTL} -n crossplane-system apply -f "$PROVIDER_DIR/examples/namespaced/providerconfig/providerconfig.yaml"
${KUBECTL} -n crossplane-system wait --for=condition=Available deployment --all --timeout=5m

echo "Setting up Garage layout"
NODE_ID="$(${KUBECTL} -n crossplane-system exec deployments/garage -- /garage status | tail -1 | cut -d' ' -f1)"
${KUBECTL} -n crossplane-system exec deployment/garage -- /garage layout assign -z dc1 -c 100M "$NODE_ID"
${KUBECTL} -n crossplane-system exec deployment/garage -- /garage layout apply --version 1
