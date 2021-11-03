up:
	docker-compose down
	docker-compose up -d --scale read-replica=2
	# For some reason, orchestrator container cannot start until main and read replica are fully prepared
	sleep 5
	docker-compose up -d
