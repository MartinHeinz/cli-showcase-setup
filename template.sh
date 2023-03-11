#!/bin/bash

# https://github.com/paxtonhare/demo-magic/blob/master/samples/demo-template.sh

source demo-magic.sh

DEMO_PROMPT="${GREEN}➜ ${CYAN}\W ${COLOR_RESET}"
# speed at which to simulate typing. bigger num = faster
TYPE_SPEED=30

function comment() {
  cmd=$DEMO_COMMENT_COLOR$1$COLOR_RESET
  echo -en "$cmd"; echo ""
}

clear

comment "# Simple commands:"
pe 'ps aux | head'
pe 'ls -l'

comment "# Print and execute immediately"
pei 'cat some-file'
echo

comment "# Long running:"
pe 'docker build -t some-app .'

comment "# Error prone (dependant on external factors like network):"

pe 'curl --silent -X GET https://httpstat.us/418 -H "Accept: application/json" | jq .'

pe 'curl -i -X POST "https://httpbin.org/status/204" --data "{'some':'data'}"'

pe 'wget -q --show-progress -O some-file.txt https://raw.githubusercontent.com/octocat/Hello-World/master/README'

comment "# Hard to type:"

pe 'openssl req -x509 -newkey rsa:4096 -sha256 -days 3650 -nodes -keyout example.key -out example.crt -subj "/CN=example.com" -addext "subjectAltName=DNS:example.com,DNS:www.example.net,IP:10.0.0.1" 2>/dev/null'

comment "#  Enter interactive mode..."
cmd  # Run 'ls -l | grep example' to show result of 'openssl ...'

pe "clear"
