/*
 * Copyright 2021 DADi590
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package UtilsInt

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"reflect"

	"Assist_Platforms_Unifier/Utils_APU"
)

// MOD_RET_ERR_PREFIX is the prefix to be used to define a constant at the submodule level, which shall be a return
// string with an error description on it. A unique string with the submodule name on it (can be abbreviated) must be
// appended to it, followed by " - ". For example, "CMD_DETECT - " for the Commands Detection submodule. A result
// example might be
//
//	"3234_ERR_GO_CMD_DETECT - Err 1: Some description here"
const MOD_RET_ERR_PREFIX = "3234_APU_ERR_"

///////////////////////////////////////
// Try / Catch / Finally

// Credits: https://dzone.com/articles/try-and-catch-in-golang

type Exception interface{}

func Throw(up Exception) {
	panic(up)
}

func (tcf Tcf) Do() {
	if tcf.Finally != nil {
		defer tcf.Finally()
	}
	if tcf.Catch != nil {
		defer func() {
			if r := recover(); r != nil {
				tcf.Catch(r)
			}
		}()
	}
	tcf.Try()
}

type Tcf struct {
	Try     func()
	Catch   func(Exception)
	Finally func()
}

/* Original example:
Tcf {
	Try: func() {
		fmt.Println("I tried")
		throw("Oh,...sh...")
	},
	Catch: func(e Exception) {
		fmt.Printf("Caught %v\n", e)
	},
	Finally: func() {
		fmt.Println("Finally...")
	},
}.Do()
*/
///////////////////////////////////////

/*
Tern behaves like the ternary operator (not present in Go) - though, inefficiently.

NOTICE: it does NOT offer conditional evaluation and requires a type assertion, so that would explain its inefficiency
(0.27 ns original vs 18.8 ns this implementation).

Example of use:

	func Abs(n int) int {
		return tern(n >= 0, n, -n).(int) // .(int) is a type assertion
	}

Equivalent (but faster - no type assertion necessary):

	test = func1()
	if condition {
		test = func2()
	}

Example with conditional evaluation (possibly the only way in Go):

	if condition {
		test = func1()
	} else {
		test = func2()
	}

Use type assertion to get the correct return type
*/
func Tern(statement bool, true_return interface{}, false_return interface{}) interface{} {
	if statement {
		return true_return
	}

	return false_return
}

/*
DelElemInSlice removes an element from a slice by its index.

Credits to https://stackoverflow.com/a/56591107/8228163 (optimized here).

-----------------------------------------------------------

> Params:

- slice – a pointer to the slice header

- index – the index of the element to remove

> Returns:

- nothing
*/
func DelElemInSlice(slice interface{}, index int) {
	var slice_value reflect.Value = reflect.ValueOf(slice).Elem()
	slice_value.Set(reflect.AppendSlice(slice_value.Slice(0, index), slice_value.Slice(index+1, slice_value.Len())))
}

/*
AddElemSlice adds an element to a specific index of a slice, keeping the elements' order.

-----------------------------------------------------------

> Params:

- slice – a pointer to the slice header

- element – the element to add

- index – the index to add the element on, with range [0, len(slice)]

> Returns:

- nothing
*/
func AddElemSlice(slice interface{}, element interface{}, index int) {
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
CopyOuterSlice copies all the values from an OUTER slice to a new slice, with the length and capacity of the original.

Note: the below described won't have any effect if the slice to copy has only one dimension - in that case, don't worry
at all as the function will copy all values normally. If the slice has more dimensions, read the below explanation.

I wrote “Outer“ in caps because of this example:

	var example [][]int = [][]int{{1}, {2}, {3}}

This function will copy the values of the outer slice only - which are pointers to the inner slices. If ANY value of the
inner slices gets changed, on the original slice that shall happen too, because both the original and the copy point to
the same inner slices. Only the outer slices differ - so one can add an inner slice to the copy, and it will not
show up on the original, and change values on that new inner slice - as long as the values of the original inner slices
don't change.

-----------------------------------------------------------

> Params:

- slice – the slice

> Returns:

- the new slice as an Interface (use type assertion to get the correct slice type)
*/
func CopyOuterSlice(slice interface{}) interface{} {
	var slice_value reflect.Value = reflect.ValueOf(slice)
	var new_slice reflect.Value = reflect.MakeSlice(slice_value.Type(), slice_value.Len(), slice_value.Cap())
	reflect.Copy(new_slice, slice_value)

	return new_slice.Interface()
}

/*
CopySlice copies all the values from slice/array to a new slice/array, with the length and capacity of the original,
provided both slices/arrays have the same type (that includes the length of each dimension with arrays).

NOTE: this function is slow, according to what someone told me. Don't use unless you really need to copy all values
from multidimensional slices/arrays.

-----------------------------------------------------------

> Params:

- destination – a pointer to an empty destination slice header/array

- source – the source slice/array

> Returns:

- nothing
*/
func CopySlice(destination interface{}, source interface{}) {
	var buf *bytes.Buffer = new(bytes.Buffer)
	var err error = gob.NewEncoder(buf).Encode(source)
	if err != nil {
		panic(err)
	}
	err = gob.NewDecoder(buf).Decode(destination)
	if err != nil {
		panic(err)
	}
}

/*
PanicInt calls the built-in panic() function, but with a string with the standard format to this module.

The format is:

	GlobalUtils_APU.APU_ERR_PREFIX + err_code + ": " + description

-----------------------------------------------------------

> Params:

- err_code – a unique error identifier for the submodule that calls the function

- description – a description of the error

> Returns:

- nothing
*/
func PanicInt(err_code float32, description string) {
	panic(Utils_APU.APU_ERR_PREFIX + fmt.Sprint(err_code) + ": " + description)
}
