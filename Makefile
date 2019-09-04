# Kolej Project Makefile

all: run

run: 
	docker-compose up -d

build:
	docker-compose build

stop:
	docker-compose down