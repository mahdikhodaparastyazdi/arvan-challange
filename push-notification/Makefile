COMPOSE_FILES=docker-compose.yml
COMPOSE_PROFILES=
COMPOSE_COMMAND=docker-compose

ifeq (, $(shell which $(COMPOSE_COMMAND)))
	COMPOSE_COMMAND=docker compose
	ifeq (, $(shell which $(COMPOSE_COMMAND)))
		$(error "No docker compose in path, consider installing docker on your machine.")
	endif
endif

env:
	@[ -e ./.env ] || cp -v ./.env.example ./.env

up:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) up -d
down:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) down --remove-orphans
mysql-shell:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec -u 0 mysql mysql -hmysql -u$(MYSQL_USER) -D$(MYSQL_DATABASE) -p$(MYSQL_PASSWORD)

redis-shell:
	$(COMPOSE_COMMAND) -f $(COMPOSE_FILES) exec -u 0 redis redis-cli