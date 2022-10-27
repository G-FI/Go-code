1. 模板方法模式是一种行为设计模式， 它在**超类中定义了一个算法的框架**， 允许子类在**不修改结构的情况下**重写算法的特定步骤。
2. 比如编译器，框架就定义了具体算法步骤lex，parseAST，emitASM，以及一个compile模板方法，其中按特定顺序调用lex(),parseAST,emitASM，子类不能修改模板方法，但是对于不同的编程语言同一个体系结构，需要特定子类进行重写其中的lex和parseAST，实现对这个特定语言的扫描以及解析关键字、表达式，此例中由于体系结相同，还可以在基类中提供emitASM这个特定步骤的默认实现
3. 当你只希望客户端**扩展某个特定算法步骤**， **而不是整个算法或其结构时**， 可使用模板方法模式。---> 让算法结构保持完整

实现：
1. 看算法是否可以拆分，将拆分出来的步骤按逻辑组合就形成了**模板方法**
2. 从所有子类出发，对于拆分出来的每个步骤，将相同的实现提取出来放到基类中
3. 