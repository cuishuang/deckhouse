/*
Copyright The Kubernetes Authors.

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/deckhouse/deckhouse/deckhouse-controller/pkg/client/informers/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// Modules returns a ModuleInformer.
	Modules() ModuleInformer
	// ModuleConfigs returns a ModuleConfigInformer.
	ModuleConfigs() ModuleConfigInformer
	// ModulePullOverrides returns a ModulePullOverrideInformer.
	ModulePullOverrides() ModulePullOverrideInformer
	// ModuleReleases returns a ModuleReleaseInformer.
	ModuleReleases() ModuleReleaseInformer
	// ModuleSources returns a ModuleSourceInformer.
	ModuleSources() ModuleSourceInformer
	// ModuleUpdatePolicies returns a ModuleUpdatePolicyInformer.
	ModuleUpdatePolicies() ModuleUpdatePolicyInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// Modules returns a ModuleInformer.
func (v *version) Modules() ModuleInformer {
	return &moduleInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ModuleConfigs returns a ModuleConfigInformer.
func (v *version) ModuleConfigs() ModuleConfigInformer {
	return &moduleConfigInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ModulePullOverrides returns a ModulePullOverrideInformer.
func (v *version) ModulePullOverrides() ModulePullOverrideInformer {
	return &modulePullOverrideInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ModuleReleases returns a ModuleReleaseInformer.
func (v *version) ModuleReleases() ModuleReleaseInformer {
	return &moduleReleaseInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ModuleSources returns a ModuleSourceInformer.
func (v *version) ModuleSources() ModuleSourceInformer {
	return &moduleSourceInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// ModuleUpdatePolicies returns a ModuleUpdatePolicyInformer.
func (v *version) ModuleUpdatePolicies() ModuleUpdatePolicyInformer {
	return &moduleUpdatePolicyInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}
