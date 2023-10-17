# Invest

A toolkit for managing your investments.

## Commands

- [`rebalance`](#Rebalance)
- [`status`](#Status)

### Rebalance

Generate a series of BUY and SELL instructions to move a portfolio from a balance of asset classes toward a target "ideal" asset class.

Takes as inputs a target spec and an existing portfolio that matches that spec.

A target or portfolio spec can be input either as a file or as command line arguments.

#### Subcommands

The Rebalance command provides additional utilities for simplifying interactions with this tool:

- `normalize-file` provides a tool that writes a tool-compatible CSV file

### Status

Calculate some basic statistics and charts to provide a high-level overview of portfolio performance.

Takes as input a portfolio spec or a portfolio history.

The portfolio spec or portfolio history can be input as a file or as command line arguments.

#### Subcommands

The Status command provides additional utilities for simplifying interactions with this tool:

- `normalize-file` provides a tool that writes a tool-compatible CSV file
