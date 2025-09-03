codegen_api_spec := ./_specs/openapi.yaml
codegen_dir := ./internal/codegen
codegen_options := #--global-property models,apis
codegen_config := ./_specs/openapi_generator_config.yaml 

api_codegen_generate:
	if [ -d ${codegen_dir} ]; then \
		echo "Codegen dir ${codegen_dir} already exists."; \
	else \
		mkdir -p ${codegen_dir}; \
		echo "Created codegen dir ${codegen_dir}."; \
	fi
	openapi-generator-cli generate -g go-server -c ${codegen_config} -i ${codegen_api_spec} -o ${codegen_dir} ${codegen_options}

api_codegen_clean:
	rm -rf ${codegen_dir}

api_codegen_regenerate: api_codegen_clean api_codegen_generate
	echo "Regenerating api endpoins..."
