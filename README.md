
# riot-cli

Simple command line client for Riot Games service written in Go.

## Install



### Build

You can build and install the official repository with Go (before running the below command, ensure you have GO111MODULE=on set in your environment):


This will checkout this repository into `$GOPATH/src/yucealiosman/riot-cli/riotcli/`, build, and install it.

It should be available in $GOPATH/bin/riotcli.


### Example Usage

riotcli {command}


##### Commands
* `help`                    Help about any command
* `champion-masteries-score` Champion masteries score by summoner name
* `choose-region`            Choose your region
* `current-game `            Current game by summoner name
* `help`                    Help about any command
* `match`                    Match detail by match id
* `match-list`               Match list by summoner name
* `server-status`            Status of selected server
* `summoners`                Detail of summoners with given name list
* `token`                    Riot api token
* `token-set`                Set your riot api token