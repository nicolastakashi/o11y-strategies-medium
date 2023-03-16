test:
	echo "GET http://localhost:9091/checkouts" | vegeta attack -duration=300s | vegeta report

setup:
	@docker-compose up --build

destroy:
	@docker-compose down -v