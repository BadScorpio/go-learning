const { buildModule } = require("@nomicfoundation/hardhat-ignition/modules");
module.exports = buildModule("erc20Module", (m) => { 
    
const erc20 = m.contract("erc20", []); 
return { erc20 }; 
});