# tidy repository
tidy:
	@go mod tidy

# reload .gitignore cache
reload:
	@git rm -rf --cached .
