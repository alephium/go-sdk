# Go API client for alephium

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.5.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import alephium "github.com/alephium/go-sdk"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), alephium.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), alephium.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), alephium.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), alephium.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://..*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*AddressesApi* | [**GetAddressesAddressBalance**](docs/AddressesApi.md#getaddressesaddressbalance) | **Get** /addresses/{address}/balance | Get the balance of an address
*AddressesApi* | [**GetAddressesAddressGroup**](docs/AddressesApi.md#getaddressesaddressgroup) | **Get** /addresses/{address}/group | Get the group of an address
*AddressesApi* | [**GetAddressesAddressUtxos**](docs/AddressesApi.md#getaddressesaddressutxos) | **Get** /addresses/{address}/utxos | Get the UTXOs of an address
*BlockflowApi* | [**GetBlockflowBlocks**](docs/BlockflowApi.md#getblockflowblocks) | **Get** /blockflow/blocks | List blocks on the given time interval
*BlockflowApi* | [**GetBlockflowBlocksBlockHash**](docs/BlockflowApi.md#getblockflowblocksblockhash) | **Get** /blockflow/blocks/{block_hash} | Get a block with hash
*BlockflowApi* | [**GetBlockflowBlocksWithEvents**](docs/BlockflowApi.md#getblockflowblockswithevents) | **Get** /blockflow/blocks-with-events | List blocks with events on the given time interval
*BlockflowApi* | [**GetBlockflowBlocksWithEventsBlockHash**](docs/BlockflowApi.md#getblockflowblockswitheventsblockhash) | **Get** /blockflow/blocks-with-events/{block_hash} | Get a block and events with hash
*BlockflowApi* | [**GetBlockflowChainInfo**](docs/BlockflowApi.md#getblockflowchaininfo) | **Get** /blockflow/chain-info | Get infos about the chain from the given groups
*BlockflowApi* | [**GetBlockflowHashes**](docs/BlockflowApi.md#getblockflowhashes) | **Get** /blockflow/hashes | Get all block&#39;s hashes at given height for given groups
*BlockflowApi* | [**GetBlockflowHeadersBlockHash**](docs/BlockflowApi.md#getblockflowheadersblockhash) | **Get** /blockflow/headers/{block_hash} | Get block header
*BlockflowApi* | [**GetBlockflowIsBlockInMainChain**](docs/BlockflowApi.md#getblockflowisblockinmainchain) | **Get** /blockflow/is-block-in-main-chain | Check if the block is in main chain
*ContractsApi* | [**GetContractsAddressState**](docs/ContractsApi.md#getcontractsaddressstate) | **Get** /contracts/{address}/state | Get contract state
*ContractsApi* | [**PostContractsCallContract**](docs/ContractsApi.md#postcontractscallcontract) | **Post** /contracts/call-contract | Call contract
*ContractsApi* | [**PostContractsCompileContract**](docs/ContractsApi.md#postcontractscompilecontract) | **Post** /contracts/compile-contract | Compile a smart contract
*ContractsApi* | [**PostContractsCompileProject**](docs/ContractsApi.md#postcontractscompileproject) | **Post** /contracts/compile-project | Compile a project
*ContractsApi* | [**PostContractsCompileScript**](docs/ContractsApi.md#postcontractscompilescript) | **Post** /contracts/compile-script | Compile a script
*ContractsApi* | [**PostContractsTestContract**](docs/ContractsApi.md#postcontractstestcontract) | **Post** /contracts/test-contract | Test contract
*ContractsApi* | [**PostContractsUnsignedTxDeployContract**](docs/ContractsApi.md#postcontractsunsignedtxdeploycontract) | **Post** /contracts/unsigned-tx/deploy-contract | Build an unsigned contract
*ContractsApi* | [**PostContractsUnsignedTxExecuteScript**](docs/ContractsApi.md#postcontractsunsignedtxexecutescript) | **Post** /contracts/unsigned-tx/execute-script | Build an unsigned script
*EventsApi* | [**GetEventsBlockHashBlockhash**](docs/EventsApi.md#geteventsblockhashblockhash) | **Get** /events/block-hash/{blockHash} | Get contract events for a block
*EventsApi* | [**GetEventsContractContractaddress**](docs/EventsApi.md#geteventscontractcontractaddress) | **Get** /events/contract/{contractAddress} | Get events for a contract within a counter range
*EventsApi* | [**GetEventsContractContractaddressCurrentCount**](docs/EventsApi.md#geteventscontractcontractaddresscurrentcount) | **Get** /events/contract/{contractAddress}/current-count | Get current value of the events counter for a contract
*EventsApi* | [**GetEventsTxIdTxid**](docs/EventsApi.md#geteventstxidtxid) | **Get** /events/tx-id/{txId} | Get contract events for a transaction
*InfosApi* | [**GetInfosChainParams**](docs/InfosApi.md#getinfoschainparams) | **Get** /infos/chain-params | Get key params about your blockchain
*InfosApi* | [**GetInfosCurrentHashrate**](docs/InfosApi.md#getinfoscurrenthashrate) | **Get** /infos/current-hashrate | Get average hashrate from &#x60;now - timespan(millis)&#x60; to &#x60;now&#x60;
*InfosApi* | [**GetInfosDiscoveredNeighbors**](docs/InfosApi.md#getinfosdiscoveredneighbors) | **Get** /infos/discovered-neighbors | Get discovered neighbors
*InfosApi* | [**GetInfosHistoryHashrate**](docs/InfosApi.md#getinfoshistoryhashrate) | **Get** /infos/history-hashrate | Get history average hashrate on the given time interval
*InfosApi* | [**GetInfosInterCliquePeerInfo**](docs/InfosApi.md#getinfosintercliquepeerinfo) | **Get** /infos/inter-clique-peer-info | Get infos about the inter cliques
*InfosApi* | [**GetInfosMisbehaviors**](docs/InfosApi.md#getinfosmisbehaviors) | **Get** /infos/misbehaviors | Get the misbehaviors of peers
*InfosApi* | [**GetInfosNode**](docs/InfosApi.md#getinfosnode) | **Get** /infos/node | Get info about that node
*InfosApi* | [**GetInfosSelfClique**](docs/InfosApi.md#getinfosselfclique) | **Get** /infos/self-clique | Get info about your own clique
*InfosApi* | [**GetInfosUnreachable**](docs/InfosApi.md#getinfosunreachable) | **Get** /infos/unreachable | Get the unreachable brokers
*InfosApi* | [**GetInfosVersion**](docs/InfosApi.md#getinfosversion) | **Get** /infos/version | Get version about that node
*InfosApi* | [**PostInfosDiscovery**](docs/InfosApi.md#postinfosdiscovery) | **Post** /infos/discovery | Set brokers to be unreachable/reachable
*InfosApi* | [**PostInfosMisbehaviors**](docs/InfosApi.md#postinfosmisbehaviors) | **Post** /infos/misbehaviors | Ban/Unban given peers
*MinersApi* | [**GetMinersAddresses**](docs/MinersApi.md#getminersaddresses) | **Get** /miners/addresses | List miner&#39;s addresses
*MinersApi* | [**GetWalletsWalletNameMinerAddresses**](docs/MinersApi.md#getwalletswalletnamemineraddresses) | **Get** /wallets/{wallet_name}/miner-addresses | List all miner addresses per group
*MinersApi* | [**PostMinersCpuMining**](docs/MinersApi.md#postminerscpumining) | **Post** /miners/cpu-mining | Execute an action on CPU miner. !!! for test only !!!
*MinersApi* | [**PostMinersCpuMiningMineOneBlock**](docs/MinersApi.md#postminerscpuminingmineoneblock) | **Post** /miners/cpu-mining/mine-one-block | Mine a block on CPU miner. !!! for test only !!!
*MinersApi* | [**PostWalletsWalletNameDeriveNextMinerAddresses**](docs/MinersApi.md#postwalletswalletnamederivenextmineraddresses) | **Post** /wallets/{wallet_name}/derive-next-miner-addresses | Derive your next miner addresses for each group
*MinersApi* | [**PutMinersAddresses**](docs/MinersApi.md#putminersaddresses) | **Put** /miners/addresses | Update miner&#39;s addresses, but better to use user.conf instead
*MultiSignatureApi* | [**PostMultisigAddress**](docs/MultiSignatureApi.md#postmultisigaddress) | **Post** /multisig/address | Create the multisig address and unlock script
*MultiSignatureApi* | [**PostMultisigBuild**](docs/MultiSignatureApi.md#postmultisigbuild) | **Post** /multisig/build | Build a multisig unsigned transaction
*MultiSignatureApi* | [**PostMultisigSubmit**](docs/MultiSignatureApi.md#postmultisigsubmit) | **Post** /multisig/submit | Submit a multi-signed transaction
*TransactionsApi* | [**GetTransactionsStatus**](docs/TransactionsApi.md#gettransactionsstatus) | **Get** /transactions/status | Get tx status
*TransactionsApi* | [**GetTransactionsUnconfirmed**](docs/TransactionsApi.md#gettransactionsunconfirmed) | **Get** /transactions/unconfirmed | List unconfirmed transactions
*TransactionsApi* | [**PostTransactionsBuild**](docs/TransactionsApi.md#posttransactionsbuild) | **Post** /transactions/build | Build an unsigned transaction to a number of recipients
*TransactionsApi* | [**PostTransactionsDecodeUnsignedTx**](docs/TransactionsApi.md#posttransactionsdecodeunsignedtx) | **Post** /transactions/decode-unsigned-tx | Decode an unsigned transaction
*TransactionsApi* | [**PostTransactionsSubmit**](docs/TransactionsApi.md#posttransactionssubmit) | **Post** /transactions/submit | Submit a signed transaction
*TransactionsApi* | [**PostTransactionsSweepAddressBuild**](docs/TransactionsApi.md#posttransactionssweepaddressbuild) | **Post** /transactions/sweep-address/build | Build unsigned transactions to send all unlocked balanced of one address to another address
*UtilsApi* | [**PostUtilsVerifySignature**](docs/UtilsApi.md#postutilsverifysignature) | **Post** /utils/verify-signature | Verify the SecP256K1 signature of some data
*UtilsApi* | [**PutUtilsCheckHashIndexing**](docs/UtilsApi.md#pututilscheckhashindexing) | **Put** /utils/check-hash-indexing | Check and repair the indexing of block hashes
*WalletsApi* | [**DeleteWalletsWalletName**](docs/WalletsApi.md#deletewalletswalletname) | **Delete** /wallets/{wallet_name} | Delete your wallet file (can be recovered with your mnemonic)
*WalletsApi* | [**GetWallets**](docs/WalletsApi.md#getwallets) | **Get** /wallets | List available wallets
*WalletsApi* | [**GetWalletsWalletName**](docs/WalletsApi.md#getwalletswalletname) | **Get** /wallets/{wallet_name} | Get wallet&#39;s status
*WalletsApi* | [**GetWalletsWalletNameAddresses**](docs/WalletsApi.md#getwalletswalletnameaddresses) | **Get** /wallets/{wallet_name}/addresses | List all your wallet&#39;s addresses
*WalletsApi* | [**GetWalletsWalletNameAddressesAddress**](docs/WalletsApi.md#getwalletswalletnameaddressesaddress) | **Get** /wallets/{wallet_name}/addresses/{address} | Get address&#39; info
*WalletsApi* | [**GetWalletsWalletNameBalances**](docs/WalletsApi.md#getwalletswalletnamebalances) | **Get** /wallets/{wallet_name}/balances | Get your total balance
*WalletsApi* | [**PostWallets**](docs/WalletsApi.md#postwallets) | **Post** /wallets | Create a new wallet
*WalletsApi* | [**PostWalletsWalletNameChangeActiveAddress**](docs/WalletsApi.md#postwalletswalletnamechangeactiveaddress) | **Post** /wallets/{wallet_name}/change-active-address | Choose the active address
*WalletsApi* | [**PostWalletsWalletNameDeriveNextAddress**](docs/WalletsApi.md#postwalletswalletnamederivenextaddress) | **Post** /wallets/{wallet_name}/derive-next-address | Derive your next address
*WalletsApi* | [**PostWalletsWalletNameLock**](docs/WalletsApi.md#postwalletswalletnamelock) | **Post** /wallets/{wallet_name}/lock | Lock your wallet
*WalletsApi* | [**PostWalletsWalletNameRevealMnemonic**](docs/WalletsApi.md#postwalletswalletnamerevealmnemonic) | **Post** /wallets/{wallet_name}/reveal-mnemonic | Reveal your mnemonic. !!! use it with caution !!!
*WalletsApi* | [**PostWalletsWalletNameSign**](docs/WalletsApi.md#postwalletswalletnamesign) | **Post** /wallets/{wallet_name}/sign | Sign the given data and return back the signature
*WalletsApi* | [**PostWalletsWalletNameSweepActiveAddress**](docs/WalletsApi.md#postwalletswalletnamesweepactiveaddress) | **Post** /wallets/{wallet_name}/sweep-active-address | Transfer all unlocked ALPH from the active address to another address
*WalletsApi* | [**PostWalletsWalletNameSweepAllAddresses**](docs/WalletsApi.md#postwalletswalletnamesweepalladdresses) | **Post** /wallets/{wallet_name}/sweep-all-addresses | Transfer unlocked ALPH from all addresses (including all mining addresses if applicable) to another address
*WalletsApi* | [**PostWalletsWalletNameTransfer**](docs/WalletsApi.md#postwalletswalletnametransfer) | **Post** /wallets/{wallet_name}/transfer | Transfer ALPH from the active address
*WalletsApi* | [**PostWalletsWalletNameUnlock**](docs/WalletsApi.md#postwalletswalletnameunlock) | **Post** /wallets/{wallet_name}/unlock | Unlock your wallet
*WalletsApi* | [**PutWallets**](docs/WalletsApi.md#putwallets) | **Put** /wallets | Restore a wallet from your mnemonic


## Documentation For Models

 - [AddressBalance](docs/AddressBalance.md)
 - [AddressInfo](docs/AddressInfo.md)
 - [Addresses](docs/Addresses.md)
 - [AssetInput](docs/AssetInput.md)
 - [AssetOutput](docs/AssetOutput.md)
 - [AssetState](docs/AssetState.md)
 - [BadRequest](docs/BadRequest.md)
 - [Balance](docs/Balance.md)
 - [Balances](docs/Balances.md)
 - [Ban](docs/Ban.md)
 - [Banned](docs/Banned.md)
 - [BlockAndEvents](docs/BlockAndEvents.md)
 - [BlockEntry](docs/BlockEntry.md)
 - [BlockHeaderEntry](docs/BlockHeaderEntry.md)
 - [BlocksAndEventsPerTimeStampRange](docs/BlocksAndEventsPerTimeStampRange.md)
 - [BlocksPerTimeStampRange](docs/BlocksPerTimeStampRange.md)
 - [BrokerInfo](docs/BrokerInfo.md)
 - [BrokerInfoAddress](docs/BrokerInfoAddress.md)
 - [BuildDeployContractTx](docs/BuildDeployContractTx.md)
 - [BuildDeployContractTxResult](docs/BuildDeployContractTxResult.md)
 - [BuildExecuteScriptTx](docs/BuildExecuteScriptTx.md)
 - [BuildExecuteScriptTxResult](docs/BuildExecuteScriptTxResult.md)
 - [BuildInfo](docs/BuildInfo.md)
 - [BuildMultisig](docs/BuildMultisig.md)
 - [BuildMultisigAddress](docs/BuildMultisigAddress.md)
 - [BuildMultisigAddressResult](docs/BuildMultisigAddressResult.md)
 - [BuildSweepAddressTransactions](docs/BuildSweepAddressTransactions.md)
 - [BuildSweepAddressTransactionsResult](docs/BuildSweepAddressTransactionsResult.md)
 - [BuildTransaction](docs/BuildTransaction.md)
 - [BuildTransactionResult](docs/BuildTransactionResult.md)
 - [CallContract](docs/CallContract.md)
 - [CallContractResult](docs/CallContractResult.md)
 - [ChainInfo](docs/ChainInfo.md)
 - [ChainParams](docs/ChainParams.md)
 - [ChangeActiveAddress](docs/ChangeActiveAddress.md)
 - [CompileContractResult](docs/CompileContractResult.md)
 - [CompileProjectResult](docs/CompileProjectResult.md)
 - [CompileScriptResult](docs/CompileScriptResult.md)
 - [CompilerOptions](docs/CompilerOptions.md)
 - [Confirmed](docs/Confirmed.md)
 - [Contract](docs/Contract.md)
 - [ContractEvent](docs/ContractEvent.md)
 - [ContractEventByBlockHash](docs/ContractEventByBlockHash.md)
 - [ContractEventByTxId](docs/ContractEventByTxId.md)
 - [ContractEvents](docs/ContractEvents.md)
 - [ContractEventsByBlockHash](docs/ContractEventsByBlockHash.md)
 - [ContractEventsByTxId](docs/ContractEventsByTxId.md)
 - [ContractOutput](docs/ContractOutput.md)
 - [ContractState](docs/ContractState.md)
 - [DecodeUnsignedTx](docs/DecodeUnsignedTx.md)
 - [DecodeUnsignedTxResult](docs/DecodeUnsignedTxResult.md)
 - [Destination](docs/Destination.md)
 - [DiscoveryAction](docs/DiscoveryAction.md)
 - [EventSig](docs/EventSig.md)
 - [FieldsSig](docs/FieldsSig.md)
 - [FixedAssetOutput](docs/FixedAssetOutput.md)
 - [FunctionSig](docs/FunctionSig.md)
 - [Group](docs/Group.md)
 - [HashesAtHeight](docs/HashesAtHeight.md)
 - [InterCliquePeerInfo](docs/InterCliquePeerInfo.md)
 - [InternalServerError](docs/InternalServerError.md)
 - [MemPooled](docs/MemPooled.md)
 - [MinerAddresses](docs/MinerAddresses.md)
 - [MinerAddressesInfo](docs/MinerAddressesInfo.md)
 - [MisbehaviorAction](docs/MisbehaviorAction.md)
 - [NodeInfo](docs/NodeInfo.md)
 - [NodeVersion](docs/NodeVersion.md)
 - [NotFound](docs/NotFound.md)
 - [Output](docs/Output.md)
 - [OutputRef](docs/OutputRef.md)
 - [PeerAddress](docs/PeerAddress.md)
 - [PeerMisbehavior](docs/PeerMisbehavior.md)
 - [PeerStatus](docs/PeerStatus.md)
 - [Penalty](docs/Penalty.md)
 - [Project](docs/Project.md)
 - [Reachable](docs/Reachable.md)
 - [ReleaseVersion](docs/ReleaseVersion.md)
 - [RevealMnemonic](docs/RevealMnemonic.md)
 - [RevealMnemonicResult](docs/RevealMnemonicResult.md)
 - [Script](docs/Script.md)
 - [SelfClique](docs/SelfClique.md)
 - [ServiceUnavailable](docs/ServiceUnavailable.md)
 - [Sign](docs/Sign.md)
 - [SignResult](docs/SignResult.md)
 - [SubmitMultisig](docs/SubmitMultisig.md)
 - [SubmitTransaction](docs/SubmitTransaction.md)
 - [SubmitTxResult](docs/SubmitTxResult.md)
 - [Sweep](docs/Sweep.md)
 - [SweepAddressTransaction](docs/SweepAddressTransaction.md)
 - [TestContract](docs/TestContract.md)
 - [TestContractResult](docs/TestContractResult.md)
 - [TestInputAsset](docs/TestInputAsset.md)
 - [Token](docs/Token.md)
 - [Transaction](docs/Transaction.md)
 - [TransactionTemplate](docs/TransactionTemplate.md)
 - [Transfer](docs/Transfer.md)
 - [TransferResult](docs/TransferResult.md)
 - [TransferResults](docs/TransferResults.md)
 - [TxNotFound](docs/TxNotFound.md)
 - [TxStatus](docs/TxStatus.md)
 - [UTXO](docs/UTXO.md)
 - [UTXOs](docs/UTXOs.md)
 - [Unauthorized](docs/Unauthorized.md)
 - [Unban](docs/Unban.md)
 - [UnconfirmedTransactions](docs/UnconfirmedTransactions.md)
 - [Unreachable](docs/Unreachable.md)
 - [UnsignedTx](docs/UnsignedTx.md)
 - [Val](docs/Val.md)
 - [ValAddress](docs/ValAddress.md)
 - [ValArray](docs/ValArray.md)
 - [ValBool](docs/ValBool.md)
 - [ValByteVec](docs/ValByteVec.md)
 - [ValI256](docs/ValI256.md)
 - [ValU256](docs/ValU256.md)
 - [VerifySignature](docs/VerifySignature.md)
 - [WalletCreation](docs/WalletCreation.md)
 - [WalletCreationResult](docs/WalletCreationResult.md)
 - [WalletDeletion](docs/WalletDeletion.md)
 - [WalletRestore](docs/WalletRestore.md)
 - [WalletRestoreResult](docs/WalletRestoreResult.md)
 - [WalletStatus](docs/WalletStatus.md)
 - [WalletUnlock](docs/WalletUnlock.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



