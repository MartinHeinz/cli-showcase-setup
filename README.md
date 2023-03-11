# Template for CLI-based Demos/Showcases

Templates and config for running scripted and/or pre-recorded CLI demos.

-----

If you find this useful, you can support me on Ko-Fi (Donations are always appreciated, but never required):

[![ko-fi](https://ko-fi.com/img/githubbutton_sm.svg)](https://ko-fi.com/K3K6F4XN6)

## Scripted Demo

Repository provides template for 2 scripting tools, below is setup needed for each of them.

Setup (`demo-magic`):

```shell
# Install "Pipe Viewer" for simulated typing
# MacOS:
brew install pv
# Ubuntu:
apt-get install pv

git clone https://github.com/paxtonhare/demo-magic.git
cp demo-magic/demo-magic.sh demo-magic.sh
```

To create and run a new demo:

```shell
cp template.sh my-cool-demo.sh
# Edit 'my-cool-demo.sh'

# Run
./my-cool-demo.sh
```

---

Setup (`https://github.com/saschagrunert/demo`):

```shell
go get github.com/saschagrunert/demo@latest

# Edit 'main.go'

go build
./cli-demo --all  # or change 'all' to name of specific demo run
```

## Recording Demo

```shell
# You might need 'sudo'
npm install -g terminalizer

# https://github.com/faressoft/terminalizer/issues/29
# Very careful with below command!
sudo chown -R <YOUR_USERNAME> /usr/lib/node_modules/terminalizer/render/

terminalizer init
# The global config directory is created at
# /home/martin/.terminalizer

# Adjust config.yaml... (e.g., change 'title' field)

# Create copy of default config in current directory
terminalizer config

# Start recording:
terminalizer record demo --config config.yaml

# ... run some commands

# CTRL+D to stop recording...

# Successfully Recorded
# The recording data is saved into the file:
# /home/.../demo.yml

terminalizer play demo  # To check
terminalizer render demo  # To render
```

## Workflow

Example usage/streamlined workflow:

```shell
# Create demo:

# ... Edit 'main.go'
go build
# OR
cp template.sh my-cool-demo.sh
# Edit 'my-cool-demo.sh'

# -----

# Start recording:
terminalizer record demo

# Start scripted demo:
./cli-demo --all -i
# OR
./my-cool-demo.sh

# ... Go through the demo
# CTRL+D

terminalizer play demo  # To check
# ... Adjust demo.yml
terminalizer render demo  # To render
```

