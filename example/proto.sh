protoc \
--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:src/app \
--ts_out=service=true:src/app \
-I . \
*.proto

protoc \
--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:src/app \
--ts_out=service=true:src/app \
-I .. \
../*.proto


protoc --go_out=./chat -I . ./*.proto
