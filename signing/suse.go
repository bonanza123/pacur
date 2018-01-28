package signing

import (
	"fmt"
	"github.com/dropbox/godropbox/errors"
	"github.com/pacur/pacur/utils"
	"os/user"
	"path/filepath"
)

func CreateSuseConf() (err error) {
	name, err := GetName()
	if err != nil {
		return
	}

	data := fmt.Sprintf("%%_signature gpg\n%%_gpg_name %s\n", name)

	usr, err := user.Current()
	if err != nil {
		err = &LookupError{
			errors.Wrap(err, "signing: Failed to get current user"),
		}
		return
	}

	err = utils.CreateWrite(filepath.Join(usr.HomeDir, ".rpmmacros"), data)
	if err != nil {
		return
	}

	return
}

func SignSuse(dir string) (err error) {
	err = CreateSuseConf()
	if err != nil {
		return
	}

	pkgs, err := utils.FindExt(dir, ".rpm")
	if err != nil {
		return
	}

	for _, pkg := range pkgs {
		err = utils.Exec("", "expect",
			"-c", "spawn rpm --resign "+pkg,
			"-c", `expect "Enter pass phrase:"`,
			"-c", `send "\r"`,
			"-c", "wait",
			"-c", "interact")
		if err != nil {
			return
		}
	}

	return
}
