#!/usr/bin/env bash

CURRENT_PATH=`pwd`
DIST_PATH=${CURRENT_PATH}/dist
DIST_TMP_PATH=${DIST_PATH}/tmp
PACKAGE_NAME="jirastats"

validateParams(){
	if [ -z "$GOPATH" ];
	then
		echo
		echo -en "\033[40;31m\033[1mGOPATH environment variable is not set. Exit with failure!\033[0m"
		echo
		echo "For more information please have a look how to setup a go workspace: https://golang.org/doc/code.html"
		exit 1
	fi

	if [ -z "$OS" ];
	then
		echo
		echo -en "\033[40;31m\033[1mParameter for operating system must be provided. Exit with failure!\033[0m"
		echo
		exit 1
	fi

	if [ -z "$VERSION" ];
	then
		echo
		echo -en "\033[40;31m\033[1mVersion parameter must be provided. Exit with failure!\033[0m"
		echo
		exit 1
	fi
}

showHelp(){
    echo
    echo "Script creates the distribution and compress it in a zip file to upload."
    echo "usage: ./scripts/distribute.sh [options]"
    echo
    echo "Options:"
    echo "-h, -?                shows this help"
    echo "-o [operating system] indicates the operating for which the application is compiled for. The following options are possible: win (for windows), lin (for linux) or mac (for Mac OS)"
    echo "-v [version number]   the version number of the application you distribute"
    echo
}

listPath(){
	echo
	echo "Actual path: $CURRENT_PATH"
	echo "Distribution path: $DIST_PATH"
	echo "Go path: $GOPATH"
	echo
}

cleanupDistFolder(){
	# cleanup dist folder
	rm -r ${DIST_PATH}/*
	mkdir ${DIST_TMP_PATH}
}

buildServer(){
	# build server application
	${CURRENT_PATH}/scripts/buildServer.sh
	cp ${GOPATH}/bin/jirastats* ${DIST_TMP_PATH}
}

buildClient(){
	# build client application
	npm run build
	cp -r ${CURRENT_PATH}/public ${DIST_TMP_PATH}
	rm ${DIST_TMP_PATH}/public/.gitignore
}

prepare(){
    rm -r node_modules
    rm -r vendor
    npm install --no-bin-links
    glide install
}

packageWindows(){
	cd ${DIST_TMP_PATH}
	zip -r ./../${PACKAGE_NAME}-v${VERSION}-${OS}-x64.zip *
}

packageLinux(){
	cd ${DIST_TMP_PATH}
	tar cfz ./../${PACKAGE_NAME}-v${VERSION}-${OS}-x64.tar.gz --exclude=.gitignore .
}

packageMac(){
	cd ${DIST_TMP_PATH}
	tar cfz ./../${PACKAGE_NAME}-v${VERSION}-${OS}-x64.tar.gz --exclude=.gitignore .
}

while getopts "h?o:v:" opt; do
    case "$opt" in
        h|\?)
            showHelp
            exit 0
            ;;
        o)
			case "$OPTARG" in
			    win)
			        OS="windows"
			        ;;
			    lin)
			        OS="linux"
			        ;;
			    mac)
			        OS="macOS"
			        ;;
			esac
            ;;
        v)
			VERSION=$OPTARG
            ;;
    esac
done

echo
echo -en "\033[40;36m\033[1mBuild Jira Stats distribution ...\033[0m"
echo

validateParams
listPath
cleanupDistFolder

# create license
cp ${CURRENT_PATH}/LICENSE ${DIST_TMP_PATH}

prepare
buildServer
buildClient

# create storage path
mkdir ${DIST_TMP_PATH}/storage

# zip it
case "$OS" in
    windows)
        packageWindows
        ;;
    linux)
        packageLinux
        ;;
    macOS)
        packageMac
        ;;
esac

cd ${CURRENT_PATH}
rm -r ${DIST_TMP_PATH}

echo
echo -en "\033[40;32m\033[1mJira Stats distribution build successful :-)\033[0m"
echo
echo
