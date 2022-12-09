mock:
	mockgen -source=worker/init.go -destination=worker/worker.mock/mock.go -package=mockWorker
	mockgen -source=util/init.go -destination=util/util.mock/mock.go -package=mockUtil

# tidy repository
tidy:
	@go mod tidy

# reload .gitignore cache
reload:
	@git rm -rf --cached .
