/*
Copyright 2018 The Skaffold Authors

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

package labels

import (
	"fmt"

	"github.com/GoogleContainerTools/skaffold/pkg/webhook/constants"
	"github.com/google/go-github/github"
)

// GenerateLabelsFromPR returns labels that should be applied to all deployment for this PR
func GenerateLabelsFromPR(prNumber int) map[string]string {
	m := map[string]string{}
	m["docs-controller-deployment"] = "true"
	k, v := retrieveLabel(prNumber)
	m[k] = v
	return m
}

// Selector returns the label associated with the pull request
func Selector(prNumber int) string {
	k, v := retrieveLabel(prNumber)
	return fmt.Sprintf("%s=%s", k, v)
}

// retrieveLabel returns the key and value for a label associated with this PR number
func retrieveLabel(prNumber int) (string, string) {
	value := fmt.Sprintf("docs-controller-deployment-%d", prNumber)
	return "deployment", value
}

// DocsLabelExists returns true if the label exists
func DocsLabelExists(labels []*github.Label) bool {
	for _, l := range labels {
		if l != nil {
			if *l.Name == constants.DocsLabel {
				return true
			}
		}
	}
	return false
}
