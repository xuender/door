# protoc \
# --plugin=protoc-gen-ts=/home/ender/.nvm/versions/node/v6.10.2/bin/protoc-gen-ts \
# --js_out=import_style=commonjs,binary:./example/src/generated \
# --ts_out=service=true:./example/src/generated \
# -I ./proto \
# proto/*.proto

protoc --go_out=. -I . ./*.proto
