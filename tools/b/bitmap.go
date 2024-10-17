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
package pkgb

import (
	"errors"
	"math"
	"strconv"
)

type BitMap struct {
	m []uint64
	n uint64
}

func New(maxNum int) (*BitMap, error) {
	if maxNum <= 0 || maxNum > math.MaxInt {
		return nil, errors.New(strconv.Itoa(maxNum) + " cap can't supported")
	}
	bm := make([]uint64, (maxNum+64)>>6, (maxNum+64)>>6)
	return &BitMap{
		m: bm,
		n: uint64(maxNum),
	}, nil
}

func (bm *BitMap) Contain(num uint64) bool {
	return num <= bm.n && (bm.m[num>>6]&(uint64(1)<<(num&63)) != 0)
}

func (bm *BitMap) Add(num uint64) {
	if num <= bm.n {
		bm.m[num>>6] |= uint64(1) << (num & 63)
	}
}

func (bm *BitMap) Del(num uint64) {
	if num <= bm.n {
		bm.m[num>>6] &= ^(uint64(1) << (num & 63))
	}
}
