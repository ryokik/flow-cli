/*
 * Flow CLI
 *
 * Copyright 2019 Dapper Labs, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package migration

import (
	"context"
	"fmt"

	"github.com/onflow/cadence"
	flowsdk "github.com/onflow/flow-go-sdk"
	"github.com/onflow/flowkit"
	"github.com/onflow/flowkit/output"
	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/internal/command"
	"github.com/onflow/flow-cli/internal/scripts"
)

var listStagedContractsflags = scripts.Flags{}

var listStagedContractsCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "flow list-staged <CONTRACT_ADDRESS>",
		Short:   "returns back the a list of staged contracts given a contract address",
		Example: `flow list-staged 0xhello`,
		Args:    cobra.MinimumNArgs(1),
	},
	Flags: &listStagedContractsflags,
	RunS:  listStagedContracts,
}

func listStagedContracts(
	args []string,
	globalFlags command.GlobalFlags,
	_ output.Logger,
	flow flowkit.Services,
	state *flowkit.State,
) (command.Result, error) {
	code, err := RenderContractTemplate(GetStagedCodeForAddressScriptFilepath, globalFlags.Network)
	if err != nil {
		return nil, fmt.Errorf("error loading staging contract file: %w", err)
	}

	contractAddress := args[0]

	caddr := cadence.NewAddress(flowsdk.HexToAddress(contractAddress))

	query := flowkit.ScriptQuery{}
	if listStagedContractsflags.BlockHeight != 0 {
		query.Height = listStagedContractsflags.BlockHeight
	} else if listStagedContractsflags.BlockID != "" {
		query.ID = flowsdk.HexToID(listStagedContractsflags.BlockID)
	} else {
		query.Latest = true
	}

	value, err := flow.ExecuteScript(
		context.Background(),
		flowkit.Script{
			Code: code,
			Args: []cadence.Value{caddr},
		},
		query,
	)
	if err != nil {
		return nil, err
	}

	return scripts.NewScriptResult(value), nil
}
