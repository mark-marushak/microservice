package serve

import (
	"context"

	"github.com/urfave/cli/v3"
)

func New() *cli.Command {
	return &cli.Command{
		Name:   "serve",
		Usage:  "usage bla bla",
		Action: action,
	}
}

func action(ctx context.Context, c *cli.Command) error {
	return nil
}
