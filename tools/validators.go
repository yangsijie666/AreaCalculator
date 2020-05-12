package tools

import (
    "AreaCalculator/types"
    "fmt"
    "os"
    "path"
    "runtime"
    "strconv"
)

func ValidateIntStr(str string) (int, error) {
    res, err := strconv.Atoi(str)
    if err != nil {
        return 0, err
    }
    return res, nil
}

func ValidateFontPath(newFontPath string) (string, error) {
    if newFontPath != "" {
        if exist(newFontPath) {
            if path.Ext(newFontPath) != ".ttf" {
                return "", fmt.Errorf(types.FontFileExtError)
            }
            return newFontPath, nil
        }
    }

    // 检查黑体
    if exist(types.DefaultFontDir) {
        return types.DefaultFontDir, nil
    }

    switch runtime.GOOS {
    case "darwin":
        if exist(types.DefaultMacFontDir) {
            return types.DefaultMacFontDir, nil
        }
    case "windows":
        if exist(types.DefaultWindowsFontDir) {
            return types.DefaultWindowsFontDir, nil
        }
    }

    return "", fmt.Errorf(types.CanNotFindFontFile)
}

func exist(filename string) bool {
    _, err := os.Stat(filename)
    return err == nil || os.IsExist(err)
}
