//  Copyright (c) 2015 Couchbase, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 		http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package json

import (
	"fmt"

	"github.com/blevesearch/bleve/v2/registry"
	"github.com/blevesearch/bleve/v2/search/highlight"
	simpleFragmenter "github.com/blevesearch/bleve/v2/search/highlight/fragmenter/simple"
	simpleHighlighter "github.com/blevesearch/bleve/v2/search/highlight/highlighter/simple"

	jsonFormatter "github.com/anyproto/anytype-heart/pkg/lib/localstore/ftsearch/jsonhighlighter/json"
)

const Name = "json"

func Constructor(config map[string]interface{}, cache *registry.Cache) (highlight.Highlighter, error) {

	fragmenter, err := cache.FragmenterNamed(simpleFragmenter.Name)
	if err != nil {
		return nil, fmt.Errorf("error building fragmenter: %w", err)
	}

	formatter, err := cache.FragmentFormatterNamed(jsonFormatter.Name)
	if err != nil {
		return nil, fmt.Errorf("error building fragment formatter: %w", err)
	}

	return simpleHighlighter.NewHighlighter(
			fragmenter,
			formatter,
			""),
		nil
}

func init() {
	registry.RegisterHighlighter(Name, Constructor)
}
