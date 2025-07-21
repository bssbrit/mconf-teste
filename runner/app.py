import requests

x = requests.get('http://localhost:3000/')

print(x.text)