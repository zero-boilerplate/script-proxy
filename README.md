# scroxy (short script-proxy)
Reduce the amount of repeated lines in script (windows batch, shell bash, etc)

## Example

Simple example in windows batch file. Lets assume you have a batch file to build and run your go app.


**run.bat**:
```
@echo off
SET ERRORLEVEL=0

scroxy go build -o my-app.exe     & if errorlevel 1 goto ERROR
scroxy my-app.exe                 & if errorlevel 1 goto ERROR


goto SUCCESS

:SUCCESS
echo Success!!
goto EOF



:ERROR
echo ERROR!!! See the last ran command
pause
goto EOF

:EOF
```

## Features

Defaults to load the config file at location `EXE_PATH``.toml`. If the config file does not exist, it defaults to log path at `EXE_PATH``.log`. The log path is used with [logrus](github.com/Sirupsen/logrus) and its `JSONFormatter` to write logs. For now it writes when the command STARTS and when it ENDS. It also catches the command's `Stdout` and `Stderr` in this log file.