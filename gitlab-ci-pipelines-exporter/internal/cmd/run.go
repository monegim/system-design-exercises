package cmd

import "github.com/urfave/cli/v2"

func Run(cliCtx *cli.Context) (int, error) {
	cfg, err := configure(cliCtx)
	if err != nil {
		return 1, err
	}
	_ = cfg
	return 0, nil
}
