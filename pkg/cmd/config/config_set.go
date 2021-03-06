package config

import (
	"github.com/dcos/dcos-cli/api"
	"github.com/spf13/cobra"
)

// newCmdConfigSet creates the `dcos config set` subcommand.
func newCmdConfigSet(ctx api.Context) *cobra.Command {
	return &cobra.Command{
		Use:   "set",
		Short: "Add or set a property in the configuration file used for the current cluster",
		Long:  "The properties that can be set are: core.dcos_url, core.dcos_acs_token, core.ssl_verify, core.timeout, core.ssh_user, core_ssh_proxy_ip, core.pagination, core.reporting, core.mesos_master_url, core_prompt_login",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := ctx.ConfigManager().Current()
			if err != nil {
				return err
			}

			conf.Set(args[0], args[1])
			return conf.Persist()
		},
	}
}
