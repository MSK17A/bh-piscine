import requests
import json

response_API = requests.post('https://learn.reboot01.com/api/graphql-engine/v1/graphql' , params='{"query":"{user(where:{login:{_eq:\"malsamma\"}}){id}}"}')

print(response_API.text)