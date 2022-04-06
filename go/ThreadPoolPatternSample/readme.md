Worker Pool (aka Thread Pool) is a pattern to achieve concurrency using fixed number of workers to execute multiple amount of tasks on a queue. In Go ecosystem, we use goroutines to spawn the worker and implement the queue using channel. The defined amount of workers will pull the task from queue, and finish up the task, and when the task has been done, the worker will keep pulling the new one until the queue is empty.

Ok, I know what is Worker Pool, but why we need Worker Pool on my code? What is the difference when I just spawn new goroutine if I need the process to run concurrently?
You don’t have unlimited resource on your machine, the minimal size of a goroutine object is 2 KB, when you spawn too many goroutine, your machine will quickly run out of memory and the CPU will keep processing the task until it reach the limit. By using limited pool of workers and keep the task on the queue, we can reduce the burst of CPU and memory since the task will wait on the queue until the the worker pull the task.

Reference:

https://medium.com/code-chasm/go-concurrency-pattern-worker-pool-a437117025b1