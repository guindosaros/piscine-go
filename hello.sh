#! /bin/bash

name=$(curl https://api.github.com/users/guindosaros| grep '"login":' | sed -e 's/^.": "//g' -e 's/",.$//g')

echo "hello: $name"