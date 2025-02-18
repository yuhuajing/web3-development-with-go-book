#  Web3-development-with-go-book
Welcome to the world of decentralized blockchain: The document begins with an introduction to blockchain technology and Ethereum, providing essential context for understanding smart contracts. It then delves into the syntax and features of Solidity, covering key concepts such as data types, functions, modifiers, and inheritance. Practical examples are included to illustrate how to write and deploy smart contracts, along with best practices for security and optimization. Additionally, the document addresses common pitfalls and debugging strategies to help learners navigate challenges they may encounter. Finally, the document provides resources for further learning, including links to online courses, documentation, and community forums. This structured approach aims to equip readers with the knowledge and skills needed to confidently create and manage their own smart contracts in Solidity

This book will guide you through the development of a decentralized application, including:
- smart-contract development (in [Solidity](https://docs.soliditylang.org/en/latest/index.html));

**This book is not for complete beginners.**

I expect you to be an experienced developer, who has ever programmed in any programming language. It'll also be helpful if you know [the syntax of Solidity](https://docs.soliditylang.org/en/v0.8.17/introduction-to-smart-contracts.html), the main programming language of this book. If not, it's not a big problem: we'll learn a lot about Solidity and Ethereum Virtual Machine during our journey.

**However, this book is for blockchain beginners.**

If you only heard about blockchains and were interested but haven't had a chance to dive deeper, this book is for you!  Yes, for you! You'll learn how to develop for blockchains (specifically, Ethereum), how blockchains work, how to program and deploy smart contracts, and how to run and test them on your computer.

Alright, let's get started!

## Useful Links
1. This book is hosted on GitHub: <https://github.com/yuhuajing/web3-development-with-go-book>

## Table of Contents
- Milestone 0. Base Golang Knowledge
  1. init
  2. rune
  3. slice
  4. data-mapping
  5. data-channel
  6. func
  7. defer
  8. GMP
  9. GC
- Milestone 1. Golang Blockchain
  1. initBlockchainConn
  2. getAccountBalance
  3. getAccountCodes
  4. generateAccount
  5. getBlockchainBlocks
  6. getBlockchainTransactions
  7. subscribeBlockchainNewHead
  8. subscribeBlockchainNewTransactions
  9. subscribeBlockchainNewLogs
- Milestone 2. Golang Blockchain Transactions
  1. buildRawTransactions
  2. callContractByRawTransactions
  3. estimateTransactionsGas
  4. readContractByABI
  5. writeContractByABI
- Milestone 3. Golang Blockchain Contract Slot
  1. slotKnowledge
  2. staticUint
  3. staticInt
  4. staticBool
  5. staticAddress
  6. staticBytes
  7. staticArray
  8. staticStruct
  9. staticString
  10. staticMapping
- Milestone 4. Golang Blockchain Tools
  1. merkleProof
  2. signature

## Running locally

To run the book locally:
1. Install Rust.
- `curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh`
2. Install mdBook:
```shell
    cargo install mdbook
    cargo install mdbook-katex
```
3. Clone the repo:
```shell
    git clone https://github.com/yuhuajing/web3-development-with-go-book.git
    cd web3-development-with-go-book
    mdbook serve --open
```
4. Visit http://localhost:3000/ 
