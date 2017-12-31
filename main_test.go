package main

import (
	"io/ioutil"
	"os"
	"os/exec"
	"testing"
)

func TestSquashedPullReqNum(t *testing.T) {
	inGitDir(t, func() {
		err := sh(`
			git commit --no-gpg-sign -m "Squashed commit (#42)" --allow-empty
		`)
		if err != nil {
			t.Fatal(err)
		}

		pr, err := SquashedPullReqNum("@")
		if err != nil {
			t.Fatal(err)
		}
		if pr != 42 {
			t.Errorf("PR number mismatch! Expected: %d, Got: %d", 42, pr)
		}

	})

	inGitDir(t, func() {
		err := sh(`
			git commit --no-gpg-sign -m "Not squashed commit. (#42)foo" --allow-empty
		`)
		if err != nil {
			t.Fatal(err)
		}

		_, err = SquashedPullReqNum("@")
		if err == nil {
			t.Error("Expected: error, but got nil")
		}
		if err.Error() != "Does not match" {
			t.Error("Expected: 'Does not match', but got %s", err.Error())
		}
	})
}

func inGitDir(t *testing.T, f func()) {
	// Prepare a git repository
	dir, err := ioutil.TempDir("", "whichpr-test-")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	if err := os.Chdir(dir); err != nil {
		t.Fatal(err)
	}
	defer os.Chdir(pwd)

	err = sh(`
		git init .
		git commit --no-gpg-sign -m "First commit" --allow-empty
	`)
	if err != nil {
		t.Fatal(err)
	}

	f()
}

func sh(shell string) error {
	return exec.Command("sh", "-c", shell).Run()
}
