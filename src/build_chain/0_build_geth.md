# PrivatePOAEthereum

## 正式使用硬件需求

Hardware Requirements
Minimum:

CPU with 2+ cores
4GB RAM
1TB free storage space to sync the Mainnet
8 MBit/sec download Internet service
Recommended:

Fast CPU with 4+ cores
16GB+ RAM
High-performance SSD with at least 1TB of free space
25+ MBit/sec download Internet service

出块节点： 开放30303 （测试可以开放8545，后续需要关闭）
同步节点：开放30303 8545

## 安装golang gcc

```shell
sudo apt install build-essential
```
## 两种方式编译Geth工具
### 1- 从源码编译
```shell
git clone https://github.com/ethereum/go-ethereum.git
``` 

编译geth工具

```shell
 cd go-ethereum && make geth
```
### 2- 直接下载geth工具
```shell
wget https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.12.1-9c216bd6.tar.gz
```
解压文件夹
```shell
tar -xzf geth-linux-amd64-1.12.1-9c216bd6.tar.gz
```
## 写入geth环境变量
```shell
vi ~/.bashrc
```
```shell
export ETHPATH=YouPathToGeth
export PATH=$ETHPATH:$PATH
```
```shell
source ~/.bashrc
```

## 生成账户地址

私钥通过用户输入的密码加密存储,可以通过golang解析出账户地址和私钥：
```text
geth account new --datadir /opt/etherData
```
> Address: 0x430CbEEffa18BD7ad0Ae5BAc062f130b6c8129B6

```shell
geth account new --datadir /opt/etherData
```
> Address: 0x413E129dD6b217E4a8702821ee069e1929D17c6a

```shell
geth account new --datadir /opt/etherData
```
> Address: 0x9449202f3E28Dd4595b0BE3c1736922Ba5aAce71

## 构造创世区块
```text
1. chainID:自定义的链ID
2. homesteadBlock、eip150Block、eip155Block、eip158Block、byzantiumBlock、constantinopleBlock、petersburgBlock：各项提案和升级的区块高度
3. period:出块时间间隔，0为不允许出空交易块，会等待有交易才出块
4. epoch:更新出块节点列表的周期
5. difficulty：POA下无作用
6. gasLimit：gasLimit限制
7. extradata：POA模式下用来指定验证者地址,账户地址去掉0x后加上64个前缀0和130个后缀0，比如0x+（64个前缀0）+5534F5024146D16a5C1ce60A9f5a2e9794e3F981+（130个后缀0）
8. alloc：用来预置账号以及账号的以太币数量，比如预置0x0B587FFD0BBa122fb5ddc19AD6eEcEB1D2dBbff7地址拥有1000ETH（1000*10^18WEI）
```
genesis.json
```json
{
   "config":{
      "chainId":12345,
      "homesteadBlock":0,
      "eip150Block":0,
      "eip155Block":0,
      "eip158Block":0,
      "byzantiumBlock":0,
      "constantinopleBlock":0,
      "petersburgBlock":0,
      "istanbulBlock":0,
      "berlinBlock":0,
      "londonBlock": 0,
      "clique":{
         "period":5,
         "epoch":300
      }
   },
   "alloc":{
      "0x6593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f":{
         "balance":"1000000000000000000000"
      },
      "0x6C345f0771a2f2B2694f97522D3371bF87b6BDF9":{
         "balance":"1000000000000000000000"
      },
      "0xab6bbb89eFd62dF605C881E692960a4951238D71":{
         "balance":"1000000000000000000000"
      }
   },
  "coinbase": "0x6593B47be3F4Bd1154c2faFb8Ad4aC4EFddD618f",
  "difficulty": "1",
  "extraData": "0x0000000000000000000000000000000000000000000000000000000000000000验证者地址0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
  "gasLimit": "80000000",
  "nonce": "0x0000000000000000",
  "mixhash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "timestamp": "0x00"
}
```