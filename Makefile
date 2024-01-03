build-base:
	docker build -f ./docker/app/base/Dockerfile -t showwin/ishocon2_app_base:latest .;

build-bench:
	docker build -f ./docker/benchmarker/Dockerfile -t showwin/ishocon2_bench:latest .;

build:
	docker compose -f ./docker-compose.yml build;

build-app:
	docker build -f "./docker/app/$(LANG)/Dockerfile" -t ishocon2-app:latest .;

up:
	docker compose up -d;

up-nod:
	docker compose up;

down: 
	docker compose down;

bench:
	docker exec -i ishocon2-bench-1 sh -c "./benchmark --ip app:443 --workload 8"

bench-with-db-init: up
	docker exec -i ishocon2-bench-1 sh -c " \
		service mysql restart \
		&& tar -jxvf /root/admin/ishocon2.dump.tar.bz2 && mysql -u root -pishocon ishocon2 < /root/admin/ishocon2.dump \
		&& ./benchmark --ip app:443 --workload 8 \
	";

change-lang:
	if sed --version 2>&1 | grep -q GNU; then \
		echo "GNU sed"; \
		sed -i 's/\(ruby\|python\|go\|php\|nodejs\|crystal\)/'"$(LANG)"'/g' ./docker-compose.yml; \
	else \
		echo "BSD sed"; \
		sed -i '' -E 's/(ruby|python|go|php|nodejs|crystal)/'"$(LANG)"'/g' ./docker-compose.yml; \
	fi;
