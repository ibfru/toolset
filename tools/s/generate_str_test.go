// Copyright 2024 Chao Feng
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package ts_test

import (
	ts "github.com/ibfru/toolset/tools/s"
	"math/rand"
	"testing"
)

func BenchmarkRandomString(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts.RandomString(int(rand.Int31()) & 127)
	}
	b.StopTimer()
}

func BenchmarkKnuthShuffle(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ts.KnuthShuffle(int(rand.Int31()))
	}
	b.StopTimer()
}
