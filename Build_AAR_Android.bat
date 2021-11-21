SET ANDROID_HOME=C:\Users\DADi590\AppData\Local\Android\Sdk
SET PATH=%PATH%;"C:\Program Files\Java\jdk-11.0.3\bin"

:: %1 == Project directory
CD %1

:: Only put here below the packages that need the exported names visible.
:: The others will be added automatically (if needed?).
:: Keep GlobalUtils here because it has the VERSION constant, which is useful to know.
:: Also keep compressdwarf here. It's default true, but they could change it to default false, so this way it's true for
:: sure.
gomobile bind^
 --target=android^
 -x^
 -ldflags="-v -s -w -compressdwarf=true"^
 -o="GeneratedBinaries_EOG/assist_platforms_unifier.aar"^
 "Assist_Platforms_Unifier/APU_GlobalUtils"^
 "Assist_Platforms_Unifier/APU_CmdDetection"

::PAUSE
