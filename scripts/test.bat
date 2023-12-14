@echo off
echo Testing Sitri!
go build
cd test
del * /S /Q
rmdir /S /Q .
..\sitri.exe init
echo If the test directory has a .sitri directory containing everything it should have, it worked!ubuntu