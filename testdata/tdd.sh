#! /bin/bash
./ua.sh
curl -L \
    -A "$UA" \
    "https://google.com/search?q=test+driven+development" \
    -o tdd.html
