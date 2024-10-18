package internal

import (
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

const commandTpl = `
docker buildx build \
	--platform=%s \
	-f Dockerfile.tmp \
	-t %s . \
	--push
`

func Syncx(ctx *cli.Context) error {
	var repo = ctx.Args().First()
	if !strings.Contains(repo, "/") {
		repo = "library/" + repo
	}

	var dst = ctx.String("to") + "/" + repo
	src := ctx.String("from") + "/" + repo

	content := []byte("FROM " + src)
	if err := os.WriteFile("Dockerfile.tmp", content, 0644); err != nil {
		return err
	}

	if err := Execute(commandTpl, ctx.String("platform"), dst); err != nil {
		return err
	}
	return os.Remove("Dockerfile.tmp")
}

func Sync(ctx *cli.Context) error {
	var repo = ctx.Args().First()
	if !strings.Contains(repo, "/") {
		repo = "library/" + repo
	}

	var dst = ctx.String("to") + "/" + repo
	src := ctx.String("from") + "/" + repo

	if err := Execute("docker pull %s", src); err != nil {
		return err
	}
	if err := Execute("docker tag %s %s", src, dst); err != nil {
		return err
	}
	if err := Execute("docker push %s", dst); err != nil {
		return err
	}
	return nil
}
