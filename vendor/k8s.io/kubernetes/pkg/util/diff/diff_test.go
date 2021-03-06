/*
Copyright 2016 The Kubernetes Authors.

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

package diff

import (
	"testing"
)

func TestObjectReflectDiff(t *testing.T) {
	expect := `
object[other]:
  a: 2
  b: <nil>
object[test]:
  a: 1
  b: 2
object[third]:
  a: <nil>
  b: 3`
	a := map[string]int{"test": 1, "other": 2}
	b := map[string]int{"test": 2, "third": 3}
	if actual := ObjectReflectDiff(a, b); actual != expect {
		t.Errorf("unexpected output: %s", actual)
	}
}
