# coding=utf-8
"""
This file updates this Go module VERSION variable each time it's ran.
"""
import datetime

LINE_BEGINNING = "const VERSION string ="
LOCAL_TIMEZONE = str(datetime.datetime.now(datetime.timezone(datetime.timedelta(0))).astimezone().tzinfo)
DATE_TIME = str(datetime.datetime.now().strftime("%Y-%m-%d -- %H:%M:%S.%f"))

# Do not add the time zone with hours here. It doesn't seem to account for Summer time. Keep the Timezone as the system
# reports it. That one is correct (at least on Windows).
FINAL_LINE = '{} "{} ({})"\n'.format(LINE_BEGINNING, DATE_TIME, LOCAL_TIMEZONE)

with open("APU_GlobalUtils/GL_CONSTS.go", "r", encoding="UTF-8") as file:
	lines = file.readlines()

#print(lines)

for counter, line in enumerate(lines):
	if LINE_BEGINNING in line:
		lines[counter] = FINAL_LINE
		#print(lines[counter])
		break

with open("APU_GlobalUtils/GL_CONSTS.go", "w", encoding="UTF-8") as file:
	file.writelines(lines)
