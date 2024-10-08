/*
Copyright The Karpor Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1beta1

import (
	v1beta1 "github.com/KusionStack/karpor/pkg/kubernetes/apis/search/v1beta1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// TrimRuleLister helps list TrimRules.
// All objects returned here must be treated as read-only.
type TrimRuleLister interface {
	// List lists all TrimRules in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1beta1.TrimRule, err error)
	// Get retrieves the TrimRule from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1beta1.TrimRule, error)
	TrimRuleListerExpansion
}

// trimRuleLister implements the TrimRuleLister interface.
type trimRuleLister struct {
	indexer cache.Indexer
}

// NewTrimRuleLister returns a new TrimRuleLister.
func NewTrimRuleLister(indexer cache.Indexer) TrimRuleLister {
	return &trimRuleLister{indexer: indexer}
}

// List lists all TrimRules in the indexer.
func (s *trimRuleLister) List(selector labels.Selector) (ret []*v1beta1.TrimRule, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1beta1.TrimRule))
	})
	return ret, err
}

// Get retrieves the TrimRule from the index for a given name.
func (s *trimRuleLister) Get(name string) (*v1beta1.TrimRule, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1beta1.Resource("trimrule"), name)
	}
	return obj.(*v1beta1.TrimRule), nil
}
