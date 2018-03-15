#!/usr/bin/env bash
TIMESTAMP=`date +"%Y%m%d_%H%M%S"`
BACKUP_ROOT_PATH=./backup
BACKUP_PATH=$BACKUP_ROOT_PATH/$TIMESTAMP
STORAGE_PATH=./storage
DELETE_BACKUP=0
SKIP_BACKUP=0

showHelp(){
    echo
    echo "Scripts resets project folder to initial state (cleanup datafiles etc)"
    echo "usage: ./scripts/resetProject.sh [options]"
    echo
    echo "Options:"
    echo "-h, -?    shows this help"
    echo "-d        deletes old backups"
    echo "-s        skips backup of files"
    echo
}

doBackup(){
    echo -en "\033[40;33m\033[1mBackup files to $BACKUP_PATH\033[0m"
    echo
    mkdir -p $BACKUP_PATH
    cp $STORAGE_PATH/*.db $BACKUP_PATH
}

deletBackups(){
    echo -en "\033[40;35m\033[1mRemove old backups\033[0m"
    echo
    rm -r $BACKUP_ROOT_PATH
}

resetProject(){
    rm $STORAGE_PATH/*.db
}

while getopts "h?ds" opt; do
    case "$opt" in
        h|\?)
            showHelp
            exit 0
            ;;
        d)
            DELETE_BACKUP=1
            ;;
        s)
            SKIP_BACKUP=1
            ;;
    esac
done


echo
echo -en "\033[40;36m\033[1mReset Jira Stats project ...\033[0m"
echo
echo

if [ $DELETE_BACKUP -eq 1 ]
then
    deletBackups
fi

if [ $SKIP_BACKUP -eq 0 ]
then
    doBackup
fi

resetProject

echo
echo -en "\033[40;32m\033[1mJira Stats reset successful :-)\033[0m"
echo
echo
