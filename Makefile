build-image:
	docker build -t gabrielcamurcab/finance .

run-app:
	docker-compose -f .devops/app.yml up -d