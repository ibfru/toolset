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
package pkgb_test

import (
	tb "github.com/ibfru/toolset/tools/b"
	"math/rand"
	"testing"
)

func TestBitMap(t *testing.T) {
	bitMap, _ := tb.New(1024)
	num := uint64(123)
	bitMap.Add(num)
	if !bitMap.Contain(num) {
		t.Error("failed")
	}

	bitMap.Del(num)
	if bitMap.Contain(num) {
		t.Error("failed")
	}
}

func BenchmarkBitMap(b *testing.B) {
	b.ResetTimer()
	bitMap, _ := tb.New(2 << 30)
	for i := 0; i < b.N; i++ {
		bitMap.Add(rand.Uint64())
		bitMap.Del(rand.Uint64())
		bitMap.Contain(rand.Uint64())
	}
	b.StopTimer()
}
