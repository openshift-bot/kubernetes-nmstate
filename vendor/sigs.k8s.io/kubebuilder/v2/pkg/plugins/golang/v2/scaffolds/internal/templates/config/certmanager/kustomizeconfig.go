/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package certmanager

import (
	"path/filepath"

	"sigs.k8s.io/kubebuilder/v2/pkg/model/file"
)

var _ file.Template = &KustomizeConfig{}

// KustomizeConfig scaffolds a file that configures the kustomization for the certmanager folder
type KustomizeConfig struct {
	file.TemplateMixin
}

// SetTemplateDefaults implements file.Template
func (f *KustomizeConfig) SetTemplateDefaults() error {
	if f.Path == "" {
		f.Path = filepath.Join("config", "certmanager", "kustomizeconfig.yaml")
	}

	f.TemplateBody = kustomizeConfigTemplate

	return nil
}

//nolint:lll
const kustomizeConfigTemplate = `# This configuration is for teaching kustomize how to update name ref and var substitution 
nameReference:
- kind: Issuer
  group: cert-manager.io
  fieldSpecs:
  - kind: Certificate
    group: cert-manager.io
    path: spec/issuerRef/name

varReference:
- kind: Certificate
  group: cert-manager.io
  path: spec/commonName
- kind: Certificate
  group: cert-manager.io
  path: spec/dnsNames
`
