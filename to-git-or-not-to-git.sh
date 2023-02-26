0NUMBER=$(curl -s https://learn.reboot01.com/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"malsamma\"}}){id}}"}' | tr -dc '0-9')

echo $NUMBER