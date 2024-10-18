package main

import (
	"github.com/lxzan/docker-utils/internal"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	app := cli.App{
		Name:  "docker-utils",
		Usage: "docker辅助工具",
		Commands: []*cli.Command{
			{
				Name:  "syncx",
				Usage: "同步镜像仓库",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "f",
						Aliases: []string{"from"},
						Value:   "docker.io",
						Usage:   "镜像源",
					},
					&cli.StringFlag{
						Name:    "t",
						Aliases: []string{"to"},
						Usage:   "目标地址, 例如 harbor.com/library/alpine:latest",
					},
					&cli.StringFlag{
						Name:        "p",
						Aliases:     []string{"platform"},
						DefaultText: "linux/amd64",
						Usage:       "平台架构, 例如 linux/amd64, 如有多个用逗号分隔",
					},
				},
				Action: internal.Syncx,
			},

			{
				Name:  "sync",
				Usage: "同步镜像仓库",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "f",
						Aliases: []string{"from"},
						Value:   "docker.io",
						Usage:   "镜像源",
					},
					&cli.StringFlag{
						Name:    "t",
						Aliases: []string{"to"},
						Usage:   "目标地址, 例如 harbor.com/library/alpine:latest",
					},
				},
				Action: internal.Sync,
			},
		},
	}
	if err := app.Run(os.Args); err != nil {
		log.Println(err.Error())
	}
}
