@echo off
chcp 65001 >nul
echo ========================================
echo   Efunc DLL 构建脚本
echo ========================================
echo.

set CGO_ENABLED=1
set GOOS=windows

echo [1/2] 清理旧文件...
if exist Efunc.dll del Efunc.dll
if exist Efunc.h del Efunc.h

echo [2/2] 编译 DLL (x86_64)...
go build -buildmode=c-shared -o Efunc.dll .

if %ERRORLEVEL% NEQ 0 (
    echo.
    echo [错误] 编译失败！请检查：
    echo   1. 是否安装了 GCC (MinGW-w64)
    echo   2. CGO 是否启用
    echo   3. 依赖是否完整 (go mod tidy)
    pause
    exit /b 1
)

echo.
echo ========================================
echo   编译成功！
echo   生成文件：
echo     - Efunc.dll
echo     - Efunc.h
echo ========================================
echo.
echo 使用方法：
echo   1. 将 Efunc.dll 复制到目标项目
echo   2. 包含 Efunc.h 头文件
echo   3. 调用 Efunc_Call(函数名, JSON参数)
echo   4. 调用完毕后用 Efunc_Free 释放内存
echo.
pause
