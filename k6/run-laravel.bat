@echo off
REM Run load test against Laravel server (port 8000)
echo Starting load test against Laravel server (http://localhost:8000)...
k6 run --env BASE_URL=http://localhost:8000 %~dp0\load-test.js
