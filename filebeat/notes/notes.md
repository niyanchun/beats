## 模块

filebeat/input代表每一个类型，比如log、stdin、redis，各个子模块实现了自己的input。

## 流程

1. libbeat中各种公共流程，最终调用`filebeat.go#Run`开始启动Filebeat
2. 