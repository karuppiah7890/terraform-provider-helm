/*
Copyright The Helm Authors.

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

package kube // import "helm.sh/helm/v3/pkg/kube"

import (
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/cli-runtime/pkg/resource"
	"k8s.io/client-go/kubernetes/scheme"
)

// AsVersioned converts the given info into a runtime.Object with the correct
// group and version set
func AsVersioned(info *resource.Info) runtime.Object {
	return convertWithMapper(info.Object, info.Mapping)
}

// convertWithMapper converts the given object with the optional provided
// RESTMapping. If no mapping is provided, the default schema versioner is used
func convertWithMapper(obj runtime.Object, mapping *meta.RESTMapping) runtime.Object {
	var gv = runtime.GroupVersioner(schema.GroupVersions(scheme.Scheme.PrioritizedVersionsAllGroups()))
	if mapping != nil {
		gv = mapping.GroupVersionKind.GroupVersion()
	}
	if obj, err := runtime.ObjectConvertor(scheme.Scheme).ConvertToVersion(obj, gv); err == nil {
		return obj
	}
	return obj
}
