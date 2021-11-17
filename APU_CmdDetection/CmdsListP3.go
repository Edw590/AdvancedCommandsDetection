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

package APU_CmdDetection

////////////////////////////////////////////////////////////////////////////////////
// ATTENTION: keep the format of the below arrays as it is. Each index must correspond to the value of the CMD_-started
// constant corresponding to the array index contents in question.
//
// ----- NEVER EVER REMOVE AN ELEMENT FROM THE ARRAYS UNLESS IT'S THE LAST ONE!!! -----
// If it's to deactivate one of the elements, delete everything from it and put a // to do without the space in the
// front to be seen well that that index is not being used (if it's to be used again, it's a notification on the IDE
// warning about unused slots, so it's good to be there).

var conditions_continue_GL = [...][][][]string{
	{}, // Ignored
	{}, // 1
	{}, // 2
	{}, // 3
	{ // 4
		{{}, {"on","off","wifi","wi-fi"}, {"on","off","wifi","wi-fi"}},
	},
	{ // 5
		{{}, {"on","off","mobile","data","connection"}, {"mobile","data","connection"}, {"on","off","mobile","data","connection"}, {"on","off","connection"}},
	},
	{ // 6
		{{}, {"on","off","bluetooth"}, {"on","off","bluetooth"}},
	},
	{}, // 7
	{ // 8
		{{}, {"flashlight", "lantern", "on", "off"}, {"flashlight", "lantern", "on", "off"}},
	},
	{}, // 9
	{ // 10
		{{}, {"on","off","speaker","speakers"}, {"on","off","speaker","speakers"}},
	},
	{ // 11
		{{}, {"on","off","airplane"}, {"mode","airplane"}, {"on","off","mode"}},
	},
	{ // 12
		{{"how"}, {"is"}, {"battery"}},
		{{"how's", "tell"}, {"battery"}, {}},
	},
	{ // 13
		{{"power","turn"}, {"off","down","phone"}, {"off","down","phone"}},
		{{"shut"}, {"down","phone"}, {"down","phone"}},
		{{"shutdown"}, {"phone"}, {}},
	},
	{}, // 14
	{ // 15
		// This is here cecause only one of the words on the 2nd sub-list of the words_list is mandatory ("take a
		// picture"). No word of the first sub-list is mandatory.
		{{}, {}, {"selfie","picture","photo","photograph"}},
	},
	{ // 16
		// Because only one of the words on the 2nd sub-list of the words_list is mandatory ("record the audio").
		// No word of the first sub-list is mandatory.
		{{}, {}, {"video","camera","audio","sound"}},
	},
	{ // 17
		{{"say"}, {"again"}},
		{{"what"}, {"say","said"}},
	},
	{}, // 18
}

var conditions_not_continue_GL = [...][][][][]string{
	{}, // Ignored
	{}, // 1
	{}, // 2
	{}, // 3
	{ // 4
		{
			{{}, {"off"}, {"on"}},  {{}, {"on"}, {"off"}},
		},
	},
	{ // 5
		{
			{{}, {"on"}, {}, {"off"}},  {{}, {"off"}, {}, {"on"}},
		},
	},
	{ // 6
		{
			{{}, {"off"}, {"on"}},  {{}, {"on"}, {"off"}},
		},
	},
	{}, // 7
	{ // 8
		{
			{{}, {"off"}, {"on"}},  {{}, {"on"}, {"off"}},
		},
	},
	{}, // 9
	{ // 10
		{
			{{}, {"off"}, {"on"}},  {{}, {"on"}, {"off"}},
		},
	},
	{ // 11
		{
			{{}, {"off"}, {}, {"on"}},  {{}, {"on"}, {}, {"off"}},
		},
	},
	{}, // 12
	{}, // 13
	{}, // 14
	{}, // 15
	{ // 16
		// Can't record frontal or rear audios xD.
		{
			{{}, {"frontal","front","rear"}, {"audio","sound"}},
		},
	},
	{ // 17
		{
			{{"say"}, {"say","said"}},
		},
		{
			{{"what"}, {"again"}},
		},
	},
	{}, // 18
}

