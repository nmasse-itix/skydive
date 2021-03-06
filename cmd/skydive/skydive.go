/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package skydive

import (
	"os"
	"path"
	"strings"

	"github.com/skydive-project/skydive/cmd"
	"github.com/skydive-project/skydive/cmd/agent"
	"github.com/skydive-project/skydive/cmd/allinone"
	"github.com/skydive-project/skydive/cmd/analyzer"
	"github.com/skydive-project/skydive/cmd/client"
	"github.com/skydive-project/skydive/cmd/completion"
	"github.com/skydive-project/skydive/cmd/config"
	"github.com/skydive-project/skydive/cmd/version"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:          "skydive [sub]",
	Short:        "Skydive",
	SilenceUsage: true,
	PersistentPreRun: func(c *cobra.Command, args []string) {
		config.LoadConfiguration(cmd.CfgBackend, cmd.CfgFiles)
	},
}

func init() {
	RootCmd.PersistentFlags().StringArrayVarP(&cmd.CfgFiles, "conf", "c", []string{}, "location of Skydive configuration files, default try loading /etc/skydive/skydive.yml if exist")
	RootCmd.PersistentFlags().StringVarP(&cmd.CfgBackend, "config-backend", "b", "file", "configuration backend (defaults to file)")

	executable := path.Base(os.Args[0])
	if strings.TrimSuffix(executable, path.Ext(executable)) == "skydive-cli" {
		RootCmd.Use = "skydive-cli"
		RootCmd.Short = "Skydive client"
		RootCmd.AddCommand(completion.BashCompletion)
		RootCmd.AddCommand(version.VersionCmd)
		client.RegisterClientCommands(RootCmd)
	} else {
		RootCmd.AddCommand(agent.AgentCmd)
		RootCmd.AddCommand(analyzer.AnalyzerCmd)
		RootCmd.AddCommand(completion.BashCompletion)
		RootCmd.AddCommand(client.ClientCmd)
		RootCmd.AddCommand(version.VersionCmd)

		if allinone.AllInOneCmd != nil {
			RootCmd.AddCommand(allinone.AllInOneCmd)
		}
	}
}
