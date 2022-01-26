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

////////////////////////////////////////////////////////////////////////////////////
// ATTENTION: keep the format of the below arrays as it is. Each index must correspond to the value of the CMD_-started
// constant corresponding to the array index contents in question.
//
// ----- NEVER EVER REMOVE AN ELEMENT FROM THE ARRAYS UNLESS IT'S THE LAST ONE!!! -----
// If it's to deactivate one of the elements, delete everything from it and put a // to do without the space in the
// front to be seen well that that index is not being used (if it's to be used again, it's a notification on the IDE
// warning about unused slots, so it's good to be there).

var left_intervs_GL = [...]map[string]string{
	{}, // Ignored
	{}, // 1
	{}, // 2
	{}, // 3
	{}, // 4
	{}, // 5
	{}, // 6
	{}, // 7
	{}, // 8
	{}, // 9
	{}, // 10
	{}, // 11
	{}, // 12
	{}, // 13
	{}, // 14
	{}, // 15
	{}, // 16
	{}, // 17
	{}, // 18
	{}, // 19
	{}, // 20
}

var right_intervs_GL = [...]map[string]string{
	{}, // Ignored
	{}, // 1
	{}, // 2
	{}, // 3
	{}, // 4
	{}, // 5
	{}, // 6
	{}, // 7
	{}, // 8
	{}, // 9
	{}, // 10
	{}, // 11
	{}, // 12
	{}, // 13
	{ // 14
		"1": "4", // "reboot the phone |into the amazing| safe mode" (some adjective someone might say?)
	},
	{}, // 15
	{}, // 16
	{}, // 17
	{}, // 18
	{}, // 19
	{}, // 20
}

var init_indexes_sub_verifs_GL = [...]map[string]string{
	{}, // Ignored
	{}, // 1
	{}, // 2
	{}, // 3
	{}, // 4
	{}, // 5
	{}, // 6
	{}, // 7
	{}, // 8
	{}, // 9
	{}, // 10
	{}, // 11
	{}, // 12
	{}, // 13
	{}, // 14
	{}, // 15
	{}, // 16
	{}, // 17
	{}, // 18
	{}, // 19
	{}, // 20
}

var exclude_word_found_GL = [...][]int{
	{},                   // Ignored
	{ALL_SUB_VERIFS_INT}, // 1
	{ALL_SUB_VERIFS_INT}, // 2
	{ALL_SUB_VERIFS_INT}, // 3
	{ALL_SUB_VERIFS_INT}, // 4
	{ALL_SUB_VERIFS_INT}, // 5
	{ALL_SUB_VERIFS_INT}, // 6
	{ALL_SUB_VERIFS_INT}, // 7
	{ALL_SUB_VERIFS_INT}, // 8
	{ALL_SUB_VERIFS_INT}, // 9
	{ALL_SUB_VERIFS_INT}, // 10
	{ALL_SUB_VERIFS_INT}, // 11
	{ALL_SUB_VERIFS_INT}, // 12
	{ALL_SUB_VERIFS_INT}, // 13
	{ALL_SUB_VERIFS_INT}, // 14
	{ALL_SUB_VERIFS_INT}, // 15
	{ALL_SUB_VERIFS_INT}, // 16
	{ALL_SUB_VERIFS_INT}, // 17
	{ALL_SUB_VERIFS_INT}, // 18
	{ALL_SUB_VERIFS_INT}, // 19
	{ALL_SUB_VERIFS_INT}, // 20
}

var return_last_match_GL = [...]bool{
	false, // Ignored
	false, // 1
	false, // 2
	false, // 3
	false, // 4
	false, // 5
	false, // 6
	false, // 7
	false, // 8
	false, // 9
	false, // 10
	false, // 11
	false, // 12
	false, // 13
	false, // 14
	false, // 15
	false, // 16
	false, // 17
	false, // 18
	false, // 19
	false, // 20
}

var ignore_repets_main_words_GL = [...]bool{
	false, // Ignored
	true,  // 1
	true,  // 2
	true,  // 3
	true,  // 4
	true,  // 5
	true,  // 6
	true,  // 7
	true,  // 8
	true,  // 9
	true,  // 10
	true,  // 11
	true,  // 12
	true,  // 13
	true,  // 14
	true,  // 15
	true,  // 16
	true,  // 17
	true,  // 18
	true,  // 19
	true,  // 20
}

var ignore_repets_cmds_GL = [...]bool{
	false, // Ignored
	false, // 1
	false, // 2
	false, // 3
	false, // 4
	false, // 5
	false, // 6
	false, // 7
	false, // 8
	false, // 9
	false, // 10
	false, // 11
	false, // 12
	false, // 13
	false, // 14
	false, // 15
	false, // 16
	false, // 17
	false, // 18
	false, // 19
	false, // 20
}

var order_words_list_GL = [...]bool{
	false, // Ignored
	false, // 1
	false, // 2
	false, // 3
	false, // 4
	false, // 5
	false, // 6
	false, // 7
	false, // 8
	false, // 9
	false, // 10
	false, // 11
	false, // 12
	false, // 13
	false, // 14
	false, // 15
	false, // 16
	false, // 17
	false, // 18
	false, // 19
	false, // 20
}

var stop_first_not_found_GL = [...]bool{
	false, // Ignored
	false, // 1
	false, // 2
	false, // 3
	false, // 4
	false, // 5
	false, // 6
	false, // 7
	false, // 8
	false, // 9
	false, // 10
	false, // 11
	false, // 12
	false, // 13
	false, // 14
	false, // 15
	false, // 16
	false, // 17
	false, // 18
	false, // 19
	false, // 20
}

var exclude_original_words_GL = [...]bool{
	false, // Ignored
	true,  // 1
	true,  // 2
	true,  // 3
	true,  // 4
	true,  // 5
	true,  // 6
	true,  // 7
	true,  // 8
	true,  // 9
	true,  // 10
	true,  // 11
	true,  // 12
	true,  // 13
	true,  // 14
	true,  // 15
	true,  // 16
	true,  // 17
	true,  // 18
	true,  // 19
	true,  // 20
}

var continue_with_words_slice_number_GL = [...]int{
	0,  // Ignored
	-1, // 1
	-1, // 2
	-1, // 3
	-1, // 4
	-1, // 5
	-1, // 6
	-1, // 7
	-1, // 8
	-1, // 9
	-1, // 10
	-1, // 11
	-1, // 12
	-1, // 13
	-1, // 14
	-1, // 15
	-1, // 16
	-1, // 17
	-1, // 18
	-1, // 19
	-1, // 20
}
