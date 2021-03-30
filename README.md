# CLI_Web_Server

### Command Line interface which is used to serve specified html files on local server.

## Available commands:

* help - use if you need any explanation of commands on your machine

* version - returns version of currently installed CLI

* run --file <file_name> - specify html file name to get its context on your local server

#### Web server starts after running 'go run cmd/cli/server_cli.go (or just built binary file) run'  command, '--file' flag is required here. 
#### Example 'Hello World!' html file is located in web/ directory
#### To run this program you need package: https://github.com/urfave/cli
