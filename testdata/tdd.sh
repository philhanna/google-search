#! /bin/bash
./ua.sh
curl \
-A "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 \
(KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36" \
"https://www.google.com/search?q=test+driven+development" \
-o tdd.html
