#!/bin/bash
#:[.''.]:>-------------------------------------------
#:[.''.]:> Author...: Marco Antonio Rubio Lopez
#:[.''.]:> Contact..: markitos.es.info@gmail.com
#:[.''.]:> createdAt: diciembre 2024
#:[.''.]:> updatedAt: enero 2025
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><docker-compose>
docker-up:
	bash bin/docker-up.sh
docker-down:
	bash bin/docker-down.sh
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><runs>
run:
	bash bin/run.sh
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><testing>
test-cache-clear:
	bash bin/test-cache-clear.sh
test: test-cache-clear
	bash bin/test.sh
testv: test-cache-clear
	bash bin/testv.sh
testc: test-cache-clear
	bash bin/testc.sh
testcv: test-cache-clear
	bash bin/testcv.sh
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><database - migrations - Xorm>
createdb:
	bash bin/createdb.sh
dropdb: 
	bash bin/dropdb.sh
recreatedb: dropdb createdb test
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:><security>
appsec-sast:
	bash bin/appsec-sast.sh
appsec-secrets:
	bash bin/appsec-secrets.sh	
appsec: appsec-sast appsec-secrets
#:[.''.]:>-------------------------------------------



#:[.''.]:>-------------------------------------------
#:[.''.]:><building>
docker-login:
	bash bin/docker-login.sh $(GITHUB_TOKEN)
docker-publish-tag: docker-login
	bash bin/docker-publish-tag.sh $(or $(TAG),1.0.0)
docker-publish-postgres: docker-login
	bash bin/docker-publish-postgres.sh $(or $(TAG),1.0.0)
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:> License: OpenSource :) to pa ti!
#:[.''.]:> markitos.es
#:[.''.]:>-------------------------------------------
