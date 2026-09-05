@echo off
REM Run load test against Go server (port 8080)
echo Starting load test against Go server (http://localhost:8080)...
k6 run --env BASE_URL=http://localhost:8080 %~dp0\load-test.js
