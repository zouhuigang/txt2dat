windows编译64位：

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  main.go

windows编译32位：
	
	CGO_ENABLED=0 GOOS=windows GOARCH=386  go build  main.go

mac版：

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

linux版：

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go


将output.txt转换为phone.dat格式数据