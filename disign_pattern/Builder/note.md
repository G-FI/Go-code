1. 产品较为复杂且需要详细配置时，使用生成器,**将太多的参数从构造函数中移除**
2. 不同的产品的创建过程相同，只是对于配置的不同，调用的**配置过程有差别**
3. 基本生成器接口中定义了所有可能的制造步骤， 具体生成器将实现这些步骤来制造特定形式的产品。 同时， 主管类将负责管理制造步骤的顺序
4. 创建的基本元素：
    1. Builder接口 
    2. 用于生成不同product的实现Builder的具体类 
    3. Manager 对于一些常用的配置，在其中自动化调用Builer具体类的方法 
    4. 客户端创建具体Builder类，可选手动配置定制化的Product，也可以将Builer与Manager关联，自动化配置Product
    5. Builder具体类中聚合了一个product，在前面配置完之后用具体builder类的getProduct()返回构建的对象，如果所有的构建的product都遵循同一接口那么，可以直接在manager中返回构建结果
    6. 具体类builder如果getProduct()返回的product不相同，也可以用Builder接口转换成具体builder然后调用getProduct()方法，因为接口中不好声明返回值不同的方法
