package evm

import (
	"context"
	_ "embed"
	"fmt"

	"github.com/onflow/cadence"
	"github.com/spf13/cobra"

	"github.com/onflow/flow-cli/flowkit"
	"github.com/onflow/flow-cli/flowkit/arguments"
	"github.com/onflow/flow-cli/flowkit/output"
	"github.com/onflow/flow-cli/internal/command"
)

//go:embed get.cdc
var getCode []byte

type flagsGet struct{}

var getFlags = flagsGet{}

var getCommand = &command.Command{
	Cmd: &cobra.Command{
		Use:     "get-account <evm address>",
		Short:   "Get account by the EVM address",
		Args:    cobra.ExactArgs(1),
		Example: "flow evm get-account 522b3294e6d06aa25ad0f1b8891242e335d3b459",
	},
	Flags: &getFlags,
	RunS:  get,
}

// todo only for demo, super hacky now

func get(
	args []string,
	_ command.GlobalFlags,
	_ output.Logger,
	flow flowkit.Services,
	state *flowkit.State,
) (command.Result, error) {
	val, _ := GetEVMAccount(args[0], flow)

	fmt.Printf("\n🔥🔥🔥🔥🔥🔥🔥 EVM Account Creation Summary 🔥🔥🔥🔥🔥🔥🔥\n")
	fmt.Println("Address:  ", "0000000000000000000000000000000000000001")
	fmt.Println("Balance:  ", val)
	fmt.Printf("\n-------------------------------------------------------------\n\n")
	return nil, nil
}

func GetEVMAccount(
	address string,
	flow flowkit.Services,
) (cadence.Value, error) {

	scriptArgs, err := arguments.ParseWithoutType([]string{address}, getCode, "")
	if err != nil {
		return nil, err
	}

	return flow.ExecuteScript(
		context.Background(),
		flowkit.Script{
			Code: getCode,
			Args: scriptArgs,
		},
		flowkit.ScriptQuery{Latest: true},
	)
}