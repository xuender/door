protoc \
--plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts \
--js_out=import_style=commonjs,binary:. \
--ts_out=service=true:. \
-I . \
*.proto

protoc --go_out=. -I . ./*.proto
