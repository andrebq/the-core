package autoupdate

import (
	"github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	"path/filepath"
)

func appFolder(app string) (string, error) {
	dir, err := homedir.Dir()
	if err != nil {
		return "", errors.Wrapf(err, "unable to get homedir for user")
	}
	return filepath.Join(dir, "opt", "the-core", app), nil
}

func binFolder(appFolder string) string {
	return filepath.Join(appFolder, "bin")
}

func resFolder(appFolder string) string {
	return filepath.Join(appFolder, "res")
}
