# Copyright 2021 DADi590
#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.

# coding=utf-8
"""
This file updates this module's VERSION variable each time it's ran.
"""
import datetime

LINE_BEGINNING = "const VERSION string ="
LOCAL_TIMEZONE = str(datetime.datetime.now(datetime.timezone(datetime.timedelta(0))).astimezone().tzinfo)
DATE_TIME = str(datetime.datetime.now().strftime("%Y-%m-%d -- %H:%M:%S.%f"))

# Do not add the time zone with hours here. It doesn't seem to account for Summer time. Keep the Timezone as the system
# reports it. That one is correct (at least on Windows).
FINAL_LINE = '{} "{} ({})"\n'.format(LINE_BEGINNING, DATE_TIME, LOCAL_TIMEZONE)

with open("GlobalUtils_APU/GL_CONSTS.go", "r", encoding="UTF-8") as file:
	lines = file.readlines()

#print(lines)

for counter, line in enumerate(lines):
	if LINE_BEGINNING in line:
		lines[counter] = FINAL_LINE
		#print(lines[counter])
		break

with open("GlobalUtils_APU/GL_CONSTS.go", "w", encoding="UTF-8") as file:
	file.writelines(lines)
