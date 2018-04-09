/*
Copyright 2017 The Kubernetes Authors.

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

package regexmatch

import (
	extensions "k8s.io/api/extensions/v1beta1"

	"k8s.io/ingress-nginx/internal/ingress/annotations/parser"
	"k8s.io/ingress-nginx/internal/ingress/resolver"
)

type regexMatch struct {
	r resolver.Resolver
}

// NewParser creates a new regexMatch annotation parser
func NewParser(r resolver.Resolver) parser.IngressAnnotation {
	return regexMatch{r}
}

// Parse parses the annotations contained in the ingress rule
// used to set location type as regex
func (r regexMatch) Parse(ing *extensions.Ingress) (interface{}, error) {
	return parser.GetBoolAnnotation("regex-match", ing)
}
