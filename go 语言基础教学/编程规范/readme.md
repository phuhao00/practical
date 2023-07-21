## 编码规范

> mandatory （必须采用）<br>
> preferable (推荐，可以不采取) <br>
> optional （可选，自己决定）<br>



# case
* content
- gofmt,列<120
- import
  （
  "内部"
  "外部"
  "匿名"
  ）
* error
    - 参数返回放最后
    - 采用独立错误流进行处理，独立判断，不要与其他变量组合逻辑判断
* 单元测试
    - TestExample,TestFoo,TestStructFoo
* 注释
    - 每一个被导出的都应该有注释
    - 所有注释代码在提交code review 前都应该被删除，除非添加注释说明为什么不删除，并且标明后续处理方式
    - 每个package 都应该有注释 //Package ....
    - 结构体字段生僻字都应该有注释
    - 自定义类型注释，全局 常量，变量，都应该 有注释
* 命名规范
- 包命名（小写） case: syscall(systemcall)
  不要使用无意义的包名，eg:common,util ,应该清晰，符合单一职责；util/encryption 是允许的
- 文件 使用小写，下划线分割单词
- 结构体应该是名词或者名词短语，不要使用不明确的，eg:Data,Info,声明和初始化，多行展开格式
- 接口 单个函数的接口 er做为后缀，eg:Reader,Writer
- 变量 apiClient -APIClient,局部变量使用短命名

* 控制结构
- if   变量在👈，常量在👉
- for 采用短声明
- range 尽量丢弃
- switch 要有default
- return 尽早返回
- goto 禁止使用
* 函数
    - 参数和返回值，小写打头，数量<5，返回值不清晰的时候可以使用命名返回
    - 尽量用值传递
    - 参数是map,slice,chan,interface的时候不要使用指针
    - defer ,当存在资源管理的时候，紧跟资源释放，禁止在循环中使用
* 方法接收器
    - 结构体首字母小写命名，如果方法超过20行，不要使用单字母
* 代码行数(包括注释，空行)
- 文件不超过 800 row
- 函数不超过 80 row
- 嵌套 深度 < 4
