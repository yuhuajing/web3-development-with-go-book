
# 启动节点
## 创建创世区块
```shell
geth --datadir /opt/etherData init ./genesis.json
```
## 节点启动
节点正常启动
```golang
nohup geth --identity "myethereum" --datadir /opt/etherData --allow-insecure-unlock --networkid 12345 --http --http.addr 0.0.0.0  --http.corsdomain "*" --ws --ws.addr 0.0.0.0 --ws.origins "*"  --http.api "eth,net,debug,txpool,web3,personal,admin,miner"  --rpc.enabledeprecatedpersonal --miner.gaslimit 80000000 --syncmode "full" --nodiscover --rpc.enabledeprecatedpersonal >> geth.log 2>&1 &
```
连接控制台
```shell
   geth attach http://localhost:8545
```
或者ipc连接
```shell
     geth attach /opt/etherData/node0/geth.ipc
```
## 查询链上状态
1. 查看所有账户列表
```golang
eth.accounts
```
2. 根据节点私钥导入账号，提供节点私钥、加密节点私钥的对称密钥
```golang
personal.importRawKey("","")
```
3. 查看账户余额
```golang
eth.getBalance(eth.accounts[0])
```
```golang
balanse=web3.fromWei(eth.getBalance(eth.accounts[0]),'ether')
```
4. 查询区块高度
```golang
eth.blockNumber
```
5. 根据交易hash查询交易数据
```golang
eth.getTransaction("TxHash")
```
6. 查看节点信息
```golang
admin.nodeInfo.enode
```
## 发送链上交易
1. 发送链上交易时，需要先解锁账户。
```golang
eth.sendTransaction({from:eth.accounts[0],to:eth.accounts[1],value:web3.toWei(4,'ether')})
```
2. 解锁账户--需要指定时间，默认解锁300s
personal.unlockAccount(address, passphrase, duration),密码和解锁时长都是可选的。如果密码为null，控制台将提示交互输密码。解密的密钥将保存在内存中直到解锁周期超时，默认的解锁周期为300秒。将解锁周期设置为0秒将解锁该密钥直到退出geth程序。
```golang
personal.unlockAccount(eth.accounts[0],'passward',0)
```
## 启动出块
1. 设置矿工地址(验证者地址)
```golang
miner.setEtherbase(eth.accounts[0])
```
2. 查看矿工账户
```golang
eth.coinbase
```
3. 设置区块GasLimit
```golang
miner.setGasLimit(80000000)
```
4. 启动出块（start（） 的参数表示出块使用的线程数）/关闭出块
必须先解锁矿工账户，否则不会启动出块
```golang
miner.start()
```
```golang
miner.stop()
```
## 关闭节点
```golang
ps aux | grep geth | grep -v grep | awk '{print $2}'| xargs kill -15
```

## 清除链数据
```golang
geth removedb --datadir "/opt/etherData/"
```