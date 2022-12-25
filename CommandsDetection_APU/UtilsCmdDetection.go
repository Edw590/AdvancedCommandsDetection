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

package CommandsDetection_APU

import (
	"strconv"
	"strings"
)

/*
GenerateListAllCmds generates a string with a list of all the available commands specified on each CMD_-started
constant, which can be used directly with Main().

-----------------------------------------------------------

> Params:

- none

> Returns:

- a string with all the available commands separated by CMDS_SEPARATOR
*/
func GenerateListAllCmds() string {
	var ret_var string = ""
	for counter := 1; counter <= HIGHEST_CMD_INT; counter++ {
		ret_var += strconv.Itoa(counter) + CMDS_SEPARATOR
	}

	return ret_var[:len(ret_var)-len(CMDS_SEPARATOR)]
}

const ADD_INFO_ERR = ERR_CMD_DETECT + "1"

/*
GetCmdAdditionalInfo returns the value associated with the requested additional information from a CMD_-started
constant.

-----------------------------------------------------------

> Params:

- cmd_constant – the CMD_-started constant

- cmdi_info_index – the index of the wanted information (one of the CMDi_INDEX_INF-started constants)

> Returns:

- the value associated with the requested information (one of the CMDi_INF-started constants)
*/
func GetCmdAdditionalInfo(cmd_constant string, cmdi_info_index int32) string {
	if cmd_info, ok := CMDi_INFO[cmd_constant]; ok {
		return strings.Split(cmd_info, CMDi_INF_SEP)[cmdi_info_index]
	} else {
		// Won't happen - just don't use strings calling this function and use actual constants (and as long as they're
		// on the map, which they must be, and are not only by brain memory leak xD).
		return "Won't happen" // lol
	}
}
