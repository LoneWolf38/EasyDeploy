# EasyDeploy


## Automatically Deploys your Website/Webapps
- Creates a EC2 instance with all the necessary softwares
- Clones your project from Github and deploys it.
- If your project needs an database, EasyDeploys creates a AWS Database server for you to use
- If your want your project/website to be linked with a domain, EasyDeploys attaches the given domain to the server

## Build From Source
```
$ git clone https://github.com/LoneWolf38/EasyDeploy.git /home/$USER/go/github.com/LoneWolf38/.

$ cd /home/$USER/go/github.com/LoneWolf38/EasyDeploy/

$ go build -o EasyDeploy main.go

$ ./EasyDeploy

```

## Basic Usage (Only One Command is now working)
```
$ ./EasyDeploy init [--region][--key]

```