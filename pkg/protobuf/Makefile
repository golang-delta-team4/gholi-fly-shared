gen-user-proto:
	protoc -I=./user -I=. --go_out=./user --go_opt=paths=source_relative --go-grpc_out=./user --go-grpc_opt=paths=source_relative ./user/*.proto
gen-bank-proto:
	protoc -I=./bank -I=. --go_out=./bank --go_opt=paths=source_relative --go-grpc_out=./bank --go-grpc_opt=paths=source_relative ./bank/*.proto
gen-role-proto:
	protoc -I=./role -I=. --go_out=./role --go_opt=paths=source_relative --go-grpc_out=./role --go-grpc_opt=paths=source_relative ./role/*.proto
gen-make-all: gen-user-proto gen-bank-proto gen-role-proto
