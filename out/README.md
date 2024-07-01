# 微信支付

## 启动

> 运行 pay.sh / pay.exe 会自动读取config.json的内容 和apiclient_key.pem 
>
> shell 会根据json内配置的端口启动http服务

## 注意事项

- 使用前请确保已获取config.json对应的参数
- 修改config.json参数时请勿修改任何的key
- pay.sh 只能在amd64下的linux环境运行
- 请确保config.json 在oss.sh / oss.exe同级目录下
- 配置http端口时请确保port的值为string类型

### API

> 成功启动通过 ip:port/uri 访问 API

| uri        | method | param                                                   |Content-Type | response  | desc  |
|------------| ------ |---------------------------------------------------------| -------- |-----------|-------|
| payment    | POST | openId:string, orderId:string, desc:string, price:int64 | application/json| 请查看微信官方文档 | 下单    |
| status     |POST| orderId:string                                          | application/json|   请查看微信官方文档        | 查询订单状态 |
| close | POST| orderId:string                                          |application/json| 200       | 关闭订单  |
| refund | POST| orderId:string, refundAmount:int64, totalAmount:int64   |application/json| 200       | 退款  |



> orderId = 开发者自己生成的订单号 <br/>
> 支付的价格单位为分 比如你传一个100 实际上付的是1.00 <br/>
> 退款时 totalAmount = 订单总价 refundAmount = 退款的价格  传入的值单位都是分 每笔订单不可重复退款 <br/>
>
> 具体可看微信支付apiv3的文档


