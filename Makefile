build:
	go build -gcflags=all="-N -l" -o program

dlv:
	dlv debug --headless --listen=:2346 --log main.go
