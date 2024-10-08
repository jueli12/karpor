// Copyright The Karpor Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package scheme

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/apimachinery/pkg/runtime/serializer"

	clusterinstall "github.com/KusionStack/karpor/pkg/kubernetes/apis/cluster/install"
	clusterv1beta1 "github.com/KusionStack/karpor/pkg/kubernetes/apis/cluster/v1beta1"
	searchinstall "github.com/KusionStack/karpor/pkg/kubernetes/apis/search/install"
	searchv1beta1 "github.com/KusionStack/karpor/pkg/kubernetes/apis/search/v1beta1"
	authenticationinstall "k8s.io/kubernetes/pkg/apis/authentication/install"
	coreinstall "k8s.io/kubernetes/pkg/apis/core/install"
	corev1 "k8s.io/kubernetes/pkg/apis/core/v1"
	rbacinstall "k8s.io/kubernetes/pkg/apis/rbac/install"
	rbacv1 "k8s.io/kubernetes/pkg/apis/rbac/v1"
)

var (
	// Scheme defines methods for serializing and deserializing API objects.
	Scheme = runtime.NewScheme()
	// Codecs provides methods for retrieving codecs and serializers for specific
	// versions and content types.
	Codecs = serializer.NewCodecFactory(Scheme)
	// ParameterCodec handles versioning of objects that are converted to query parameters.
	ParameterCodec = runtime.NewParameterCodec(Scheme)

	Versions = []schema.GroupVersion{
		clusterv1beta1.SchemeGroupVersion,
		searchv1beta1.SchemeGroupVersion,
		corev1.SchemeGroupVersion,
		rbacv1.SchemeGroupVersion,
	}
)

func init() {
	clusterinstall.Install(Scheme)
	searchinstall.Install(Scheme)
	coreinstall.Install(Scheme)
	rbacinstall.Install(Scheme)
	authenticationinstall.Install(Scheme)

	metav1.AddToGroupVersion(Scheme, schema.GroupVersion{Version: "v1"})
	unversioned := schema.GroupVersion{Group: "", Version: "v1"}
	Scheme.AddUnversionedTypes(
		unversioned,
		&metav1.Status{},
		&metav1.APIVersions{},
		&metav1.APIGroupList{},
		&metav1.APIGroup{},
		&metav1.APIResourceList{},
	)
}
