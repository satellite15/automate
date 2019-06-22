# we use pushd/popd here, and /bin/sh of our chefes/buildkite image is not bash
# so we have to override the default shell here
SHELL=bash
SWAGGER_SOURCE_DIRS=../automate-gateway/api ../../api/external
SWAGGER_RESULT_FILE=static/api-docs/all-apis.swagger.json

themes/chef:
	git clone https://${GITHUB_TOKEN}@github.com/chef/chef-hugo-theme.git themes/chef

clean:
	rm -rf site/themes/chef
	rm -f $(SWAGGER_RESULT_FILE)

sync: themes/chef
	pushd themes/chef && git fetch && git reset --hard origin/master && popd

SWAGGER_FILES = $(shell find $(SWAGGER_SOURCE_DIRS) -name '*.swagger.json')

$(SWAGGER_RESULT_FILE): $(SWAGGER_FILES)

EXPECTED_CONFLICTS=47

generate_swagger: $(SWAGGER_RESULT_FILE)
	swagger mixin -c=$(EXPECTED_CONFLICTS) $(SWAGGER_FILES) > $(SWAGGER_RESULT_FILE) 2> /dev/null

$(SWAGGER_RESULT_FILE).ruby: $(SWAGGER_FILES)

generate_swagger_ruby: $(SWAGGER_RESULT_FILE).ruby
	DEBUG=true ruby scripts/generate_all_swagger.rb $(SWAGGER_FILES) > $(SWAGGER_RESULT_FILE).ruby

serve: sync generate_swagger
	hugo server --buildDrafts --noHTTPCache

lint: themes/chef
	hugo -D
