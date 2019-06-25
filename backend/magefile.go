// +build mage

package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

func GenerateProtoc() {
	mg.Deps(GenerateStorage, GenerateBridge)
}

func GenerateBridge() error {
	return runCmd(exec.Command("protoc", "-I", "proto", "proto/bridge.proto",
	"--go_out=plugins=grpc:api"))
}

func GenerateStorage() error {
	return runCmd(exec.Command("protoc", "-I", "proto", "proto/storage.proto",
	"--go_out=plugins=grpc:api"))
}

func Generate() error {
	mg.Deps(GenerateProtoc)
	return nil
}

func Build() error {
	return runCmd(exec.Command("go", "build", "-v", "./..."))
}

func Test() error {
	mg.Deps(Build)
	return runCmd(exec.Command("go", "test", "./..."))
}

func Install() error {
	mg.Deps(Build)
	return nil
}

// Manage your deps, or running package managers.
func InstallDeps() error {
	fmt.Println("Installing deps...")
	cmd := exec.Command("go", "get", "-v", ".")
	return cmd.Run()
}

// Dist generates binary artificats for the current operating system
func Dist() {
	mg.Deps(DistStoraged, DistStorageClient)
}

// DistStoraged builds the storage server
func DistStoraged() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", path("./dist/storaged.exe"),
	"./cmd/storaged"))
}

// DistStorageClient builds the storage client
func DistStorageClient() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", path("./dist/the-core-storage-cli.exe"),
	"./cmd/the-core-storage-cli"))
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("dist")
}


func runCmd(cmd *exec.Cmd) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	return cmd.Run()
}

func path(str string) string {
	return filepath.FromSlash(str)
}