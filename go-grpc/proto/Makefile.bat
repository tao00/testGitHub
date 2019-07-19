protoc ^
			-I . ^
			-I %GOPATH%/src ^
			-I %GOPATH%/src/github.com/tao00/go-grpc/vendor/github.com/grpc-ecosystem/grpc-gateway ^
			-I %GOPATH%/src/github.com/tao00/go-grpc/vendor/github.com/gogo/googleapis ^
			-I %GOPATH%/src/github.com/tao00/go-grpc/vendor ^
			-I %GOPATH%/src/github.com/google/protobuf/src ^
		--gogo_out=plugins=grpc,^
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,^
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,^
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:^
. ^
		--grpc-gateway_out=^
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,^
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,^
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:^
. ^
		--govalidators_out=gogoimport=true,^
Mgoogle/protobuf/timestamp.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types,^
Mgoogle/protobuf/empty.proto=github.com/gogo/protobuf/types,^
Mgoogle/api/annotations.proto=github.com/gogo/googleapis/google/api,^
Mgoogle/protobuf/field_mask.proto=github.com/gogo/protobuf/types:^
. ^
				proto.proto
		REM del protopb_test.go
		
		REM cd ../../../
		REM rpc_helper.exe rpcitr -cli 1