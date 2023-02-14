# go-trongrid-api
Base:https://github.com/fbsobreira/gotron-sdk

## 成本估算

### 合约方式
- 部署子合约：111.1992 TRX+314带宽+397140能量

### TRC20转账
- 激活账户有1.1TRX手续费
- TRC20转账 345Bandwidth(0.345TRX) 13,045Energy(5.5TRX) 大约:5.85TRX
- 带宽价格：0.001TRX
- 能量价格：0.00042TRX 

## 波场账户管理
- 下载本地钱包工具：http://192.168.10.119/tronctl.zip
- 创建一个账户：`./tronctl keys add examplename` 得到`TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG`
- 水龙头领取矿工费 `https://t.me/troncoredevscommunity` `https://discordapp.com/invite/hqKvyAM`
```shell
!shasta TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG
```
- 区块浏览器 https://shasta.tronscan.org
- 查询余额
```shell
$ ./tronctl account balance TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG -n grpc.shasta.trongrid.io:50051
{
  "address": "TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG",
  "allowance": 0,
  "balance": 7911.198038,
  "rewards": 0,
  "type": 0
}
```
- 获得TRC20-USDT
- 查询TRC20-USDT余额：
```shell
$ ./tronctl trc20 balance TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG TGT6r3rS1tyA77B8LT1gE1pfMFxK9f2XJT -n grpc.shasta.trongrid.io:50051
{
  "balance": "999939898.8 xUSDT"
}
```
- 转账TRC20-USDT:从TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG转给TY9tyzi6xtt5s3QDfcK5yrcGX2ovHZPzEZ,1.23USDT（TGT6r3rS1tyA77B8LT1gE1pfMFxK9f2XJT）
```shell
$  ./tronctl trc20 send TY9tyzi6xtt5s3QDfcK5yrcGX2ovHZPzEZ 1.23 TGT6r3rS1tyA77B8LT1gE1pfMFxK9f2XJT -n grpc.shasta.trongrid.io:50051 -s TWva2iKq2uVZo51EsuUFtrkjaN5VL4K5PG
{
  "blockNumber": 30043699,
  "contractAddress": "TGT6r3rS1tyA77B8LT1gE1pfMFxK9f2XJT",
  "message": "",
  "receipt": {
    "energyFee": 7834960,
    "energyUsage": 0,
    "energyUsageTotal": 27982,
    "fee": 7834960,
    "netFee": 0,
    "netUsage": 345,
    "originEnergyUsage": 0
  },
  "resMessage": "",
  "success": true,
  "txID": "0x622ef0cdcc6528244c49f075ae1be3b47499d5d6d907bad58281c655463a9afc"
}
```