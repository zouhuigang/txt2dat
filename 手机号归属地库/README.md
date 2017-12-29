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

