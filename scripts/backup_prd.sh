#!/bin/bash
DATE=`date +"%Y-%m-%d"`
USER=`whoami`
docker exec yajirobe_mysql_prod_1 mysqldump yajirobe_db > "/home/$USER/app/backups/yajirobe/$DATE.sql"