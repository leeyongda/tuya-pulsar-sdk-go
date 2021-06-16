# tuya-pulsar-sdk-go
tuya-pulsar-sdk-go

### 基于官方SDK封装
[依赖库]
https://github.com/apache/pulsar-client-go
## 使用前准备

1. AccessID：由涂鸦平台提供
2. AccessKey：由涂鸦平台提供
3. pulsar地址：根据不同的业务区域选择 Pulsar 地址。可以从涂鸦对接文档中查询获取。

参数说明如下：
1. AccessID：填写云开发平台中 API 授权密钥的 Access ID。
2. AccessKey：填写云开发平台中 API 授权密钥的 Access Secret。
3. Pulsar URL：根据调用的区域进行选择
    > CN_SERVER_URL(中国区)：pulsar+ssl://mqe.tuyacn.com:7285/

    > US_SERVER_URL(美国区)：pulsar+ssl://mqe.tuyaus.com:7285/

    > EU_SERVER_URL(欧洲区)：pulsar+ssl://mqe.tuyaeu.com:7285/

    > IND_SERVER_URL(印度区)：pulsar+ssl://mqe.tuyain.com:7285/

### [example](https://github.com/leeyongda/tuya-pulsar-sdk-go/blob/main/example)

## 注意事项

1. 确保accessID，accessKey是正确的
2. 确保pulsar地址是正确的
3. 尽量确保你使用的sdk代码版本是最新的

## 涂鸦技术支持

你可以通过以下方式获得Tua开发者技术支持：

- [涂鸦帮助中心](https://support.tuya.com/zh/help)
- [涂鸦技术工单平台](https://iot.tuya.com/council)
- [基于 Puslar SDK 获取消息推送](https://developer.tuya.com/cn/docs/iot/Puslar-SDK-get-message?id=Kan0klj9qbv3l)
