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
package tg

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"runtime"
	"time"
)

type BasisDataType struct {
	A int
	B string
}

type ComplexDataType struct {
	C bool
	D int
	E *BasisDataType
	F []byte
	G string
}

func (d *BasisDataType) ConvertToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *BasisDataType) ConvertFromBytes(b []byte) error {
	if b == nil {
		return errors.New("no data to convert")
	}

	buf := new(bytes.Buffer)
	buf.Write(b)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(d); err != nil {
		return err
	}

	return nil
}

func (d *ComplexDataType) ConvertToBytes() ([]byte, error) {
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	if err := enc.Encode(d); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (d *ComplexDataType) ConvertFromBytes(b []byte) error {
	if b == nil {
		return errors.New("no data to convert")
	}

	buf := new(bytes.Buffer)
	buf.Write(b)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(d); err != nil {
		return err
	}

	return nil
}

func CurrentDate() string {
	return time.Now().Format("2006-01-02")
}

func CurrentDateTime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

func CostTime(start time.Time) {
	_, _, lineNo, _ := runtime.Caller(1)
	fmt.Printf("Line %d: \u001B[0;31;6m cost time(ms) = %d \n \u001B[0;30;6m", lineNo, time.Now().UnixMilli()-start.UnixMilli())
}
