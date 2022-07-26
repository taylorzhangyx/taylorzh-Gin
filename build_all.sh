#!/bin/bash
funCheck(){
    if [ $? -ne 0 ]; then
        echo -e "\033[31m FAILED! \033[0m"
        exit 127
    else
        echo -e "\033[32m OK! \033[0m"
    fi
}


echo "==========  Make Build  =========="
GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

funCheck

echo "========== Docker Build =========="
IMAGE="test-app"
TAG="v1.0.0"
REGISTRY="ccr.ccs.tencentyun.com/tiemsdev"

docker build -t $REGISTRY/$IMAGE:$TAG -f Dockerfile .

echo "====== Docker Push (ccr.ccs) ======"
docker push $REGISTRY/$IMAGE:$TAG

