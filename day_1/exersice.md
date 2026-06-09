# Exersice 1 - Check your node 
- Check blockchain info

[command]
bitcoin-cli -regtest getblockchaininfo

[output]
{
  "chain": "regtest",
  "blocks": 104,
  "headers": 104,
  "bestblockhash": "2fbdbd823581c10c0d6cc8806f04e616d7e8c56c418992d96a165b3a4395d7e9",
  "bits": "207fffff",
  "target": "7fffff0000000000000000000000000000000000000000000000000000000000",
  "difficulty": 4.656542373906925e-10,
  "time": 1780934099,
  "mediantime": 1780924863,
  "verificationprogress": 0.5898664255748441,
  "initialblockdownload": false,
  "chainwork": "00000000000000000000000000000000000000000000000000000000000000d2",
  "size_on_disk": 31772,
  "pruned": false,
  "warnings": [
  ]
}
- check network info
[command]
bitcoin-cli -regtest getnetworkinfo

[output]
{
  "version": 310000,
  "subversion": "/Satoshi:31.0.0/",
  "protocolversion": 70016,
  "localservices": "0000000000000c09",
  "localservicesnames": [
    "NETWORK",
    "WITNESS",
    "NETWORK_LIMITED",
    "P2P_V2"
  ],
  "localrelay": true,
  "timeoffset": 0,
  "networkactive": true,
  "connections": 0,
  "connections_in": 0,
  "connections_out": 0,
  "networks": [
    {
      "name": "ipv4",
      "limited": false,
      "reachable": true,
      "proxy": "",
      "proxy_randomize_credentials": false
    },
    {
      "name": "ipv6",
      "limited": false,
      "reachable": true,
      "proxy": "",
      "proxy_randomize_credentials": false
    },
    {
      "name": "onion",
      "limited": true,
      "reachable": false,
      "proxy": "",
      "proxy_randomize_credentials": false
    },
    {
      "name": "i2p",
      "limited": true,
      "reachable": false,
      "proxy": "",
      "proxy_randomize_credentials": false
    },
    {
      "name": "cjdns",
      "limited": true,
      "reachable": false,
      "proxy": "",
      "proxy_randomize_credentials": false
    }
  ],
  "relayfee": 0.00000100,
  "incrementalfee": 0.00000100,
  "localaddresses": [
  ],
  "warnings": [
  ]
}

# Exersice 2 - Create wallets
[command]
bitcoin-cli -regtest createwallet "alice"
bitcoin-cli -regtest createwallet "bob"

# Exersice 3 - Generate Addresses
bitcoin-cli -regtest -rpcwallet=bob getnewaddress
bcrt1qelm9x7m62yspltsr0s0t04da4avv6ydxx05jsx




