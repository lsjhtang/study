cd pdfiles && protoc --go_out=plugins=grpc:../services Prod.proto
cd ../..