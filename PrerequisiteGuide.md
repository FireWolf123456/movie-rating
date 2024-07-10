# Prerequisite Guide

This guide will walk you through how to install and verify the required tools for the interview coding exercise. 
It will also confirm that you can run the exercise applications locally.
Please complete this guide prior to the interview. Total completion time is usually no more than one hour. This can vary based on how many tools are already installed on your machine.

You'll need to be able to run a golang API and a react app on your local machine in order to complete the coding exercise.

Formatted `text like this` represents a command to run in your terminal.

## Install Tools/Languages

### 1) Install your preferred code editor or IDE
The following are list of good ones but feel free to use whatever you would like:

- [Webstorm](https://www.jetbrains.com/webstorm/)
- [GoLand](https://www.jetbrains.com/go/)
- [VSCode](https://code.visualstudio.com)

### 2) Install NPM
#### Mac
1. Install Homebrew: https://docs.brew.sh/Installation
2. `brew update`
3. `brew doctor`
4. `export PATH="/usr/local/bin:$PATH"` _This adds Homebrew’s location to your $PATH in your .bash_profile or .zshrc file and may not be needed._
5. Install Node: `brew install node`

#### Windows and alternatively Mac
1. https://nodejs.org/en/download/
2. Accept "Automatically install necessary tools"

- Note: Alternatively you can install homebrew for windows and `brew install node` but node installation is easier for windows than brew

### 3) Install Docker
1. Docker installation: https://docs.docker.com/get-docker/

2. **For Windows Machines Only:** To run docker you may need to install the linux kernel update package in step 4: https://docs.microsoft.com/en-us/windows/wsl/install-manual#step-4---download-the-linux-kernel-update-package

### 4) Run dev container
- [Webstorm](https://www.jetbrains.com/help/idea/connect-to-devcontainer.html#create_dev_container_inside_ide)
- [GoLand](https://www.jetbrains.com/help/idea/connect-to-devcontainer.html#create_dev_container_inside_ide)
- [VSCode](https://code.visualstudio.com/docs/devcontainers/containers#_quick-start-open-an-existing-folder-in-a-container)

### 5) Run API and Frontend
- [API README](./api/README.md)   
- [Frontend README](./frontend/README.md)
