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

import (
	"math"
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

> Params:
  - str – the string to check

> Returns:
  - true if it's a special command, false otherwise
*/
func isSpecialCommand(str string) bool {
	return strings.HasPrefix(str, ";") && strings.HasSuffix(str, ";")
}
