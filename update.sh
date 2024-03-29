#!/bin/bash

VERSION=v2.5.6
curl https://raw.githubusercontent.com/alephium/alephium/${VERSION}/api/src/main/resources/openapi.json -o openapi.json
openapi-generator generate -i ./openapi.json -g go -o ./ --skip-validate-spec --package-name alephium --git-repo-id go-sdk --git-user-id alephium
rm openapi.json
rm git_push.sh
