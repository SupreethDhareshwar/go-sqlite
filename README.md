# Go Sqlite Setup

## Pre-requisites

- Install Git `sudo apt-get install git`
- Configure user.name and user.email
  `git config --global user.email "you@example.com"`
  `git config --global user.name "Your Name"`
- Setup Github Desktop
  Go to https://github.com/shiftkey/desktop/releases and download latest release
  To install `sudo dpkg -i <file>.deb`
- Install Docker and Docker-Compose
- Setup Go (We are using 1.12) at ~/go
- Setup Env variables like GOPATH and GOROOT in ~/.bashrc
  `export GOPATH=$HOME/.go`
  `export PATH="$PATH:$HOME/go/bin:$GOROOT/bin:$GOPATH/bin"`
- Setup Project at ~/.go/src/
- Install Go Plugin for VSCode
- Install dep package to manage all project package dependencies
- `$ cd ~/.go/src/go-sqlite`
- `dep init` - This creates Gopkg.toml,Gopkg.lock and vendor folder
- Project is all setup, time to start installing packages and start coding

## Setup Project

- Create a new file main.go. This will be a satrting point of application

- Create a folder model. This will store all your application model files
