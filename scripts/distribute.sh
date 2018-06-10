#!/usr/bin/env bash
CURRENT_PATH=`pwd`
DIST_PATH=${CURRENT_PATH}/dist
DIST_TMP_PATH=${DIST_PATH}/tmp
PACKAGE_NAME="jirastats"
WINDOWS_BUILD=0

validateParams(){
	if [ -z "$GOPATH" ];
	then
		echo
		echo -en "\033[40;31m\033[1mGOPATH environment variable is not set. Exit with failure!\033[0m"
		echo
		echo "For more information please have a look how to setup a go workspace: https://golang.org/doc/code.html"
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
    echo "-w                    creates a Windows distribution. Linux is default, so theres no need for a parameter."
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

packageWindows(){
	cd ${DIST_TMP_PATH}
	zip -r ./../${PACKAGE_NAME}-v${VERSION}-win-x86.zip *
}

packageLinux(){
	cd ${DIST_TMP_PATH}
	tar cfz ./../${PACKAGE_NAME}-v${VERSION}-linux-x86.tar.gz --exclude=.gitignore .
}

while getopts "h?wv:" opt; do
    case "$opt" in
        h|\?)
            showHelp
            exit 0
            ;;
        w)
			WINDOWS_BUILD=1
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

buildServer
buildClient

# create storage path
mkdir ${DIST_TMP_PATH}/storage

# zip it
if [ ${WINDOWS_BUILD} -eq 1 ];
then
	packageWindows
else
	packageLinux
fi

cd ${CURRENT_PATH}
rm -r ${DIST_TMP_PATH}

echo
echo -en "\033[40;32m\033[1mJira Stats distribution build successful :-)\033[0m"
echo
echo
