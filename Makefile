generate-booking-pb:
	protoc -I=proto proto/*.proto --go_out=:pb  --go-grpc_out=:pb
kill-port:
	sudo kill -9 `sudo lsof -t -i:2223`