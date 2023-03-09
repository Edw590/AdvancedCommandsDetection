:: Execute with the current directory as the project root directory

:: Update the module version
python VersionUpdater.py

:: Only put here below the packages that need the exported names visible. The others will be added automatically (if
:: needed?).
:: Keep compressdwarf here. It's default true, but they could change it to default false, so this way it's true for
:: sure.
gomobile bind^
 -target=android^
 -x^
 -v^
 -ldflags="-v -s -w -compressdwarf=true"^
 -o="GeneratedBinaries/AdvancedCommandsDetection.aar"^
 "AdvancedCommandsDetection/ACD"

echo Error code: %ERRORLEVEL%
