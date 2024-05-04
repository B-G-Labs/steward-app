#!/bin/bash

# fnm
export PATH="/home/ubuntu/.local/share/fnm:$PATH"
eval "`fnm env`"

export PNPM_HOME="/home/ubuntu/.local/share/pnpm"
export PATH="$PNPM_HOME:$PATH"

# 1. Fetch the latest code from remote
cd ~/Projects/steward-app && git pull -f origin master

# 2. Install dependencies
pnpm install

# 3. Build step that compiles code, bundles assets, etc.
turbo build

# 4. Kill all screens
screen -ls | grep 'nuxt' | awk '{print $1}' | xargs -I % -t screen -X -S % quit
pkill -9 -f 'node .output/server/index.mjs'

# 5. Make new screen and serve page
screen -dmS nuxt bash -c 'cd ~/Projects/steward-app && turbo serve; exec bash'