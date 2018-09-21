/**
 * Copyright (C) 2015 Red Hat, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *         http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
/*
Copyright 2018 The Kubernetes Authors.

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

// This file was automatically generated by lister-gen

package internalversion

import (
	servicecatalog "github.com/kubernetes-incubator/service-catalog/pkg/apis/servicecatalog"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// ClusterServiceBrokerLister helps list ClusterServiceBrokers.
type ClusterServiceBrokerLister interface {
	// List lists all ClusterServiceBrokers in the indexer.
	List(selector labels.Selector) (ret []*servicecatalog.ClusterServiceBroker, err error)
	// Get retrieves the ClusterServiceBroker from the index for a given name.
	Get(name string) (*servicecatalog.ClusterServiceBroker, error)
	ClusterServiceBrokerListerExpansion
}

// clusterServiceBrokerLister implements the ClusterServiceBrokerLister interface.
type clusterServiceBrokerLister struct {
	indexer cache.Indexer
}

// NewClusterServiceBrokerLister returns a new ClusterServiceBrokerLister.
func NewClusterServiceBrokerLister(indexer cache.Indexer) ClusterServiceBrokerLister {
	return &clusterServiceBrokerLister{indexer: indexer}
}

// List lists all ClusterServiceBrokers in the indexer.
func (s *clusterServiceBrokerLister) List(selector labels.Selector) (ret []*servicecatalog.ClusterServiceBroker, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*servicecatalog.ClusterServiceBroker))
	})
	return ret, err
}

// Get retrieves the ClusterServiceBroker from the index for a given name.
func (s *clusterServiceBrokerLister) Get(name string) (*servicecatalog.ClusterServiceBroker, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(servicecatalog.Resource("clusterservicebroker"), name)
	}
	return obj.(*servicecatalog.ClusterServiceBroker), nil
}