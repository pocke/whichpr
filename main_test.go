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
			t.Errorf("Expected: 'Does not match', but got %s", err.Error())
		}
	})
}

func TestIsParent_whenEqual(t *testing.T) {
	if !isParent("51ed8a271", "51ed8a271d02d1c2a1fb5166d0679e23a0436fd7") {
		t.Error(":(")
	}
}

func TestIsParent_whenNotEq(t *testing.T) {
	inGitDir(t, func() {
		err := sh(`
			git tag first-commit
			git checkout -b foobar
			git commit --no-gpg-sign --allow-empty -m 'foo'
			git tag 1br
			git commit --no-gpg-sign --allow-empty -m 'bar'
			git tag 2br
			git checkout master
			git commit --no-gpg-sign --allow-empty -m 'baz'
			git tag 1master
			git commit --no-gpg-sign --allow-empty -m 'piyo'
			git tag 2master
			git merge --no-gpg-sign -m 'merge' foobar
		`)
		if err != nil {
			t.Error(err)
		}
		if !isParent("1br", "2br") {
			t.Error(":(")
		}
		if !isParent("1master", "2master") {
			t.Error(":(")
		}
		if isParent("2br", "1br") {
			t.Error(":(")
		}
		if isParent("2br", "1master") {
			t.Error(":(")
		}
	})
}

// helpers

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

	if e := os.Chdir(dir); e != nil {
		t.Fatal(e)
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
