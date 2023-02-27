VARIABLE_ID=$(printenv HERO_ID)
curl https://learn.reboot01.com/assets/superhero/all.json | jq '.[] | select( .id == '${VARIABLE_ID}' ).connections.relatives' | tr -d '"'