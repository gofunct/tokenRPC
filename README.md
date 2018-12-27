# tokenRPC

**A gRPC Powered Blockchain Microservice**

**Author: Coleman Word**


# Usage
```commandline
Usage:
  token [command]

Available Commands:
  client      A brief description of your command
  gateway     A brief description of your command
  help        Help about any command
  proxy       A brief description of your command

Flags:
  -c, --client-addr string     the ethereum client address (default "http://192.168.99.100:8545")
  -t, --contract-addr string   contract address
  -h, --help                   help for token
  -k, --key string             deployer's private key
  -s, --proxy-addr string      server address to dial (default "127.0.0.1:5555")

Use "token [command] --help" for more information about a command.

```

# Protocol Documentation

## Table of Contents

<div id="toc-container">

*   [token_service/token_service.proto](#token_service%2ftoken_service.proto)
    *   [<span class="badge">M</span>AllowanceReq](#token_service.AllowanceReq)
    *   [<span class="badge">M</span>AllowanceResp](#token_service.AllowanceResp)
    *   [<span class="badge">M</span>ApproveAndCallReq](#token_service.ApproveAndCallReq)
    *   [<span class="badge">M</span>ApproveAndCallResp](#token_service.ApproveAndCallResp)
    *   [<span class="badge">M</span>ApproveReq](#token_service.ApproveReq)
    *   [<span class="badge">M</span>ApproveResp](#token_service.ApproveResp)
    *   [<span class="badge">M</span>BalanceOfReq](#token_service.BalanceOfReq)
    *   [<span class="badge">M</span>BalanceOfResp](#token_service.BalanceOfResp)
    *   [<span class="badge">M</span>BurnFromReq](#token_service.BurnFromReq)
    *   [<span class="badge">M</span>BurnFromResp](#token_service.BurnFromResp)
    *   [<span class="badge">M</span>BurnReq](#token_service.BurnReq)
    *   [<span class="badge">M</span>BurnResp](#token_service.BurnResp)
    *   [<span class="badge">M</span>DecimalsResp](#token_service.DecimalsResp)
    *   [<span class="badge">M</span>Empty](#token_service.Empty)
    *   [<span class="badge">M</span>NameResp](#token_service.NameResp)
    *   [<span class="badge">M</span>OnApprovalReq](#token_service.OnApprovalReq)
    *   [<span class="badge">M</span>OnBurnReq](#token_service.OnBurnReq)
    *   [<span class="badge">M</span>OnTransferReq](#token_service.OnTransferReq)
    *   [<span class="badge">M</span>SymbolResp](#token_service.SymbolResp)
    *   [<span class="badge">M</span>TotalSupplyResp](#token_service.TotalSupplyResp)
    *   [<span class="badge">M</span>TransactOpts](#token_service.TransactOpts)
    *   [<span class="badge">M</span>TransactionReq](#token_service.TransactionReq)
    *   [<span class="badge">M</span>TransactionResp](#token_service.TransactionResp)
    *   [<span class="badge">M</span>TransferFromReq](#token_service.TransferFromReq)
    *   [<span class="badge">M</span>TransferFromResp](#token_service.TransferFromResp)
    *   [<span class="badge">M</span>TransferReq](#token_service.TransferReq)
    *   [<span class="badge">M</span>TransferResp](#token_service.TransferResp)
    *   [<span class="badge">S</span>Token](#token_service.Token)
*   [Scalar Value Types](#scalar-value-types)

</div>

<div class="file-heading">

## token_service/token_service.proto

[Top](#title)</div>

### AllowanceReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>arg2</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### AllowanceResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### ApproveAndCallReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>spender</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>extra_data</td>

<td>[bytes](#bytes)</td>

<td>repeated</td>

<td></td>

</tr>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### ApproveAndCallResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>success</td>

<td>[bool](#bool)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### ApproveReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>spender</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### ApproveResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>success</td>

<td>[bool](#bool)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### BalanceOfReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### BalanceOfResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### BurnFromReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>from</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### BurnFromResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>success</td>

<td>[bool](#bool)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### BurnReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### BurnResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>success</td>

<td>[bool](#bool)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### DecimalsResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### Empty

### NameResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### OnApprovalReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>owner</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>spender</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### OnBurnReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>from</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### OnTransferReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>from</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>to</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### SymbolResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TotalSupplyResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>arg</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransactOpts

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>from_address</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>private_key</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>nonce</td>

<td>[int64](#int64)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[int64](#int64)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>gas_price</td>

<td>[int64](#int64)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>gas_limit</td>

<td>[int64](#int64)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransactionReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransactionResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>tx_hash</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransferFromReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>from</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>to</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransferFromResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>success</td>

<td>[bool](#bool)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransferReq

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>to</td>

<td>[string](#string)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>value</td>

<td>[bytes](#bytes)</td>

<td></td>

<td></td>

</tr>

<tr>

<td>opts</td>

<td>[TransactOpts](#token_service.TransactOpts)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### TransferResp

<table class="field-table">

<thead>

<tr>

<td>Field</td>

<td>Type</td>

<td>Label</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>success</td>

<td>[bool](#bool)</td>

<td></td>

<td></td>

</tr>

</tbody>

</table>

### Token

<table class="enum-table">

<thead>

<tr>

<td>Method Name</td>

<td>Request Type</td>

<td>Response Type</td>

<td>Description</td>

</tr>

</thead>

<tbody>

<tr>

<td>allowance</td>

<td>[AllowanceReq](#token_service.AllowanceReq)</td>

<td>[AllowanceResp](#token_service.AllowanceResp)</td>

<td></td>

</tr>

<tr>

<td>approve</td>

<td>[ApproveReq](#token_service.ApproveReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>approveAndCall</td>

<td>[ApproveAndCallReq](#token_service.ApproveAndCallReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>balanceOf</td>

<td>[BalanceOfReq](#token_service.BalanceOfReq)</td>

<td>[BalanceOfResp](#token_service.BalanceOfResp)</td>

<td></td>

</tr>

<tr>

<td>burn</td>

<td>[BurnReq](#token_service.BurnReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>burnFrom</td>

<td>[BurnFromReq](#token_service.BurnFromReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>decimals</td>

<td>[Empty](#token_service.Empty)</td>

<td>[DecimalsResp](#token_service.DecimalsResp)</td>

<td></td>

</tr>

<tr>

<td>name</td>

<td>[Empty](#token_service.Empty)</td>

<td>[NameResp](#token_service.NameResp)</td>

<td></td>

</tr>

<tr>

<td>symbol</td>

<td>[Empty](#token_service.Empty)</td>

<td>[SymbolResp](#token_service.SymbolResp)</td>

<td></td>

</tr>

<tr>

<td>totalSupply</td>

<td>[Empty](#token_service.Empty)</td>

<td>[TotalSupplyResp](#token_service.TotalSupplyResp)</td>

<td></td>

</tr>

<tr>

<td>transfer</td>

<td>[TransferReq](#token_service.TransferReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>transferFrom</td>

<td>[TransferFromReq](#token_service.TransferFromReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>onApproval</td>

<td>[OnApprovalReq](#token_service.OnApprovalReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>onBurn</td>

<td>[OnBurnReq](#token_service.OnBurnReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

<tr>

<td>onTransfer</td>

<td>[OnTransferReq](#token_service.OnTransferReq)</td>

<td>[TransactionResp](#token_service.TransactionResp)</td>

<td></td>

</tr>

</tbody>

</table>

## Scalar Value Types

<table class="scalar-value-types-table">

<thead>

<tr>

<td>.proto Type</td>

<td>Notes</td>

<td>C++ Type</td>

<td>Java Type</td>

<td>Python Type</td>

</tr>

</thead>

<tbody>

<tr id="double">

<td>double</td>

<td></td>

<td>double</td>

<td>double</td>

<td>float</td>

</tr>

<tr id="float">

<td>float</td>

<td></td>

<td>float</td>

<td>float</td>

<td>float</td>

</tr>

<tr id="int32">

<td>int32</td>

<td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead.</td>

<td>int32</td>

<td>int</td>

<td>int</td>

</tr>

<tr id="int64">

<td>int64</td>

<td>Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead.</td>

<td>int64</td>

<td>long</td>

<td>int/long</td>

</tr>

<tr id="uint32">

<td>uint32</td>

<td>Uses variable-length encoding.</td>

<td>uint32</td>

<td>int</td>

<td>int/long</td>

</tr>

<tr id="uint64">

<td>uint64</td>

<td>Uses variable-length encoding.</td>

<td>uint64</td>

<td>long</td>

<td>int/long</td>

</tr>

<tr id="sint32">

<td>sint32</td>

<td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.</td>

<td>int32</td>

<td>int</td>

<td>int</td>

</tr>

<tr id="sint64">

<td>sint64</td>

<td>Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.</td>

<td>int64</td>

<td>long</td>

<td>int/long</td>

</tr>

<tr id="fixed32">

<td>fixed32</td>

<td>Always four bytes. More efficient than uint32 if values are often greater than 2^28.</td>

<td>uint32</td>

<td>int</td>

<td>int</td>

</tr>

<tr id="fixed64">

<td>fixed64</td>

<td>Always eight bytes. More efficient than uint64 if values are often greater than 2^56.</td>

<td>uint64</td>

<td>long</td>

<td>int/long</td>

</tr>

<tr id="sfixed32">

<td>sfixed32</td>

<td>Always four bytes.</td>

<td>int32</td>

<td>int</td>

<td>int</td>

</tr>

<tr id="sfixed64">

<td>sfixed64</td>

<td>Always eight bytes.</td>

<td>int64</td>

<td>long</td>

<td>int/long</td>

</tr>

<tr id="bool">

<td>bool</td>

<td></td>

<td>bool</td>

<td>boolean</td>

<td>boolean</td>

</tr>

<tr id="string">

<td>string</td>

<td>A string must always contain UTF-8 encoded or 7-bit ASCII text.</td>

<td>string</td>

<td>String</td>

<td>str/unicode</td>

</tr>

<tr id="bytes">

<td>bytes</td>

<td>May contain any arbitrary sequence of bytes.</td>

<td>string</td>

<td>ByteString</td>

<td>str</td>

</tr>

</tbody>

</table> 