include Makefile.variables

.PHONY: format todo test check prepare

local: | quiet
	@$(eval DOCKRUN= )
	@mkdir -p tmp
	@touch tmp/dev_image_id
quiet:
	@:

prepare: tmp/dev_image_id
tmp/dev_image_id:
	@mkdir -p tmp
	@docker rmi -f ${DEV_IMAGE} > /dev/null 2>&1 || true
	@docker build -t ${DEV_IMAGE} -f Dockerfile.dev .
	@docker inspect -f "{{ .ID }}" ${DEV_IMAGE} > tmp/dev_image_id

test: check db_prepare
	${DOCKTEST} bash ./scripts/test.sh

check:prepare format
	${DOCKRUN} bash ./scripts/check.sh

format:
	${DOCKRUN} bash ./scripts/format.sh

db_start: db_stop
	@docker run --name student-portal-mysql-db -e MYSQL_ALLOW_EMPTY_PASSWORD=yes -p 3306:3306 -d mysql:5.6
	@docker run --name student-portal-mongo-db -p 27015-27017:27015-27017 -d mongo:4.2.0

db_prepare: db_start
	@docker cp student_portal.sql student-portal-mysql-db:student_portal.sql
	@echo "Executing databases...wait for 15 seconds"
	@sleep 15
	@docker exec -i student-portal-mysql-db sh -c 'mysql -uroot < student_portal.sql'

db_stop:
	bash ./scripts/db_stop.sh

codegen: prepare
	${DOCKRUN} bash ./scripts/swagger.sh

