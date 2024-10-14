/*******************************************************************************
 * Copyright 2023-2024 Edw590
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 ******************************************************************************/

package ACD

import (
	"bytes"
	"encoding/gob"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// _MOD_RET_ERR_PREFIX is the prefix to be used to define a constant at the submodule level, which shall be a return
// string with an error description on it. A unique string with the submodule name on it (can be abbreviated) must be
// appended to it, followed by " - ". For example, "CMD_DETECT - " for the Commands Detection submodule. A result
// example might be
//
//	"3234_ERR_GO_CMD_DETECT - Err 1: Some description here"
const _MOD_RET_ERR_PREFIX = "3234_ACD_ERR"

/*
GetSubCmdIndex returns the index of a returned float by Main(), in the original command information array.

For example, for 10.00023, it multiplies ".00023" by MAX_SUB_CMDS and subtracts 1, which currently (100k as max) will
return 22. Or for 10.0023, it will return 229.
*/
func GetSubCmdIndex(returned_float string) int {
	s, _ := strconv.ParseFloat("."+strings.Split(returned_float, ".")[1], 32)

	return int(math.Round(s*MAX_SUB_CMDS) - 1)
}

// Exported functions above
//////////////////////////////////////////////////////////////
// Not exported functions below

/*
isSpecialCommand check if a string is an internal command, like NONE.

-----------------------------------------------------------

– Params:
  - str – the string to check

– Returns:
  - true if it's a special command, false otherwise
*/
func isSpecialCommand(str string) bool {
	return strings.HasPrefix(str, ";") && strings.HasSuffix(str, ";")
}

/*
DelElemSLICES removes an element from a slice by its index.

If the index is out of range (index < 0 || index >= len(slice)), nothing happens.

Credits to https://stackoverflow.com/a/56591107/8228163 (optimized here).

-----------------------------------------------------------

– Params:
  - slice – a pointer to the slice header
  - index – the index of the element to remove

– Returns:
  - true if the element was removed, false otherwise
*/
func DelElemSLICES(slice any, index int) bool {
	var slice_value reflect.Value = reflect.ValueOf(slice).Elem()

	if index < 0 || index >= slice_value.Len() {
		return false
	}

	slice_value.Set(reflect.AppendSlice(slice_value.Slice(0, index), slice_value.Slice(index+1, slice_value.Len())))

	return true
}

/*
AddElemSLICES adds an element to a specific index of a slice, keeping the elements' order.

-----------------------------------------------------------

– Params:
  - slice – a pointer to the slice header
  - element – the element to add
  - index – the index to add the element on, with range [0, len(slice)]

– Returns:
  - nothing
*/
func AddElemSLICES[T any](slice *[]T, element T, index int) {
	var slice_value reflect.Value = reflect.ValueOf(slice).Elem()
	var element_value reflect.Value = reflect.ValueOf(element)
	var result reflect.Value
	if index > 0 {
		result = reflect.AppendSlice(slice_value.Slice(0, index), slice_value.Slice(index-1, slice_value.Len()))
		result.Index(index).Set(element_value)
	} else {
		var element_slice reflect.Value = reflect.MakeSlice(reflect.SliceOf(element_value.Type()), 1, slice_value.Len()+1)
		element_slice.Index(0).Set(element_value)
		result = reflect.AppendSlice(element_slice, slice_value.Slice(0, slice_value.Len()))
	}
	slice_value.Set(result)
}

/*
CopyOuterSLICES copies all the values from an OUTER slice to a new slice internally created, with the length and
capacity of the original.

Note: the below described won't have any effect if the slice to copy has only one dimension - in that case,
don't worry at all as the function will copy all values normally. If the slice has more dimensions, read the
below explanation.

I wrote “Outer“ in caps because of this example:

var example [][]int = [][]int{{1}, {2}, {3}}

This function will copy the values of the outer slice only - which are pointers to the inner slices. If ANY
value of the inner slices gets changed, on the original slice that shall happen too, because both the original
and the copy point to the same inner slices. Only the outer slices differ - so one can add an inner slice to the
copy, and it will not show up on the original, and change values on that new inner slice - as long as the values
of the original inner slices don't change.

-----------------------------------------------------------

– Params:
  - slice – the slice

– Returns:
  - the new slice
*/
func CopyOuterSLICES[T any](slice T) T {
	var slice_value reflect.Value = reflect.ValueOf(slice)
	var new_slice reflect.Value = reflect.MakeSlice(slice_value.Type(), slice_value.Len(), slice_value.Cap())
	reflect.Copy(new_slice, slice_value)

	return new_slice.Interface().(T)
}

/*
CopyFullSLICES copies all the values from slice/array to a provided slice/array with the length and capacity of the
original.

Note 1: both slices/arrays must have the same type (that includes the length of each dimension with arrays).

NOTE 2: this function is slow, according to what someone told me. Don't use unless it's really necessary to copy all
values from multidimensional slices/arrays.

-----------------------------------------------------------

– Params:
  - dest – a pointer to an empty destination slice/array header
  - src – the source slice/array

– Returns:
  - true if the slice was fully copied, false if an error occurred
*/
func CopyFullSLICES[T any](dest *T, src T) bool {
	var buf *bytes.Buffer = new(bytes.Buffer)
	var err error = gob.NewEncoder(buf).Encode(src)
	if err != nil {
		return false
	}
	err = gob.NewDecoder(buf).Decode(dest)
	if err != nil {
		return false
	}

	return true
}
