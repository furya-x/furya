package cli

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"
	"github.com/stafihub/stafihub/x/mining/types"
)

var _ = strconv.Itoa(0)

func CmdStakeRecord() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "stake-record [user-address] [stake-token-denom] [stake-record-index]",
		Short: "Query stake record",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			reqUserAddress := args[0]
			reqStakeTokenDenom := args[1]
			reqStakeRecordIndex, err := sdk.ParseUint(args[2])
			if err != nil {
				return err
			}

			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryStakeRecordRequest{
				UserAddress:      reqUserAddress,
				StakeTokenDenom:  reqStakeTokenDenom,
				StakeRecordIndex: uint32(reqStakeRecordIndex.Uint64()),
			}

			res, err := queryClient.StakeRecord(cmd.Context(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}