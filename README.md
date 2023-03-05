## Configuration and Template for CLI-based Demos/Showcases

### Scripted Demo

Setup (`demo-magic`):

```shell
# MacOS:
brew install pv

# Ubuntu:
apt-get install pv

git clone https://github.com/paxtonhare/demo-magic.git
cp demo-magic/demo-magic.sh demo-magic.sh
```

To create new demo:

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
./cli-demo --demo-0  # or change 'demo-0' to actual name
```

### Record Demo

```shell
# You might need 'sudo'
npm install -g terminalizer

# https://github.com/faressoft/terminalizer/issues/29
# Very careful with below command!
sudo chown -R <YOUR_USERNAME> /usr/lib/node_modules/terminalizer/render/

terminalizer init
# The global config directory is created at
# /home/martin/.terminalizer

# Adjust config.yaml... (change 'title' field)

# Make it the default:
cp config.yaml /home/martin/.terminalizer/config.yaml

# Start recording:
terminalizer record demo

# ... run some commands

# CTRL+D

# Successfully Recorded
# The recording data is saved into the file:
# /home/.../demo.yml

terminalizer play demo  # To check
terminalizer render demo  # To render
```

### Workflow

```shell
# Create demo:

# -----
# ... Edit 'main.go'

go build

# OR

cp template.sh my-cool-demo.sh
# Edit 'my-cool-demo.sh'

# -----

# Start recording:
terminalizer record demo

# Start scripted demo:
./cli-demo --some-demo
# OR
./my-cool-demo.sh

# ... Go through the demo
# CTRL+D

terminalizer play demo  # To check
# ... Adjust demo.yml
terminalizer render demo  # To render
```

