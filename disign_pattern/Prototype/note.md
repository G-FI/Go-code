拥有一个接口就能进行clone出一模一样的对象
优点:
    1. 不需要关心实现接口的concrete类到底是什么，就能clone出这个对象，**与所属的具体类解耦合**
    2. 将clone的工作交给框架去做，客户端只需要调用clone方法就行
方法：
    1. 可以用interface，也可以用继承，但前提是都有一个clone方法供客户端调用
    2. clone方法也只是返回一个接口类型，客户端全程不与具体类打交道
    3. clone实现：在clone中创建本类对象，然后将各个field进行赋值，如果是有拷贝构造函数的话，就直接返回 new ConcreteClass(this)

