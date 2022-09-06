// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

package cmd

import (
	fbcmd "github.com/elastic/beats/v7/filebeat/cmd"
	cmd "github.com/elastic/beats/v7/libbeat/cmd"
	"github.com/elastic/beats/v7/x-pack/libbeat/management"

	// Register the includes.
	_ "github.com/elastic/beats/v7/x-pack/filebeat/include"
	inputs "github.com/elastic/beats/v7/x-pack/filebeat/input/default-inputs"
	_ "github.com/elastic/beats/v7/x-pack/libbeat/include"
)

const Name = fbcmd.Name

// Filebeat build the beat root command for executing filebeat and it's subcommands.
func Filebeat() *cmd.BeatsRootCmd {
	management.ConfigTransform.SetTransform(filebeatCfg)
	settings := fbcmd.FilebeatSettings()
	settings.ElasticLicensed = true
	command := fbcmd.Filebeat(inputs.Init, settings)
	return command
}
