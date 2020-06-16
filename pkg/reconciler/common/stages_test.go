/*
Copyright 2020 The Knative Authors

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

package common

import (
	"context"
	"testing"

	mf "github.com/manifestival/manifestival"
	"github.com/manifestival/manifestival/fake"
	util "knative.dev/operator/pkg/reconciler/common/testing"
)

func TestStages(t *testing.T) {
	client := fake.New()
	parent, _ := mf.ManifestFrom(mf.Slice{}, mf.UseClient(client))
	stages := ReconcileStages{Parent(parent)}
	manifest, _ := mf.NewManifest("testdata/manifest.yaml")
	util.AssertEqual(t, manifest.Client, nil)
	stages.Execute(context.TODO(), &manifest, nil)
	util.AssertEqual(t, manifest.Client != nil, true)
}
