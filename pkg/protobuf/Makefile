gen-user-proto:
	protoc -I=./user -I=. --go_out=./user --go_opt=paths=source_relative --go-grpc_out=./user --go-grpc_opt=paths=source_relative ./user/*.proto
gen-bank-proto:
	protoc -I=./bank -I=. --go_out=./bank --go_opt=paths=source_relative --go-grpc_out=./bank --go-grpc_opt=paths=source_relative ./bank/*.proto
gen-factor-proto:
	protoc -I=./factor -I=. --go_out=./factor --go_opt=paths=source_relative --go-grpc_out=./factor --go-grpc_opt=paths=source_relative ./factor/*.proto
gen-wallet-proto:
	protoc -I=./wallet -I=. --go_out=./wallet --go_opt=paths=source_relative --go-grpc_out=./wallet --go-grpc_opt=paths=source_relative ./wallet/*.proto
gen-make-all: gen-user-proto gen-bank-proto gen-factor-proto gen-wallet-proto
