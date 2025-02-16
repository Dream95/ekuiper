// Copyright 2023-2024 EMQ Technologies Co., Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package json

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/lf-edge/ekuiper/v2/pkg/ast"
	mockContext "github.com/lf-edge/ekuiper/v2/pkg/mock/context"
)

func BenchmarkSimpleTuples(b *testing.B) {
	benchmarkByFiles("./testdata/simple.json", b, nil)
}

func BenchmarkSimpleTuplesWithSchema(b *testing.B) {
	schema := map[string]*ast.JsonStreamField{
		"key": {
			Type: "string",
		},
	}
	benchmarkByFiles("./testdata/simple.json", b, schema)
}

func BenchmarkSmallJSON(b *testing.B) {
	benchmarkByFiles("./testdata/small.json", b, nil)
}

func BenchmarkSmallJSONWithSchema(b *testing.B) {
	schema := map[string]*ast.JsonStreamField{
		"sid": {
			Type: "bigint",
		},
	}
	benchmarkByFiles("./testdata/small.json", b, schema)
}

func BenchmarkMediumJSON(b *testing.B) {
	benchmarkByFiles("./testdata/medium.json", b, nil)
}

func BenchmarkMediumJSONWithSchema(b *testing.B) {
	schema := map[string]*ast.JsonStreamField{
		"person": {
			Type: "struct",
			Properties: map[string]*ast.JsonStreamField{
				"id": {
					Type: "string",
				},
			},
		},
	}
	benchmarkByFiles("./testdata/medium.json", b, schema)
}

func BenchmarkLargeJSON(b *testing.B) {
	benchmarkByFiles("./testdata/large.json", b, nil)
}

func BenchmarkLargeJSONWithSchema(b *testing.B) {
	schema := map[string]*ast.JsonStreamField{
		"users": {
			Type: "array",
			Items: &ast.JsonStreamField{
				Type: "struct",
				Properties: map[string]*ast.JsonStreamField{
					"id": {
						Type: "bigint",
					},
				},
			},
		},
	}
	benchmarkByFiles("./testdata/large.json", b, schema)
}

func BenchmarkComplexTuples(b *testing.B) {
	benchmarkByFiles("./testdata/MDFD.json", b, nil)
}

func BenchmarkComplexTuplesWithSchema(b *testing.B) {
	schema := map[string]*ast.JsonStreamField{
		"STD_AbsoluteWindDirection": {
			Type: "float",
		},
	}
	benchmarkByFiles("./testdata/MDFD.json", b, schema)
}

func benchmarkByFiles(filePath string, b *testing.B, schema map[string]*ast.JsonStreamField) {
	ctx := mockContext.NewMockContext("test", "test")
	payload, err := os.ReadFile(filePath)
	if err != nil {
		b.Fatal(err)
	}
	f := NewFastJsonConverter(schema, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Decode(ctx, payload)
	}
}

func BenchmarkNativeFloatParse(b *testing.B) {
	m := make(map[string]interface{})
	data := `{"id":1.2}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		json.Unmarshal([]byte(data), &m)
	}
}

func BenchmarkFloatParse(b *testing.B) {
	ctx := mockContext.NewMockContext("test", "test")
	f := NewFastJsonConverter(nil, map[string]any{"useInt64ForWholeNumber": true})
	data := `{"id":1.2}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Decode(ctx, []byte(data))
	}
}

func BenchmarkIntParse(b *testing.B) {
	ctx := mockContext.NewMockContext("test", "test")
	f := NewFastJsonConverter(nil, map[string]any{"useInt64ForWholeNumber": true})
	data := `{"id":1}`
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		f.Decode(ctx, []byte(data))
	}
}
