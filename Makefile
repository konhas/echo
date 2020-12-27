default: build

build: clean
		docker-compose build
		docker-compose -f docker-compose.test.yml build

run:
		docker-compose up --build app

clean:
		go mod tidy

test:
		docker-compose -f docker-compose.test.yml run test