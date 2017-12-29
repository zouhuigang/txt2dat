windows编译64位：

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  main.go

windows编译32位：
	
	CGO_ENABLED=0 GOOS=windows GOARCH=386  go build  main.go

mac版：

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

linux版：

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go

1701：
--
> output_1701.txt->phone_1701.dat


	归属地信息库文件大小：3,203,029 字节
	归属地信息库最后更新：2017年1月
	手机号段记录条数：354522

1704:

 	归属地信息库文件大小：3,257,452 字节
	归属地信息库最后更新：2017年4月
	手机号段记录条数：360569

1707:

	归属地信息库文件大小：3,293,083 字节
	归属地信息库最后更新：2017年7月
	手机号段记录条数：364528

1712:

	归属地信息库文件大小：3,372,351 字节
	归属地信息库最后更新：2017年12月
	手机号段记录条数：373606
	



output.txt 格式说明:

	<手机号前7位>|<卡类型>|<归属地省份>|<归属地城市>|<邮编>|<区号>

	1300010|2|北京|北京|100000|010


卡类型(isp):

		CMCC:   "中国移动",      //1
		CUCC:   "中国联通",      //2
		CTCC:   "中国电信",      //3
		CTCC_v: "中国电信虚拟运营商", //4
		CUCC_v: "中国联通虚拟运营商", //5
		CMCC_v: "中国移动虚拟运营商", //6