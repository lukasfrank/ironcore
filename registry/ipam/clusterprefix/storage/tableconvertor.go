// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package storage

import (
	"context"

	"github.com/onmetal/controller-utils/conditionutils"
	"github.com/onmetal/onmetal-api/apis/ipam"
	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/meta/table"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

type convertor struct{}

var (
	objectMetaSwaggerDoc = metav1.ObjectMeta{}.SwaggerDoc()

	headers = []metav1.TableColumnDefinition{
		{Name: "Name", Type: "string", Format: "name", Description: objectMetaSwaggerDoc["name"]},
		{Name: "Prefix", Type: "string", Description: "The managed prefix, if any"},
		{Name: "Parent", Type: "string", Description: "The parent, if any"},
		{Name: "State", Type: "string", Description: "The allocation of the prefix"},
		{Name: "Age", Type: "string", Format: "date", Description: objectMetaSwaggerDoc["creationTimestamp"]},
	}
)

func newTableConvertor() *convertor {
	return &convertor{}
}

func clusterPrefixReadyState(clusterPrefix *ipam.ClusterPrefix) string {
	readyCond := ipam.ClusterPrefixCondition{}
	conditionutils.MustFindSlice(clusterPrefix.Status.Conditions, string(ipam.ClusterPrefixReady), &readyCond)
	return readyCond.Reason
}

func (c *convertor) ConvertToTable(ctx context.Context, obj runtime.Object, tableOptions runtime.Object) (*metav1.Table, error) {
	tab := &metav1.Table{
		ColumnDefinitions: headers,
	}

	if m, err := meta.ListAccessor(obj); err == nil {
		tab.ResourceVersion = m.GetResourceVersion()
		tab.SelfLink = m.GetSelfLink()
		tab.Continue = m.GetContinue()
	} else {
		if m, err := meta.CommonAccessor(obj); err == nil {
			tab.ResourceVersion = m.GetResourceVersion()
			tab.SelfLink = m.GetSelfLink()
		}
	}

	var err error
	tab.Rows, err = table.MetaToTableRow(obj, func(obj runtime.Object, m metav1.Object, name, age string) (cells []interface{}, err error) {
		clusterPrefix := obj.(*ipam.ClusterPrefix)

		cells = append(cells, name)
		if prefix := clusterPrefix.Spec.Prefix; prefix.IsValid() {
			cells = append(cells, prefix.String())
		} else {
			cells = append(cells, "<none>")
		}
		if parentRef := clusterPrefix.Spec.ParentRef; parentRef != nil {
			cells = append(cells, parentRef.Name)
		} else {
			cells = append(cells, "<none>")
		}
		if readyState := clusterPrefixReadyState(clusterPrefix); readyState != "" {
			cells = append(cells, readyState)
		} else {
			cells = append(cells, "<unknown>")
		}
		cells = append(cells, age)

		return cells, nil
	})
	return tab, err
}