# 在 rpc 目录下执行（进入 apps/user/rpc 目录）
# goctl rpc protoc user.proto --go_out=. --go-grpc_out=. --zrpc_out=.

# 在 im_demo 根目录下执行
goctl rpc protoc apps/user/rpc/user.proto --go_out=apps/user/rpc --go-grpc_out=apps/user/rpc --zrpc_out=apps/user/rpc

goctl model mysql ddl -src="./deploy/sql/user.sql" -dir="./apps/user/models" -c

goctl api go -api apps/user/api/user.api -dir apps/user/api -style gozero