
# riot-cli

Simple command line client for Riot Games service written in Go.

## Install



### Build

You can build and install the official repository with [Go](https://golang.org/dl/) (before running the below command, ensure you have `GO111MODULE=on` set in your environment):

	go get github.com/yucealiosman/riot-cli/riotcli

This will checkout this repository into `$GOPATH/src/yucealiosman/riot-cli/riotcli`, build, and install it.

It should then be available in $GOPATH/bin/riotcli.


### Example Usage

jiracli {command}


##### Commands
Where the individual commands are maps with these keys:
* `help: string` This is help message displayed in the usage for the command