# create mock files for packages
generate-mock:
	mockgen -source=util/json/init.go -destination=util/json/mock/json_mock.go -package=mockUtilJSON
	mockgen -source=util/yaml/init.go -destination=util/yaml/mock/yaml_mock.go -package=mockUtilYAML
	mockgen -source=worker/init.go -destination=worker/mock/worker_mock.go -package=mockWorker

# remove mock files from packages
remove-mock:
	@rm -rf util/json/mock
	@rm -rf util/yaml/mock
	@rm -rf worker/mock
