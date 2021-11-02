# TuuzGoWeb
TuuzGoWeb基于Gin，四层写法，Gorose数据库，离合transaction写法


基于TuuzGo开发脚手架，要不是go语言err！=nil不然开发效率就能更高了，C#是真好……加上dynamic类型，感觉用起来比interface{}舒服

# Gorose-Pro

GorosePro为TuuzGoWeb的ORM支持，支持nested transaction，对复杂逻辑下的解耦有非常重要的帮助


# TGW目录说明

- app
  - 这是所有的代码逻辑控制面
- common
  - 这是所有本项目里面（移植项目）里面会用到的控制面
- config
  - 设定面
- extend
  - 插件控制面
- route
  - 路由控制面
- tuuz
  - 框架面


# TGW的四段路由说明

~~~
{{host}}/v1/RouterName/ControllerName/FuncName
~~~

- 我强烈推荐使用四段路由，四段路由可以将你的动作进行有效的切分
- 四段路由在Thinkphp的默认三段的基础上，加入了版本属性
- 虽然你可以使用版本控制或者注入来完成新接口的切换，但是这并不适用于复制型的项目 
- 使用四段地址，可以让你从"前端"对接中解放出来，如果你愿意你可以让你的项目有多个版本共存
- 前端合作中需要使用什么功能，直接翻之前的接口文档对接即可，这是TGW选择使用四段路由的重要原因！

当然如果你相对OldSchool，你也可以将所有的Controller路由放到Onroute中

## TGW的route用法说明
- OnRoute是路由的入口，有且只有一个Onroute.go文件，在Onroute中，我们需要定义版本以及下级（版本）路由的入口
- v1文件夹同app文件夹中的v1，目的是将所有属于v1这个router的下的路由全部集合
- 在v1文件夹下，你可以定义多个route，请使用xxxRouter结尾，这个习惯有利于ide区别controller和router，命名不会影响功能，但是会影响日后接手人员的心情
- 编写方法可以参考示例
### TGW的route目录构建逻辑
- route
  - v1
    - indexRouter.go
    - userRouter.go
  - v2
    - indexRouter.go
    - balanceRouter.go
  - OnRoute.go

## TGW的app用法说明
- 写法以及构建方法可参考示例
- 这里强调下，当你写了一个controller之后，你需要将这个controller的名称写入到对应的router中方可使用
- TGW推荐使用MAC方式开发
- 传统框架只能使用MVC方式，复杂业务中数据库的单例很难实现，这是因为原版的MysqlDriver对NestedTransaction支持很差
- 本框架使用Gorose-Pro，支持NestedTransaction等操作，你可以放心大胆的使用MAC架构进行解耦！
- 如果有必要，你可以在MAC的基础上加入Util或者Logic，做成（伪）四层的形式

### TGW的app目录构建逻辑
- app
  - v1
    - index
      - action
      - controller
        - index.go
      - model
        - IndexModel
          - IndexModel.go
    - user
      - action
      - controller
        - user.go
      - model
        - UserModel
          - UserModel.go
        - UserInfoModel
          - UserInfoModel.go
  - v2
      - index
        - action
        - controller
            - index.go
        - model
            - IndexModel
                - IndexModel.go
      - balance
          - action
          - controller
              - index.go
          - model
              - BalanceModel
                  - BalanceModel.go
  - cron
    - InvestCron
      - invest.go
      - timeinvest.go

### TGW的其他说明
接口性能：
- 在本地数据库的情况下，平均（列表x20/单条数据)，平均效率在6ms-30ms左右
- 在RDS数据库情况下，平均（列表x20/单条数据)，平均效率在1ms-10ms左右

## TGW输出说明
- 你可以使用Gin的默认输出
- Context:示例中的c为*gin.context
- Code:错误码可以与前端自定
- Data:如果传输nil，最终json中的data字段会变成"[]"
- Echo:这里只能传string

```
GinContext，错误码，数据Array，操作指示（成功失败）

RET.Success(c, 0, nil, "验证成功")
RET.Fail(c, 400, nil, "验证失败")
```
