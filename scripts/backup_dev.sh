#!/bin/bash
DATE=`date +"%Y-%m-%d"`
docker exec -it yajirobe_mysql_dev_1 mysqldump yajirobe_db > "./backend/db/backup/$DATE.sql"