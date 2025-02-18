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
## 编译Geth工具
### 从源码编译
```shell
git clone https://github.com/ethereum/go-ethereum.git
``` 

编译geth工具

```shell
 cd go-ethereum && make geth
```
### 直接下载geth工具
```shell
wget https://gethstore.blob.core.windows.net/builds/geth-linux-amd64-1.12.1-9c216bd6.tar.gz
```
解压文件夹
```shell
tar -xzf geth-linux-amd64-1.12.1-9c216bd6.tar.gz
```
## 写入geth的环境变量
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
## 节点启动命令

### 创建创世区块
```shell
geth --datadir /opt/etherData init /opt/etherData/genesis.json
```
### 节点启动
1. 节点正常启动
```golang
nohup geth --identity "myethereum" --datadir /opt/etherData --allow-insecure-unlock --networkid 12345 --http --http.addr 0.0.0.0  --http.corsdomain "*" --ws --ws.addr 0.0.0.0 --ws.origins "*"  --http.api "eth,net,debug,txpool,web3,personal,admin,miner"  --rpc.enabledeprecatedpersonal --miner.gaslimit 80000000 --syncmode "full" --nodiscover --rpc.enabledeprecatedpersonal >> geth.log 2>&1 &
```

2. 连接控制台
```shell
   geth attach http://localhost:8545
```
或者ipc连接
```shell
     geth attach /opt/etherData/node0/geth.ipc
```
3. 根据节点私钥导入账号，提供节点私钥、加密节点私钥的对称密钥
```golang
personal.importRawKey("","")
```
4. 设置矿工地址(验证者地址)
```golang
miner.setEtherbase(eth.accounts[0])
```
5. 查看矿工账户
```golang
eth.coinbase
```
6. 解锁账户--需要指定时间，默认解锁300s
personal.unlockAccount(address, passphrase, duration),密码和解锁时长都是可选的。如果密码为null，控制台将提示交互输密码。解密的密钥将保存在内存中直到解锁周期超时，默认的解锁周期为300秒。将解锁周期设置为0秒将解锁该密钥直到退出geth程序。
```golang
personal.unlockAccount(eth.accounts[0],'passward',0)
```

矿工设置GasLimit
```golang
miner.setGasLimit(80000000)
```

7. 启动挖矿（start（） 的参数表示挖矿使用的线程数）/关闭挖矿
```golang
miner.start()
```
```golang
miner.stop()
```
8. 查看节点信息
```golang
admin.nodeInfo.enode
```
通过addPeer命令添加节点.
```golang
admin.addPeer("节点信息")
```
### 查询链上数据--连接控制台后的操作
1. 查看所有账户列表
```golang
eth.accounts
```
2. 查看所有账户余额
```golang
eth.getBalance(eth.accounts[0])
```
```golang
balanse=web3.fromWei(eth.getBalance(eth.accounts[0]),'ether')
```
3. 查询区块高度
```golang
eth.blockNumber
```
4. 交易操作
涉及链上交易时，需要先解锁账户。
```golang
eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[1],value:web3.toWei(4,'ether')})
```
5. 根据交易hash查询交易数据
```golang
eth.getTransaction("TxHash")
```
## 连接小狐狸钱包
### 连接网络
RPC： http://xxx:8545
chainId: 12345
名称和代币符号任意。
### 导入节点
输入节点私钥导入节点账户

## 新机器操作
### 数据备份节点
1. 相同命令启动后，连接控制台
   ```golang
   geth attach http://localhost:8545
   ```
   或者ipc连接
   ```golang
     geth attach /opt/etherData/node0/geth.ipc
   ```
2. 建立连接
```golang
admin.nodeInfo.enode
```
通过addPeer命令添加节点.
```golang
admin.addPeer("节点信息")
```
### 出块节点（创世区块中的验证者地址）
1. 连接控制台
2. 根据节点私钥导入账号，提供节点私钥、加密节点私钥的对称密钥
```golang
personal.importRawKey("08a7533871d3a2e01d3a8849320cbfb703eb20c5dd2a9ccd2d9780eba5659c8e","KEY")
```
3. 设置矿工地址(验证者地址)
```golang
miner.setEtherbase(eth.accounts[0])
```
4. 查看矿工账户
```golang
eth.coinbase
```
5. 解锁账户--需要指定时间，默认解锁300s
personal.unlockAccount(address, passphrase, duration),密码和解锁时长都是可选的。如果密码为null，控制台将提示交互输密码。解密的密钥将保存在内存中直到解锁周期超时，默认的解锁周期为300秒。将解锁周期设置为0秒将解锁该密钥直到退出geth程序。
```golang
personal.unlockAccount(eth.accounts[0],'passward',0)
```
6. 启动挖矿
```golang
miner.start()
```

### 出块节点（新的验证者地址）
新机器执行以下操作
1. 生成节点
```golang
geth account new --datadir /opt/etherData
```
> Address: 0x68d866baAfa993bc002cd35218c13f10aC54221d

> PrivateKey: e5ff5392711a137f3a4ac680e85ed29cb896427e89b4e0aa582b785722a84c49

2. 连接控制台
3. 根据节点私钥导入新的验证者账号，提供节点私钥、加密节点私钥的对称密钥
```golang
personal.importRawKey("e5ff5392711a137f3a4ac680e85ed29cb896427e89b4e0aa582b785722a84c49","KEY")
```
4. 在导入创世区块验证者节点的机器上执行以下操作：（需要半数以上的验证者同意）

目前的验证节点通过发起提案增加出块节点，增加后的节点和当前的验证者轮流出块。

```shell
clique.propose("新机器上生成的新验证者地址",true)
```

回到新服务器的终端：

1. 在新服务器上设置矿工地址(新的验证者地址)
```golang
miner.setEtherbase(eth.accounts[0])
```
2. 查看矿工账户
```golang
eth.coinbase
```
3. 解锁账户--需要指定时间，默认解锁300s
personal.unlockAccount(address, passphrase, duration),密码和解锁时长都是可选的。如果密码为null，控制台将提示交互输密码。解密的密钥将保存在内存中直到解锁周期超时，默认的解锁周期为300秒。将解锁周期设置为0秒将解锁该密钥直到退出geth程序。
```golang
personal.unlockAccount(eth.accounts[0],'passward',0)
```
4. 启动挖矿（start（） 的参数表示挖矿使用的线程数）/关闭挖矿
```golang
miner.start()
```

## 关闭节点
```golang
ps aux | grep geth | grep -v grep | awk '{print $2}'| xargs kill -15
```

## 清除链数据
```golang
geth removedb --datadir "/opt/etherData/"
```


nohup geth --identity "myethereum" --datadir /opt/etherData/ --networkid 12345 --authrpc.port 8551 --http --http.port 8545  --http.corsdomain "*" --ws --ws.port 8546 --ws.addr 0.0.0.0 --ws.origins "*" --port 30303 --http.addr 0.0.0.0 --http.api "eth,net,web3,personal,admin,miner,debug,txpool" --allow-insecure-unlock --rpc.enabledeprecatedpersonal --syncmode "full" --mine --miner.etherbase 0x430CbEEffa18BD7ad0Ae5BAc062f130b6c8129B6 --unlock 0x430CbEEffa18BD7ad0Ae5BAc062f130b6c8129B6 --keystore /opt/etherData/node0/keystore/ --password /opt/etherData/node0/password.txt --nodiscover >> geth1.log 2>&1 &
