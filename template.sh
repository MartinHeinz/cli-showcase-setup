#!/bin/bash

# https://github.com/paxtonhare/demo-magic/blob/master/samples/demo-template.sh

########################
# include the magic
########################


source demo-magic.sh
# . demo-magic.sh  # for Mac

DEMO_PROMPT="${GREEN}➜ ${CYAN}\W ${COLOR_RESET}"
# speed at which to simulate typing. bigger num = faster
TYPE_SPEED=30

function comment() {
  cmd=$DEMO_COMMENT_COLOR$1$COLOR_RESET
  echo -en "$cmd"; echo ""
}

clear

#p '# Just a comment'
comment '# Just a comment'
pe echo 'Hello!'

# print and execute immediately: ls -l
pei "ls -l"

p '# Enter interactive mode...'
cmd

pe "clear"