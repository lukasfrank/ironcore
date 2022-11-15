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

package table

import (
	"fmt"
	"io"
)

type Header struct {
	Name string
}

type Cell any

type Row []Cell

type Table struct {
	Headers []Header
	Rows    []Row
}

func Write(table *Table, w io.Writer) {
	for i, header := range table.Headers {
		if i != 0 {
			_, _ = fmt.Fprint(w, "\t")
		}
		_, _ = fmt.Fprint(w, header.Name)
	}
	if len(table.Headers) > 0 {
		_, _ = fmt.Fprintln(w)
	}
	for i, row := range table.Rows {
		if i != 0 {
			_, _ = fmt.Fprintln(w)
		}
		for j, cell := range row {
			if j != 0 {
				_, _ = fmt.Fprint(w, "\t")
			}
			_, _ = fmt.Fprint(w, cell)
		}
	}
	if len(table.Rows) > 0 {
		_, _ = fmt.Fprintln(w)
	}
}
