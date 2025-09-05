include .env

codegen_api_spec := ./_specs/openapi.yaml
codegen_dir := ./internal/codegen
codegen_options := #--global-property models,apis
codegen_config := ./_specs/openapi_generator_config.yaml 

pg_dsn := "postgres://${PG_USER}:${PG_PASSWORD}@${PG_HOST_EXTERNAL}:${PG_PORT_EXTERNAL}/${PG_DB}?sslmode=${PG_SSLMODE}" 
migrations_dir := ./migrations

build_source_file := ./cmd/app.go
build_dir := ./build
build_exec_name := app

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


app_build_local:
	go build -o ${build_dir}/${build_exec_name} ${build_source_file}

app_run_local: app_build_local
	./${build_dir}/${build_exec_name}

migrations_install_goose:
	go install github.com/pressly/goose/v3/cmd/goose@latest

migrations_up:
	goose -dir ${migrations_dir} postgres ${pg_dsn} up -v

migrations_down:
	goose -dir ${migrations_dir} postgres ${pg_dsn} down -v

