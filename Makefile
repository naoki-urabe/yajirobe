build-prod-api-image:
	docker build -f docker/api/Dockerfile.api.prod -t api-prod-image:latest .
set-mysql-dev-login-info:
	docker exec -it yajirobe_mysql_dev_1 mysql_config_editor set -u root -p

set-mysql-dev-dump-info:
	docker exec -it yajirobe_mysql_dev_1 mysql_config_editor set --login-path=mysqldump -u root -p

mysql-backup-dev:
	bash scripts/backup_dev.sh

mysql-restore-dev:
	docker exec -i yajirobe_mysql_dev_1 mysql yajirobe_db < ./backend/db/backup/backup.sql

set-mysql-prod-login-info:
	docker exec -it yajirobe_mysql_prod_1 mysql_config_editor set -u root -p

set-mysql-prod-dump-info:
	docker exec -it yajirobe_mysql_prod_1 mysql_config_editor set --login-path=mysqldump -u root -p

mysql-backup-prod:
	bash scripts/backup_prd.sh

mysql-restore-prod:
	docker exec -i yajirobe_mysql_prod_1 mysql yajirobe_db < ./backend/db/backup/backup.sql

migrate-up:
	migrate -database 'mysql://root:$(DB_USER_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/yajirobe_db' -path backend/db/migrations up

migrate-down:
	migrate -database 'mysql://root:$(DB_USER_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/yajirobe_db' -path backend/db/migrations down
