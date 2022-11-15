// Copyright 2022 OnMetal authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package orictl

import (
	goflag "flag"

	"github.com/onmetal/onmetal-api/orictl/cmd/orictl/common"
	"github.com/onmetal/onmetal-api/orictl/cmd/orictl/create"
	"github.com/onmetal/onmetal-api/orictl/cmd/orictl/delete"
	"github.com/onmetal/onmetal-api/orictl/cmd/orictl/get"
	"github.com/onmetal/onmetal-api/orictl/cmd/orictl/update"
	"github.com/spf13/cobra"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

func Command(streams common.Streams) *cobra.Command {
	var (
		zapOpts       zap.Options
		dialOpts      common.DialOptions
		clientFactory = dialOpts.Dial
	)

	cmd := &cobra.Command{
		Use: "orictl",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger := zap.New(zap.UseFlagOptions(&zapOpts))
			ctrl.SetLogger(logger)
			cmd.SetContext(ctrl.LoggerInto(cmd.Context(), ctrl.Log))
		},
	}

	goFlags := goflag.NewFlagSet("", 0)
	zapOpts.BindFlags(goFlags)

	cmd.PersistentFlags().AddGoFlagSet(goFlags)
	dialOpts.AddFlags(cmd.PersistentFlags())

	cmd.AddCommand(
		get.Command(streams, clientFactory),
		create.Command(streams, clientFactory),
		delete.Command(streams, clientFactory),
		update.Command(streams, clientFactory),
	)

	return cmd
}
