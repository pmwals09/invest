package cmd

import (
	"github.com/pmwals09/invest/internal/rebalance"
	"github.com/spf13/cobra"
)

var portfolioFile string
var targetFile string
var rawAssets []string
var rawAssetTargets []string

var rebalanceCmd = &cobra.Command{
	Use:   "rebalance",
	Short: "Input a portfolio and desired weights, receive instructions to achieve the weighting.",
	Long: `Input an existing asset portfolio, either with a properly-formatted file
or with command-line flag, as well as desired weighting of those asset types and
receive the BUY and SELL instructions required to achieve that target.`,
	Run: func(cmd *cobra.Command, args []string) {
    rebalance.RebalancePortfolio(portfolioFile, targetFile, rawAssets, rawAssetTargets)
	},
}

func init() {
	rootCmd.AddCommand(rebalanceCmd)

  rebalanceCmd.Flags().StringVarP(
    &portfolioFile,
    "portfolio-file",
    "f",
    "",
    "Formatted portfolio file",
  )
  rebalanceCmd.Flags().StringArrayVarP(
    &rawAssets,
    "asset",
    "a",
    []string{},
    "An existing asset as SYMBOL,ASSETCLASS,AMOUNT. Repeat for each asset",
  )
  rebalanceCmd.MarkFlagsMutuallyExclusive("portfolio-file", "asset")

  rebalanceCmd.Flags().StringVarP(
    &targetFile,
    "target-file",
    "g",
    "",
    "Formatted target file",
  )
  rebalanceCmd.Flags().StringArrayVarP(
    &rawAssetTargets,
    "target",
    "t",
    []string{},
    "An asset class target as ASSETCLASS,PERCENT. Repeat for each target",
  )
  rebalanceCmd.MarkFlagsMutuallyExclusive("target-file", "target")
}
