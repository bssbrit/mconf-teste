import requests
import sys

# x = requests.get('http://localhost:3000/')

book_name = " ".join(sys.argv[1:])
print(book_name)

params = {'book_name': book_name}
x = requests.get('http://localhost:3000/', params=params)

print(x.text)
print(x.status_code)