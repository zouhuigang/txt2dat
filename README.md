### 2018-12-18 更新手机号库


### 文件

	本库只保留一份最新的output.txt和phone.dat，历史版本数据会被移除。


### 基本格式说明

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


##### 生成phone.dat：

### 将获取到的文件导入到sql中,然后格式化获取Phone.txt

### 目标output.txt ,格式如下:手机号|卡类型|省份|城市|邮编|区号
	
	1300000|2|山东|济南|250000|0531
	1300001|2|江苏|常州|213000|0519
	1300002|2|安徽|巢湖|238000|0551
	1300003|2|四川|宜宾|644000|0831
	1300004|2|四川|自贡|643000|0813
	1300005|2|陕西|西安|710000|029


导入之后，核对条数:

	SELECT count(*) FROM `phone`;

输出:

	415967


### 修改isp为符合的数据

	UPDATE phone set isp="1" where isp="移动" =>202553
	UPDATE phone set isp="2" where isp="联通" =>109605
	UPDATE phone set isp="3" where isp="电信" =>81998

	UPDATE phone set isp="4" where isp="物联网/电信" =>150
	UPDATE phone set isp="6" where isp="虚拟/移动" =>5217
	UPDATE phone set isp="4" where isp="虚拟/电信" =>2874
	UPDATE phone set isp="5" where isp="虚拟/联通"=>12970
	UPDATE phone set isp="4" where isp="卫星/电信" =>600
	

### 查询是否还有含有中文的isp没有被替换

	SELECT isp FROM phone WHERE length(isp)!=char_length(isp)  LIMIT 10;
	
	
直到查询为NULL,结束。


output.txt格式数据:

	SELECT phone,isp,province,city,post_code,city_code FROM `phone` where 1=1  ORDER BY phone asc LIMIT 10;

输出:

	phone	isp	province	city	post_code	city_code
	1300000	2	山东	济南	250000	0531
	1300001	2	江苏	常州	213000	0519
	1300002	2	安徽	巢湖	238000	0551
	1300003	2	四川	宜宾	644000	0831
	1300004	2	四川	自贡	643000	0813
	1300005	2	陕西	西安	710000	029
	1300006	2	江苏	南京	210000	025
	1300007	2	陕西	西安	710000	029
	1300008	2	湖北	武汉	430000	027
	1300009	2	陕西	西安	710000	029


### 然后将查询结果导出为output.csv


output.csv用记事本打开，替换掉其中的字符


	符号" => 空的(没有任何字符)。
	符号, => |


然后保存为output.txt。



### 运行程序，生成phone.dat文件


修改main.go中的版本号：

	
	//开始写入数据
	HeaderData.Write([]byte("1712")) //2017年12月

=>

	//开始写入数据
	HeaderData.Write([]byte("1812")) //2018年12月



linux编译，生成main

	./main

即可，生成phone.dat。

	






---

windows编译64位：

	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build  main.go

windows编译32位：
	
	CGO_ENABLED=0 GOOS=windows GOARCH=386  go build  main.go

mac版：

	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go

linux版：

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go



### 小提示


txt转excel的时候，需要在excel中设置单元格格式为文本格式，否则后面的区号021等，会变成21了


excel转cvs，需要用记事本打开cvs，然后另存，选择utf-8编码


	SELECT count(id) FROM `phonedat`;

输出：

	373606
	

phone.dat格式:

	SELECT phone,isp,province,city,post_code,city_code FROM `phonedat` where 1=1 LIMIT 10;




outdat.txt生成流程:

	excel->cvs->txt

phone.dat生成流程:

	txt->程序->dat

