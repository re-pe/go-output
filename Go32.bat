@echo off
::chcp 65001
set WDIR=%cd%\..\..
set BITS=32
set PROJ_HOME=C:\Projects
set GITPATH=C:\Tools\Git\cmd;C:\Tools\Git\bin
set GOROOT=C:\Tools\Go%BITS%
set GOARCH=386
set GOPATH=%WDIR%;%PROJ_HOME%\Go
PROMPT ----------------------------------------------------$_Go%BITS%: $p$g 
C:\Windows\system32\cmd.exe /K path %GOROOT%\bin;%GITPATH%;%PATH%
@echo on