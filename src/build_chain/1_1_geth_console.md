
# 启动节点2
## 新机器操作
### 数据备份节点
1. 拷贝 genesis.json,创建相同的0号区块
2. 节点正常启动
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
3. 建立连接
查询节点信息
```golang
admin.nodeInfo.enode
```
通过addPeer命令添加节点.
```golang
admin.addPeer("节点信息")
```

### 出块节点（新的验证者地址）
新机器执行以下操作
1. 生成节点
```golang
geth account new --datadir /opt/etherData
```
> Address: 0x68d866baAfa993bc002cd35218c13f10aC54221d

2. 连接控制台
3. 在已连接的节点上执行以下操作：（需要半数以上的验证者同意）

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
3. 解锁账户
```golang
personal.unlockAccount(eth.accounts[0],'passward',0)
```
4. 启动挖矿（start（） 的参数表示挖矿使用的线程数）/关闭挖矿
```golang
miner.start()
```