# create mock files for packages
generate-mock:
	mockgen -source=worker/init.go -destination=worker/mock/worker_mock.go -package=mock
	mockgen -source=util/json/init.go -destination=util/json/mock/json_mock.go -package=mock
	mockgen -source=util/yaml/init.go -destination=util/yaml/mock/yaml_mock.go -package=mock
	mockgen -source=util/ptrconv/init.go -destination=util/ptrconv/mock/ptrconv_mock.go -package=mock

# remove mock files from packages
remove-mock:
	@rm -rf worker/mock
	@rm -rf util/json/mock
	@rm -rf util/yaml/mock
	@rm -rf util/ptrconv/mock
