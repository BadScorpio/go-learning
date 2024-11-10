const Web3 = require('web3');
const fs = require('fs');

// 连接到 ganache-cli
const web3 = new Web3(new Web3.providers.HttpProvider("http://127.0.0.1:8545"));

// 读取 ABI 和字节码
const abi = JSON.parse(fs.readFileSync("MyContract_sol_MyContract.abi", "utf-8"));
const bytecode = fs.readFileSync("MyContract_sol_MyContract.bin", "utf-8");

async function deploy() {
    // 获取可用账户
    const accounts = await web3.eth.getAccounts();
    const deployAccount = accounts[0];  // 使用第一个账户进行部署

    // 创建合约实例
    const contract = new web3.eth.Contract(abi);

    // 部署合约
    const deployedContract = await contract.deploy({
        data: "0x" + bytecode  // 加上 "0x" 前缀
    }).send({
        from: deployAccount,
        gas: 1500000,
        gasPrice: '30000000000'
    });

    console.log("合约部署地址:", deployedContract.options.address);
}

deploy();
