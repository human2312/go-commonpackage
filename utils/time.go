/**
 * @Author: 作者名
 * @Description: 描述
 * @File: tims //文件名称
 * @Version: 1.0.0
 * @Date: 2021/9/23 8:41 下午  //时间
 * @Software : GoLand //开发软件【可选】
 */
package utils

import "time"

type Time struct {
}

func (this *Time) Now() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
