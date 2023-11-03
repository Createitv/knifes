// Copyright 2023 xfy150150@gmail.com. All rights reserved.
// @Time        : 2023/11/2 21:27
// @Author      : Createitv
// @FileName    : os_test.go.go
// @Software    : GoLand
// @WeChat      : Navalism1
// @Description :

package system

import (
	"testing"
)

func init() {
	SetOsEnv("NUMBER_OF_PROCESSORS", "12")
	SetOsEnv("OS", "Windows_NT")
	SetOsEnv("choco", "123")
}

func TestIsWindows(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"test windows os", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWindows(); got != tt.want {
				t.Errorf("IsWindows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareOsEnv(t *testing.T) {
	type args struct {
		key         string
		comparedEnv string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"compare OS", args{
			key:         "OS",
			comparedEnv: "darwin",
		}, false},
		{"compare NUMBER_OF_PROCESSORS-12", args{
			key:         "NUMBER_OF_PROCESSORS",
			comparedEnv: "13",
		}, false},
		{"compare NUMBER_OF_PROCESSORS-12", args{
			key:         "NUMBER_OF_PROCESSORS",
			comparedEnv: "12",
		}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareOsEnv(tt.args.key, tt.args.comparedEnv); got != tt.want {
				t.Errorf("CompareOsEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetOsEnv(t *testing.T) {
	type args struct {
		key string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"NUMBER_OF_PROCESSORS", args{"NUMBER_OF_PROCESSORS"}, "12"},
		{"Get Os System", args{"OS"}, "Windows_NT"},
		{"choco", args{"choco"}, "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetOsEnv(tt.args.key); got != tt.want {
				t.Errorf("GetOsEnv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsLinux(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"test linux os", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsLinux(); got != tt.want {
				t.Errorf("IsLinux() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsMac(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"test mac os", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsMac(); got != tt.want {
				t.Errorf("IsMac() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveOsEnv(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := RemoveOsEnv(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RemoveOsEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestSetOsEnv(t *testing.T) {
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SetOsEnv(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("SetOsEnv() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
