build-image:
	docker build -t gabrielcamurcab/finance -f Dockerfile.app .
	docker build -t bitnami/prometheus -f Dockerfile.prometheus .


run-app:
	docker-compose -f .devops/app.yml up -d
	docker-compose -f .devops/prometheus.yml up -d

run-test:
	docker-compose -f .devops/app.yml run app go test ./...