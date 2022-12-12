// +build !windows
// +build !darwin

// 12 december 2022

package ui

// #cgo LDFLAGS: ${SRCDIR}/libui_linux_arm64.a -lm -ldl
// #cgo pkg-config: gtk+-3.0
import "C"
