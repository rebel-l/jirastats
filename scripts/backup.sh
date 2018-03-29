#!/usr/bin/env bash
TIMESTAMP=`date +"%Y%m%d_%H%M%S"`
BACKUP_ROOT_PATH=./backup
BACKUP_PATH=$BACKUP_ROOT_PATH/$TIMESTAMP
STORAGE_PATH=./storage
DELETE_BACKUP=0
DO_BACKUP=1
REMOVE_ORIGINAL=0

showHelp(){
    echo
    echo "Scripts resets project folder to initial state (cleanup datafiles etc)"
    echo "usage: ./scripts/backup.sh [options]"
    echo
    echo "Options:"
    echo "-h, -?    shows this help"
    echo "-d        deletes old backups"
    echo "-s        skips backup of files"
    echo "-r        removes original data file: BE CAREFUL, ALL DATA IS LOST IF YOU COMBINE IT WITH OPTION -s!"
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

while getopts "h?drs" opt; do
    case "$opt" in
        h|\?)
            showHelp
            exit 0
            ;;
        d)
            DELETE_BACKUP=1
            ;;
        r)
            REMOVE_ORIGINAL=1
            ;;
        s)
            DO_BACKUP=0
            ;;
    esac
done


echo
echo -en "\033[40;36m\033[1mBackup Data of Jira Stats project ...\033[0m"
echo
echo

if [ $DELETE_BACKUP -eq 1 ]
then
    deletBackups
fi

if [ $DO_BACKUP -eq 1 ]
then
    doBackup
fi

if [ $REMOVE_ORIGINAL -eq 1 ]
then
    resetProject
fi

echo
echo -en "\033[40;32m\033[1mJira Stats backup successful :-)\033[0m"
echo
echo
