//go:build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

var Default = Build

func Build() error {
	mg.Deps(CheckTmp)
	if err := sh.RunV("go", "mod", "tidy"); err != nil {
		return err
	}

	env := map[string]string{"GO111MODULE": "on"}
	env["CGO_ENABLE"] = "0"

	args := []string{"build"}
	args = append(args, "-o", "./tmp/weaver")
	args = append(args, ".")
	if err := sh.RunWithV(env, "go", args...); err != nil {
		return err
	}
	return nil
}

func CheckTmp() error {
	_, err := os.Stat("tmp")
	if err != nil && err != os.ErrNotExist {
		err = os.Mkdir("tmp", 0755)
	} else if err != nil && err == os.ErrExist {
		return nil
	}
	return err
}