/*
Each sub-array ("condition") contains a set of conditions that can come in the results. Each array in the sub-array
("sub-condition") in question is related to an index of the results of the wordsVerificationDADi() function call.

---------------

> Format of the conditions:

- Last index – a 1-element array containing the RET_-started constant to return in affirmative case of the result on
that sub-array happening

- All other indexes – see below (they are all the same - the only one different is the last one)

NOTES:

- If there are no results to check (like just in case the function detected the words, everything is fine), use only
the last part (means, for example, put an array with RET_2 alone on the only index of the sub-array).

- If for example
it's needed to check 2 conditions and return one thing, and in case those 2 don't apply but *any* other applies, return
something else, put those 2 conditions on 2 separate arrays and then one below those 2.
Example of this:
		{{"0", "thing1"}, {RET_0_OPTION1}},
		{{"1", "thing2"}, {RET_0_OPTION1}},
		{{RET_0_OPTION2}},
In the example above, it's either OPTION1 by the first 2 conditions, or it's OPTION2 in *any other combination* of
results (this is equivalent to an "else" statement).

---------------

> Format of each sub-condition:

- 1st index – the index of the word on the results from the wordsVerificationDADi(). Use -1 to refer to the main word
that started the command (for example for "set alarm", that word is "set", and will not be in the results from the
function).

- All other indexes ("parameters") – words that the word on the specified index can have in order for the result to be a
success. If the results include an index which is not present on these indexes mentioned here, the word on it will be
ignored. The order of these arrays will be respected, and if one of them gives a positive result, the others will not be
considered.

*/
var conditions_return_GL = [...][][][]string{
	{}, // Ignored
	{ // 1
		{{"-1", "stop"}, {RET_1_STOP}},
		{{"-1", "pause"}, {RET_1_PAUSE}},
		{{"-1", "play", "continue", "resume"}, {RET_1_PLAY}},
		{{"-1", "next"}, {RET_1_NEXT}},
		{{"-1", "previous"}, {RET_1_PREVIOUS}},
	},
	{ // 2
		{{RET_2}},
	},
	{ // 3
		{{RET_3}},
	},
	{ // 4
		{{"0", "on"}, {RET_4_ON}},
		{{"1", "on"}, {RET_4_ON}},

		{{RET_4_OFF}},
	},
	{ // 5
		{{"0", "on"}, {RET_5_ON}},
		{{"2", "on"}, {RET_5_ON}},
		{{"3", "on"}, {RET_5_ON}},

		{{RET_5_OFF}},
	},
	{ // 6
		{{"0", "on"}, {RET_6_ON}},
		{{"2", "on"}, {RET_6_ON}},

		{{RET_6_OFF}},
	},
	{ // 7
		{{RET_7}},
	},
	{ // 8
		{{"0", "on"}, {RET_8_ON}},
		{{"1", "on"}, {RET_8_ON}},

		{{RET_8_OFF}},
	},
	{ // 9
		{{RET_9}},
	},
	{ // 10
		{{"0", "on"}, {RET_10_ON}},

		{{RET_10_OFF}},
	},
	{ // 11
		{{"0", "on"}, {RET_11_ON}},
		{{"2", "on"}, {RET_11_ON}},

		{{RET_11_OFF}},
	},
	{ // 12
		{{RET_12}},
	},
	{ // 13
		{{RET_13}},
	},
	{ // 14
		{{RET_14}},
	},
	{ // 15
		// If done right, the conditions below in the order they are in, should be equivalent to the old conditions:
		// results[0][0] == "frontal" || results[0][0] == "front" || (results[1][0] == "selfie" && results[0][0]!="rear")
		{{"0", "rear"}, {RET_15_REAR}},
		{{"0", "frontal", "front"}, {RET_15_FRONTAL}},
		{{"1", "selfie"}, {RET_15_FRONTAL}},

		{{RET_15_REAR}}, // All the above, or "else", this
	},
	{ // 16
		{{"1", "audio", "sound"}, {RET_16_AUDIO}},

		// Based on the conditions of no continuation, the line below works for a frontal video, as "audio" or "sound"
		// will not be on the other indexes - check the conditions.
		{{"0", "frontal", "front"}, {RET_16_VIDEO_FRONTAL}},
		{{RET_16_VIDEO_REAR}}, // Either is "frontal" or "front", or doesn't matter and it's rear.
	},
	{ // 17
		{{RET_17}},
	},
	{ // 18
		{{RET_18}},
	},
}
