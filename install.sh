#!/usr/bin/bash

install_repo="github.com/jecqiang/mygo"
build_repo="github.com/jecqiang/mygo/mygo"

install_path=$( cd `dirname $0`; pwd )/mygo

if [ -d ${install_path} ]; then
    install_path=${install_path}-$( date +%Y%m%d%H%M%S)
fi

if [ -z ${GOPATH} ]; then
    GOPATH=`go env GOPATH`
fi

build(){
    go get ${build_repo}
    cd $GOPATH/src/${build_repo}
    go run build.go
}

install(){
    mkdir ${install_path}
    cd ${install_path}
    mkdir bin log src
    cp $GOPATH/src/${build_repo}/mygo ./bin/
    cp $GOPATH/src/${install_repo}/mygo.example.ini ./etc/mygo.ini
    cp -r $GOPATH/src/${install_repo}/public ./public
}

build

install

echo "Installing binary: ${install_path}/bin"
echo "Installing config: ${install_path}/etc"
echo "Installing web public: ${install_path}/public"
echo "Install complete"