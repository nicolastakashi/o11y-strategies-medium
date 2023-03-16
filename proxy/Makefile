test:
	echo "GET http://localhost:8080/checkouts" | vegeta attack -duration=60s | vegeta report

setup:
	@docker-compose up

destroy:
	@docker-compose down -v