package util

import (
	"crypto/sha256"
	"fmt"
	"path"
	"runtime"
)

type Level int

const (
	LevelSelf = iota + 1
	LevelParent
	Level3
	Level4
	Level5
)

func GetCallerID(skip Level) string {
	pc, file, lineNo, ok := runtime.Caller(int(skip))
	if !ok {
		return "unknow id"
	}
	funcName := runtime.FuncForPC(pc).Name()
	str := fmt.Sprintf("FuncName:%s, file:%s, line:%d ", funcName, file, lineNo)
	sha := sha256.New()
	sha.Write([]byte(str))
	return fmt.Sprintf("%x", sha.Sum(nil))
}

func GetCallerFile(skip Level) string {
	_, file, _, ok := runtime.Caller(int(skip))
	if !ok {
		return ""
	}
	fileName := path.Base(file)
	return fileName
}
