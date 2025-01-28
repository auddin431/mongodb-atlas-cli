// Copyright 2025 MongoDB Inc
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

package processor

import (
	"context"
	"fmt"

	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/cli/require"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/config"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/store"
	"github.com/mongodb/mongodb-atlas-cli/atlascli/internal/usage"
	"github.com/spf13/cobra"
)

type StopOpts struct {
	cli.ProjectOpts
	cli.OutputOpts
	cli.StreamsOpts
	processorName string
	store         store.ProcessorStopper
}

func (opts *StopOpts) Run() error {
	err := opts.store.StopStreamProcessor(opts.ConfigProjectID(), opts.Instance, opts.processorName)
	if err != nil {
		return err
	}

	return opts.Print("Successfully stopped Stream Processor")
}

func (opts *StopOpts) initStore(ctx context.Context) func() error {
	return func() error {
		var err error
		opts.store, err = store.New(store.AuthenticatedPreset(config.Default()), store.WithContext(ctx))
		return err
	}
}

// atlas streams processor stop <processorName>.
func StopBuilder() *cobra.Command {
	opts := &StopOpts{}
	cmd := &cobra.Command{
		Use:   "stop <processorName>",
		Short: "Stop an Atlas Stream Processor in a Stream Processing Instance.",
		Long:  fmt.Sprintf(usage.RequiredOneOfRoles, streamsRoles),
		Example: `  # stop Stream Processor 'ExampleProcessor' for an instance 'ExampleInstance' for the project with ID 5e2211c17a3e5a48f5497de3:
  atlas streams processors stop ExampleProcessor --projectId 5e2211c17a3e5a48f5497de3 --instance ExampleInstance`,
		Args: require.ExactArgs(1),
		Annotations: map[string]string{
			"processorNameDesc": "Name of the Stream Processor",
		},
		PreRunE: func(cmd *cobra.Command, args []string) error {
			if err := opts.PreRunE(
				opts.ValidateProjectID,
				opts.ValidateInstance,
				opts.initStore(cmd.Context()),
			); err != nil {
				return err
			}
			opts.processorName = args[0]
			return nil
		},
		RunE: func(_ *cobra.Command, _ []string) error {
			return opts.Run()
		},
	}

	opts.AddProjectOptsFlags(cmd)
	opts.AddStreamsOptsFlags(cmd)

	return cmd
}
