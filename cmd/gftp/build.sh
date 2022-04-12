export GOOS=windows
go build  -o gftp.exe .

export GOOS=linux
go build  -o gftp .
