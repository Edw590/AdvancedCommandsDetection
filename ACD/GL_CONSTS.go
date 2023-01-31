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

package ACD

// VERSION is the constant to check to know the version of the compiled module. The format is
//
//	"yyyy-MM-dd -- HH:mm:ss.SSSSSS ([timezone taken from the system])"
const VERSION string = "2023-01-31 -- 20:26:00.366568 (Hora padrão de GMT)"

const MOD_NAME string = "Advanced Commands Detection"

// APU_ERR_PREFIX is the prefix to be used after MOD_RET_ERR_PREFIX and its additions to return a custom error (that is,
// an error that is not from third-party/Go libraries). A string must be appended with the format "X: Y", in which X is
// a unique error identifier (float - 1 or 1.1) for the submodule, and Y is an error description. A result example
// might be
//
//	"Err 1: Some description here"
const APU_ERR_PREFIX = "Err "
