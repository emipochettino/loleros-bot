.PHONY: deploy
deploy:
	@docker image rm -f loleros-bot;\
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .;\
docker build -t loleros-bot .;\
rm main;\
docker rm -f loleros-bot;\
docker run -d --restart unless-stopped --name loleros-bot --env-file config.env loleros-bot;\

.PHONY: start
start:
	@docker run -d --restart unless-stopped --name loleros-bot --env-file config.env loleros-bot

.PHONY: stop
stop:
	@docker stop loleros-bot

.PHONY: logs
logs:
	@tail -f -n 2000 logs.txt

