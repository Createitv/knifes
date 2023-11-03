// Copyright 2023 panghuang xfy150150@gmail.com. All rights reserved.
// Use of this source code is governed by MIT license.

// Package system
package system

import (
	"os"
	"runtime"
)

// IsWindows checks whether the current system OS is windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

func IsMac() bool {
	return runtime.GOOS == "darwin"
}

func IsLinux() bool {
	return runtime.GOOS == "linux"
}

func GetOsArch() string {
	return runtime.GOARCH
}

func GetOsEnv(key string) string {
	return os.Getenv(key)
}

func SetOsEnv(key, value string) error {
	return os.Setenv(key, value)
}

func RemoveOsEnv(key string) error {
	return os.Unsetenv(key)
}

func CompareOsEnv(key, comparedEnv string) bool {
	env := GetOsEnv(key)
	if env == "" {
		return false
	}
	return env == comparedEnv
}
