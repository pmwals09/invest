package rebalance

import (
  "fmt"
)

type Asset struct {
	Symbol string
	Class  string
	Amount string // TODO: This should be an arbitrary-precision float
}

type Target struct {
	Class   string
	Percent string // TODO: This should be an arbitrary-precision float
}

type AssetClass struct{}

func RebalancePortfolio(
	portfolioFile string,
	targetFile string,
	rawAssets []string,
	rawAssetTargets []string,
) {
	portfolio := buildPortfolio(portfolioFile, rawAssets)
	targets := buildTargets(targetFile, rawAssetTargets)

  fmt.Println(portfolio)
  fmt.Println(targets)
}

func buildPortfolio(portfolioFile string, rawAssets []string) []Asset {
	if portfolioFile == "" {
		res := []Asset{}
		for _, s := range rawAssets {
			res = append(res, parseAssetString(s))
		}
		return res
	}
	assetStrings := getStringsFromFile(portfolioFile)
	res := []Asset{}
	for _, s := range assetStrings {
		res = append(res, parseAssetString(s))
	}
	return res
}

func parseAssetString(s string) Asset {
	return Asset{}
}

func getStringsFromFile(path string) []string {
	return []string{}
}

func buildTargets(targetFile string, rawAssetTargets []string) []Target {
	if targetFile == "" {
		res := []Target{}
		for _, s := range rawAssetTargets {
			res = append(res, parseTargetString(s))
		}
		return res
	}
	targetStrings := getStringsFromFile(targetFile)
	res := []Target{}
	for _, s := range targetStrings {
		res = append(res, parseTargetString(s))
	}
	return res
}

func parseTargetString(s string) Target {
  return Target{}
}
