/*
通过改变offset文件偏移量，将数据写入文件中的任意位置
作者:邹慧刚
github:zouhuigang
*/
package model

import "io"

type Writer struct {
	Offset int64
	Wat    io.WriterAt //准备写入的文件，打开状态
}

var M_Writer = Writer{}

//初始化，offset的值
func (this *Writer) NewWriter(w io.WriterAt, offset int64) io.Writer {
	return &Writer{offset, w}
}

//写数据,返回写入的长度。然后再次更新偏移量
func (this *Writer) Write(b []byte) (n int, err error) {
	n, err = this.Wat.WriteAt(b, this.Offset)
	this.Offset += int64(n)
	return
}
