#!/usr/bin/env bash
CURRENT_PATH=`pwd`

install(){
    TOOL_PATH=${CURRENT_PATH}/tools/jirastats-$1

    echo "Build $1 ... path: $TOOL_PATH"
    cd ${TOOL_PATH}
    go install
}

echo
echo -en "\033[40;36m\033[1mBuild Jira Stats Server ...\033[0m"
echo

echo
echo "Actual path: $CURRENT_PATH"
echo

install collector
install server
install setup

cd ${CURRENT_PATH}
echo
echo -en "\033[40;32m\033[1mJira Stats Server build successful :-)\033[0m"
echo
echo
