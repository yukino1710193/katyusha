#!/bin/bash

RED="\e[31m"
GREEN="\e[32m"
BLUE="\e[34m"
YELLOW="\e[33m"
NC="\e[0m"

REDBGR='\033[0;41m'
NCBGR='\033[0m'

########## CONFIG ##########
DOCKER_REGISTRY="bonavadeur"
IMAGE="katyusha" # docker.io/{DOCKER_REGISTRY}/{IMAGE}
NAMESPACE="knative-serving"
component="activator"
OPTION=$1
############################

logSuccess() { echo -e "$GREEN-----$1-----$NC";}
logInfo() { echo -e "$BLUE-----$1-----$NC";}
logError() { echo -e "$RED-----$1-----$NC";}
logStage() { echo -e "$YELLOW###############---$1---###############$NC";}

koBuild() {
    logStage "Ko build image"
    export KO_DOCKER_REPO="ko.local"
    ko build ./cmd/$component/
    if [ "$?" -ne "0" ]; then
        logError "ko build error"
        exit 1
    else
        logSuccess "ko build successfully"
    fi
    endTime=`date +%s`
    koBuildTime=`expr $endTime - $startTime`
    logInfo "KoBuild time was $koBuildTime seconds."
}

dockerBuild() {
    compress=$1
    mode=$2
    
    logStage "Go build binary"
    CGO_ENABLED=0 go build -ldflags="-s -w" -o ./$IMAGE ./cmd/$IMAGE

    logStage "Compress binary"
    if [ $compress == "fast" ]; then
        upxx -1 $IMAGE
    elif [ $compress == "high" ]; then
        upxx -9 $IMAGE
    fi

    logStage "Docker build image"
    docker rmi $DOCKER_REGISTRY/$IMAGE:dev
    docker build --no-cache -t $DOCKER_REGISTRY/$IMAGE:dev .
    docker save -o $IMAGE.tar $DOCKER_REGISTRY/$IMAGE:dev

    sudo crictl rmi docker.io/$DOCKER_REGISTRY/$IMAGE:dev
    sudo ctr -n=k8s.io images import $IMAGE.tar
    if [ "$mode" == "mnode" ]; then
        ssh node2 "crictl rmi docker.io/$DOCKER_REGISTRY/$IMAGE:dev"
        ssh node2 "ctr -n=k8s.io images import $IMAGE.tar"
        ssh node3 "crictl rmi docker.io/$DOCKER_REGISTRY/$IMAGE:dev"
        ssh node3 "ctr -n=k8s.io images import $IMAGE.tar"
    fi

    logStage "Clean up"
    rm -rf ./$IMAGE
    rm -rf $IMAGE.tar
    if [ "$mode" == "mnode" ]; then
        ssh node2 "rm -rf $IMAGE.tar"
        ssh node3 "rm -rf $IMAGE.tar"
    fi
}

convertImage() {
    logStage "change image from docker to crictl"
    image=$(docker images | grep ko.local | grep $component | grep latest | awk '{print $1}'):latest
    docker rmi -f $DOCKER_REGISTRY/$IMAGE:dev
    docker image tag $image $DOCKER_REGISTRY/$IMAGE:dev
    docker rmi $image
    image=$(docker images | grep ko.local | grep $component | awk '{print $1}'):$(docker images | grep ko.local | grep $component | awk '{print $2}')
    docker rmi $image
    docker save -o $IMAGE.tar $DOCKER_REGISTRY/$IMAGE:dev
    logSuccess "Saved atarashi-imeji to .tar file"
    
    sudo crictl rmi docker.io/$DOCKER_REGISTRY/$IMAGE:dev
    sudo ctr -n=k8s.io images import $IMAGE.tar
    if [ "$plane" == "data" ]; then
        ssh node2 "crictl rmi docker.io/bonavadeur/ikukantai-$component:dev"
        ssh node2 "ctr -n=k8s.io images import /root/ikukantai-$component.tar"
        ssh node3 "crictl rmi docker.io/bonavadeur/ikukantai-$component:dev"
        ssh node3 "ctr -n=k8s.io images import /root/ikukantai-$component.tar"
    fi
    logSuccess "Untar atarashi-imeji"

    rm -rf $IMAGE.tar
    if [ "$plane" == "data" ]; then
        ssh node2 "rm -rf /root/ikukantai-$component.tar"
        ssh node3 "rm -rf /root/ikukantai-$component.tar"
    fi
    message="Untar atarashi-imeji" && logSuccess
}

pushDockerImage() {
    logStage "pushing image to Docker Hub"
    tag=$1
    CONTAINER_REGISTRY="docker.io"/$DOCKER_REGISTRY
    docker tag $CONTAINER_REGISTRY/$IMAGE:dev $CONTAINER_REGISTRY/$IMAGE:$tag
    docker push $CONTAINER_REGISTRY/$IMAGE:$tag
    docker rmi $CONTAINER_REGISTRY/$IMAGE:$tag
}

deployNewVersion() {
    logStage "remove current Pod"
    pods=($(kubectl -n $NAMESPACE get pod | grep $component | awk '{print $1}'))
    for pod in ${pods[@]}
    do
        kubectl -n $NAMESPACE delete pod/$pod &
    done
}

logPod() {
    sleep 1
    pods=($(kubectl -n $NAMESPACE get pod | grep $component | grep Running | awk '{print $1}'))
    while [ "${pods[0]}" == "" ];
    do
        sleep 1
        pods=($(kubectl -n $NAMESPACE get pod | grep $component | grep Running | awk '{print $1}'))
    done
    echo "pod:"${pods[0]}
    kubectl -n $NAMESPACE wait --for=condition=ready pod ${pods[0]} > /dev/null 2>&1
    clear
    endTime=`date +%s`
    logInfo "KoBuild time was $koBuildTime seconds."
    logInfo "Build time was `expr $endTime - $startTime` seconds."
    logStage "$IMAGE logs"
    echo "pod:"${pods[0]}
    kubectl -n $NAMESPACE logs ${pods[0]} -c queue-proxy -f
}
#
#
#
#
#
#
#
#
#
#
clear
echo -e "$REDBGR このスクリプトはボナちゃんによって書かれています $NCBGR"

startTime=`date +%s`

if [ $OPTION == "ful" ]; then
    # ./build.sh ful snode
    # ./build.sh ful mnode
    dockerBuild fast $2
    deployNewVersion
    logPod
elif [ $OPTION == "push" ]; then
    # ./build.sh ful snode
    # ./build.sh ful mnode
    dockerBuild "high" $2
    pushDockerImage $2
    deployNewVersion
    sleep 1
elif [ $OPTION == "log" ]; then
    deployNewVersion
    logPod
# elif [ $OPTION == "ko" ]; then
#     image=$(docker images | grep ko.local | grep $IMAGE | awk '{print $3}')
#     docker rmi -f $image
#     koBuild
# elif [ $OPTION == "build" ]; then
#     dockerBuild
# elif [ $OPTION == "dep" ]; then
#     convertImage
#     deployNewVersion
#     logPod
# elif [ $OPTION == "debug" ]; then
#     koBuild
#     convertImage
fi
