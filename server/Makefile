start:
	touch data.json
	docker compose -f docker-compose.yml up -d
	./data-poller.sh & disown

stop:
	- docker compose -f docker-compose.yml down
	- rm data.json
