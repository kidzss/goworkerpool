# goworkerpool

思路参考
http://marcio.io/2015/07/handling-1-million-requests-per-minute-with-golang/

其实就是个多worker和调度器的基本框架

会将每一个请求放到独立的worker goroutine里去处理

用channel来控制worker的数量，需要时从channel中取出worker

任务来了时的流程：

-> 先放进全局的jobqueue

-> 在dispatcher的for select loop中，监听到全局jobqueue有job进入

-> dispatcher的for loop启动一个新的goroutine，从workerpool拿出一条jobchannel(这个jobchannel是在worker启动阶段被塞进去的)，并把job传给这个goroutine

-> 将这个job放到取出来的jobchannel里去

-> worker的for select loop发现有job来了，取出开始处理


**dispatcher的事件loop一个是在goroutine中执行，但另一个不是，需要研究一下为什么**


####这个东西可以做什么
可以拿来做一些小工具，比如异步任务执行，异步发短信什么的
如果需要实现任务可追溯，最好消息还是要落地，不要程序一崩溃什么都没了。
