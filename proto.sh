protoc \
--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:src \
--ts_out=service=true:src \
-I . \
*.proto

protoc --go_out=. -I . ./*.proto
