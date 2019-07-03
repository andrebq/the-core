// +build mage

package main

import (
	"fmt"
	"strings"
	"runtime"
	"os"
	"os/exec"
	"path/filepath"
	"path"
	"text/template"
	"io/ioutil"
	"bytes"

	"github.com/magefile/mage/mg" // mg contains helpful utility functions, like Deps
)

var (
	goos = runtime.GOOS
	goarch = runtime.GOARCH
)

func TargetWindows() {
	goos = "windows"
}

func TargetLinux() {
	goos = "linux"
}

func GenerateProtoc() {
	mg.Deps(GenerateStorage, GenerateNetport)
}

func GenerateNetport() error {
	return runCmd(exec.Command("protoc", "-I", "proto", "proto/netport.proto",
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

func Fmt() error {
	return runCmd(exec.Command("go", "fmt", "-x", "./..."))
}

func Build() error {
	mg.Deps(Fmt)
	return runCmd(exec.Command("go", "build", "-v", "-x", "./..."))
}

func Test() error {
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
	mg.Deps(DistStoraged, DistStorageClient, DistRelayServer, DistRelayTestClient,
		DistAutoupdate)
}

// Publish the binaries
func Publish() {
	mg.Deps(Dist, WriteManifest)
}

// PublishLinux is self explanatory
func PublishLinux() {
	mg.Deps(TargetLinux, Publish)
}

// WriteManifest to the dist folder
func WriteManifest() error {
	replace := struct{
		AppsUrlPrefix string
	} {
		AppsUrlPrefix: "http://localhost:8081",
	}
	buf := bytes.Buffer{}
	err := template.Must(template.New("apps").Parse(appsJson)).Execute(&buf, replace)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(toFilepath("./dist/apps.json"), buf.Bytes(), 0644)
}

// DistStoraged builds the storage server
func DistStoraged() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", appPath("storaged", goos, goarch),
	"./cmd/storaged"))
}

// DistStorageClient builds the storage client
func DistStorageClient() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", appPath("the-core-storage-cli", goos, goarch),
	"./cmd/the-core-storage-cli"))
}

// DistRelayServer builds the relay server
func DistRelayServer() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", appPath("relay-server", goos, goarch),
	"./cmd/relay-server"))
}

// DistRelayTestClient builds the ping/pong over relay
func DistRelayTestClient() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", appPath("relay-test-client", goos, goarch),
	"./cmd/relay-test-client"))
}

// DistAutoupdate builds the ping/pong over relay
func DistAutoupdate() error {
	mg.Deps(Test)
	return runCmd(exec.Command("go", "build", "-v", "-o", appPath("autoupdate", goos, goarch),
	"./cmd/autoupdate"))
}

// Clean up after yourself
func Clean() {
	fmt.Println("Cleaning...")
	os.RemoveAll("dist")
}


func runCmd(cmd *exec.Cmd, vars ...string) error {
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	cmd.Env = vars
	return cmd.Run()
}

func appPath(bin, goos, goarch string) string {
	if goos == "windows" {
		bin = bin + ".exe"
	}
	return toFilepath(path.Join(
		".", "dist", strings.ToLower(goos), strings.ToLower(goarch), bin))
}

func toFilepath(str string) string {
	return filepath.FromSlash(str)
}

const (
	appsJson = `
{
    "apps": [
        {
            "application": "autoupdate",
            "version": "0.0.0",
            "bundles":[
                {
                    "name": "main",
                    "kind": "bin",
                    "os": "windows",
                    "arch": "amd64",
                    "downloadUrl": "{{.AppsUrlPrefix}}/windows/amd64/autoupdate.exe"
                },
                {
                    "name": "main",
                    "kind": "bin",
                    "os": "linux",
                    "arch": "amd64",
                    "downloadUrl": "{{.AppsUrlPrefix}}/linux/amd64/autoupdate"
                }
            ]
        }
    ]
}
	`
)