/*
 * @Author: xiaohuyi
 * @Date: 2022-04-14 14:26:54
 * @Description:
 */
package handler

type Handler interface {
	RecordOffset(offset int64, taskID string)             //由调用方去自己实现记录offset的方式
	HandlerBefore() interface{}                           // 正式handle之前的hook
	Handler(line string, opts ...interface{}) interface{} // handle逻辑
	HandlerAfter() interface{}                            // 正式handle之后的hook
	StartHandler(line string)
}
