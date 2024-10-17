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

import "strconv"

// 仅仅是个占位符, 不占内存, 地址相同
type void struct{}

// returns a copy of the argument slice with any duplicate entries removed
func Deduplication[S ~[]E, E comparable](arr S) S {
	var set = make(map[E]void)
	var v void

	for _, ele := range arr {
		set[ele] = v
	}

	size := len(set)
	var arrCopy = make([]E, size)
	i := 0
	for key := range set {
		arrCopy[i] = key
		i++
	}

	return arrCopy
}

// returns the sum (concatenation, for strings) of its arguments.
// ~int 可以识别 int intA // type intA int
func Sum[T ~int | ~float64 | ~string](x ...T) float64 {
	var tmp float64

	for _, value := range x {
		switch interface{}(value).(type) {
		case string:
			t, err := strconv.ParseFloat(any(value).(string), 64)
			if err != nil {
				panic("string can not convert to number!")
			}
			tmp += t
		case int:
			tmp += float64(any(value).(int))
		default:
			tmp += any(value).(float64)
		}
	}
	//return any(tmp).(T)
	return tmp
}
