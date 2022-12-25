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

// Note: CMDi_-started constants have their name due to CMD Information, as they are related to additional command
// information.

// CMDi_INF_SEP is used to split a the value got from the CMDi_INFO map to obtain additional information about the
// command. The constants never end in this separator. After splitting, use one of the CMDi_INDEX_INF-started constants.
const CMDi_INF_SEP string = "/"

// CMDi_INDEX_INF1 is the index to use to get CMDi_INF1_-started values after splitting the CMD Info string.
const CMDi_INDEX_INF1 int32 = 0

// CMDi_INF1_DO_SOMETHING signals that the referring command requires the assistant to do something.
const CMDi_INF1_DO_SOMETHING string = "0_0"

// CMDi_INF1_ONLY_SPEAK signals that the referring command only requires the assistant to say something (like asking
// what time is it).
const CMDi_INF1_ONLY_SPEAK string = "0_1"

// CMDi_INFO a map of each command and its additional information. Use with GetCmdAdditionalInfo().
var CMDi_INFO map[string]string = map[string]string{
	// 0 - Ignored
	CMD_TOGGLE_MEDIA:            CMDi_INF1_DO_SOMETHING, // 1
	CMD_ASK_TIME:                CMDi_INF1_ONLY_SPEAK,   // 2
	CMD_ASK_DATE:                CMDi_INF1_ONLY_SPEAK,   // 3
	CMD_TOGGLE_WIFI:             CMDi_INF1_DO_SOMETHING, // 4
	CMD_TOGGLE_MOBILE_DATA:      CMDi_INF1_DO_SOMETHING, // 5
	CMD_TOGGLE_BLUETOOTH:        CMDi_INF1_DO_SOMETHING, // 6
	CMD_ANSWER_CALL:             CMDi_INF1_DO_SOMETHING, // 7
	CMD_TOGGLE_FLASHLIGHT:       CMDi_INF1_DO_SOMETHING, // 8
	CMD_END_CALL:                CMDi_INF1_DO_SOMETHING, // 9
	CMD_TOGGLE_SPEAKERS:         CMDi_INF1_DO_SOMETHING, // 10
	CMD_TOGGLE_AIRPLANE_MODE:    CMDi_INF1_DO_SOMETHING, // 11
	CMD_ASK_BATTERY_PERCENT:     CMDi_INF1_ONLY_SPEAK,   // 12
	CMD_SHUT_DOWN_DEVICE:        CMDi_INF1_DO_SOMETHING, // 13
	CMD_REBOOT_DEVICE:           CMDi_INF1_DO_SOMETHING, // 14
	CMD_TAKE_PHOTO:              CMDi_INF1_DO_SOMETHING, // 15
	CMD_RECORD_MEDIA:            CMDi_INF1_DO_SOMETHING, // 16
	CMD_SAY_AGAIN:               CMDi_INF1_ONLY_SPEAK,   // 17
	CMD_MAKE_CALL:               CMDi_INF1_DO_SOMETHING, // 18
	CMD_TOGGLE_POWER_SAVER_MODE: CMDi_INF1_DO_SOMETHING, // 19
	CMD_STOP_RECORD_MEDIA:       CMDi_INF1_DO_SOMETHING, // 20
}
