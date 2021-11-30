/*
 * Copyright 2021 DADi590
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
 */

package CommandsDetection_APU

import "strconv"

/*
GenerateListAllCmds generates a string with a list of all the available commands specified on each CMD_-started
constant. This string can be used directly with Main().

-----------------------------------------------------------

> Params:

- none


> Returns:

- a string with all the available commands separated by CMDS_SEPARATOR
*/
func GenerateListAllCmds() string {
	var ret_var string = ""
	for counter := 1; counter < HIGHEST_CMD_INT; counter++ {
		ret_var += strconv.Itoa(counter) + CMDS_SEPARATOR
	}

	return ret_var[:len(ret_var)-len(CMDS_SEPARATOR)]
}
