setup:
	git clone https://github.com/datacharmer/test_db

up:
	docker-compose down
	docker-compose up -d
	# For some reason, orchestrator container cannot start until main and read replica are fully prepared
	sleep 10
	docker-compose up -d
	docker-compose ps
	# Failed to register a main DB for some reasons. Wait
	sleep 10
	docker exec -it cluster_orchestrator_1 orchestrator-client -c discover -i main:3306
