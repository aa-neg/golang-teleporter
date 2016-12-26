
@echo off
setlocal enabledelayedexpansion
(teleporter %*) > tp_tmpfile.txt
if %errorlevel% equ 2 (
    for /f "delims=" %%A in ('cat tp_tmpfile.txt') do set output=%%A
    del tp_tmpfile.txt
    cd /d !output!
    pushd .
    endlocal
    popd
) else (
    type tp_tmpfile.txt
    del tp_tmpfile.txt
)
