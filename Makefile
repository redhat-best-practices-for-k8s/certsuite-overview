# Define a variable with a default value for the Grafana container name
GRAFANA_CONTAINER_NAME?=grafana

vet:
	go vet ./...

build:
	go build ./...

lint:
	golangci-lint run ./...

# The 'stop-running-grafana-container' command stops and removes any running Grafana container
stop-running-grafana-container:
	docker ps -q --filter "name=${GRAFANA_CONTAINER_NAME}" | xargs -r docker stop
	docker ps -aq --filter "name=${GRAFANA_CONTAINER_NAME}" | xargs -r docker rm

# The 'clone-tnf-secrets' command clones the 'tnf-secrets' repository from GitHub and updates the local secrets folder
clone-tnf-secrets:
	rm -rf tnf-secrets
	git clone https://github.com/redhat-best-practices-for-k8s/tnf-secrets.git

clone certsuite-overview:
	rm -rf certsuite-overview
	git clone https://github.com/redhat-best-practices-for-k8s/certsuite-overview.git

# The 'run-grafana' command starts the Grafana container, modifies the datasource.yaml file, and mounts the necessary volumes
run-grafana: clone-certsuite-overview clone-tnf-secrets stop-running-grafana-container
	sed \
		-e 's/MysqlUsername/$(shell jq -r ".MysqlHost" "./tnf-secrets/certsuite-overview-secrets.json" | base64 -d)/g' \
		-e 's/DB_PASSWORD/$(shell jq -r ".DB_PASSWORD" "./tnf-secrets/certsuite-overview-secrets.json" | base64 -d)/g' \
		./grafana/datasource/datasource.yaml > datasource-certsuite-overview.yaml
	
# Run the Grafana container with the necessary environment variables and mounted volumes
	docker run -d -p 3000:3000 --name=${GRAFANA_CONTAINER_NAME} \
		-e "GF_SECURITY_ADMIN_USER=$(shell jq -r ".GrafanaUsername" "./tnf-secrets/certsuite-overview-secrets.json")" \
		-e "GF_SECURITY_ADMIN_PASSWORD=$(shell jq -r ".GrafanaPassword" "./tnf-secrets/certsuite-overview-secrets.json")" \
		-v ./grafana/dashboard:/etc/grafana/provisioning/dashboards \
		-v ./datasource.yaml:/etc/grafana/provisioning/datasources/datasource-config-prod.yaml \
		-v ./datasource-certsuite-overview.yaml:/etc/grafana/provisioning/datasources/datasource-certsuite-overview.yaml
		grafana/grafana
	
# Clean up by removing the generated datasource.yaml file and the 'tnf-secrets' folder
	rm datasource-config-prod.yaml
	rm -rf tnf-secrets
