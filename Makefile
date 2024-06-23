update_docs:
	docker exec zord-http swag init -g ./cmd/http/main.go --parseDependency --propertyStrategy pascalcase
