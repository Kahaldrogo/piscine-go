


curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg HERO_ID $HERO_ID '.[0] | .connections.relatives' | sed 's/"//g'