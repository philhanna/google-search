#! /bin/bash

# Note: only the first line of the header seems to be necessary ("Mozilla/5.0")
export UA=$(cat << EOF
Mozilla/5.0
(Windows NT 10.0; Win64; x64)
AppleWebKit/537.36
(KHTML, like Gecko)
Chrome/91.0.4472.124
Safari/537.36
EOF
)
