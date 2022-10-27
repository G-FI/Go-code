1.client需要使用一个如method(data)
3. 有一个服务类可以完成客户端的请求，但它的提供的方法时serviceMethod(special data)
问题：
   1. client需要的接口和第三方所提供的库接口不一样，但功能就是client想要的
   2. 
解决方案：将对method(data)的调用转化为对serviceMethod(special data)的调用
   1. 创建adaptor类，实现客户端method对应接口，此时就可以将adaptor传给client需要接口的地方了
   2. adaptor类封装一个需要被适配的对象引用
   3. 在method方法中，调用serviceMethod(special data)