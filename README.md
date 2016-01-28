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

## TODO

The ideal is to have a central config file (like in the user's folder) to define settings like `LogPath`. This could then be used to alway log the commands called via `scroxy`. We can also have settings like `MustLogDuration` which could log the running time of the command.