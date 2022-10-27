1. 备忘录模式是一种行为设计模式， 允许在不暴露对象实现细节的情况下保存和恢复对象之前的状态。
2. 需要有一个originator原发器，就像整个文字编辑器状态类，我们的**备忘录就是保存这个originator某个时刻的快照(memento的field和originator可以一样)**
3. originator有创建快照和从快照恢复的方法
4. 其他类不能访问memento的内部字段(封装originator的私有属性)，但originator需要访问(恢复快照时)->可以使用内部嵌套类
5. 需要一个负责人类，其中包含memento的栈结构，用于给originator记录状态以及恢复
6. 这个负责人类可以是一个command，对于text editor来说，每个UI对象封装一个command，UI对象可以将请求委托给command，command将请求再委托给真正的service，并且可以在委托之前生成快照，用于撤销操做
