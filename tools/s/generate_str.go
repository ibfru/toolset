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
package ts

import (
	"crypto/rand"
	"math"
	"math/big"
)

const (
	baseChar = "0123456789AQZVFRTWESDXCBHYUGNPLOIKMJpqzmlkhjgfasdxcnbvwoeiuyrt"
)

func RandomString(length int) string {
	if length <= 0 || length > len(baseChar) {
		return ""
	}
	result := make([]byte, length)
	arr := []byte(baseChar)
	for i, region, j := 0, len(arr), 0; i < length; i++ {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(region)))
		j = int(r.Int64())
		result[i] = arr[j]
		region -= 1
		arr[j] = arr[region]
	}
	return string(result)
}

func KnuthShuffle(pokerNum int) []int {
	if pokerNum <= 0 || pokerNum >= math.MaxInt32 {
		return nil
	}

	result := make([]int, pokerNum)
	i := pokerNum
	for i != 0 {
		i -= 1
		result[i] = i + 1
	}

	i = pokerNum
	randomIdx, tmp := 0, 0
	for i != 2 {
		r, _ := rand.Int(rand.Reader, big.NewInt(int64(i)))
		randomIdx = int(r.Int64())
		i -= 1
		tmp = result[i]
		result[i] = result[randomIdx]
		result[randomIdx] = tmp
	}

	return result
}
