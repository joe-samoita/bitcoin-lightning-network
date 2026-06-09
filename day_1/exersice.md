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
[command]
bitcoin-cli -regtest -rpcwallet=bob getnewaddress

[output]
bcrt1qelm9x7m62yspltsr0s0t04da4avv6ydxx05jsx

# Exersice 4 - Mine Blocks
[commands]
ALICE=$(bitcoin-cli -regtest -rpcwallet=alice getnewaddress)

bitcoin-cli -regtest generatetoaddress 101 "$ALICE"
[output]

[
  "1fd41fb38b80abdde34146f39d92acfb81f281ea3f1ab284cf22af1f29a61354",
  "7b768c083c8327a1c27ced3397bf070f8c64fc80115053256d561aecf173a5d0",
  "7087fea4ad277a3ad0a8ce29007fb7a996f856b0a63c170162f5104b42d2a2ab",
  "37afe5e94395c487cb5cadde7869f21a56afdede6d439ca8a5e3d691a3cc0e52",
  "474f1ce081db277ceac237fca8244705cb75d21a09025666fe20074384ba8b81",
  "0bc854b5d1829f7ea373717eace8ab785287edf58adb0b7e30ddff7f6ebfe71e",
  "54e2593687af8758781ebeebc4ced9c83c060d485bfdad98a2183e76cb4a8b3d",
  "33daf0ad35cfb67cde0fb321e9e7075d804692f7220534b54079fcd4167a4d27",
  "2119cf58455e34a88e56617f9f2ce2b1c6054cdebb4196cbfb1e6b1f7c87c992",
  "0cc42d237045cee9af786c5da8a5f669b24f3c471ea82bdce022d70d1b1d30a3",
  "09773bbafd37edf7cf15986f17b5331bcf4532698866698c6e5b93588a95f3f2",
  "2204f4387dac27a84b5d5c7edf81d5afba83fbb4763e75e53be01b78f5c43f0b",
  "75abefb12229ccdade20e387057f0fdb801ce8ddb372d91f0adc0a0c247d1a1f",
  "208a3c33d2ed836f663e967ff6389f30453303a037c050f3925358e40b8444d4",
  "5412b92ffdb6d30cf517dc1131f7a4e0a1f57ad0df70ae5293d2fa1591442d82",
  "18748fa4281d6a82cf0f37e413758bde2e86a72c57a06df8c9591acd314b5693",
  "1204b065279753706ed17f7579875559ba7058efcad5605c3719bd06bac9374c",
  "3ab124e45ef7fb9fa5897cb0e0f2cda295f7b70dd5523d8d2302259ebc3fe383",
  "4b3990d1bafed1536d84d08c1cebbf41db44f7333f2c2397f1c9043ef8754d4c",
  "5d5006c6492ab750e9678de05c6a249b6c27e458e71b576b2d52489f9fda1fff",
  "27c28827534a2cfc95f14b0db0a388c67802ccc31577ad4776725e07c531fe8a",
  "3aadce699a8d9db8be1d44e25b3b1a6ce9638d55cf49797f204ea3df5d66291b",
  "46ded70d72f2cc253d098d31935a3b53110147570464ea6101efeee19faa71b2",
  "5df76fe07b0a7133e606fa3b0b04e62a273519882f882c493552e97ab54c7a61",
  "6866be6a31a51315103aad4e9e6429104b134d44201fb54c831bfc13da4339e1",
  "081a479fc38ffb6c731d4716430e9d9eea94e0c2cdf11d9036963cd58bad305e",
  "5417aa128a19a43e8f8bbf63092f94ca84e81d1c9515fd33f165a4894db27a1e",
  "7993a70d8a3895a7971836f70c8a35fbfdbb490981ea6d167fe6d18782411627",
  "1fa47c2030916625c44cbf2f10e123445c709d0a987983a5a8b81767eebab5f1",
  "636529dbd114c665370c9b88aea5b54bb2ec7599e715fdb2775448c7a51947ea",
  "0707905b54c48989707c2b1b528749d4db4da36523743aaab89b3c22e1833eff",
  "666473d09d29ed90c3ead70a15cf747171d6d25cb550f242f5a332c9933bdd78",
  "2de7bf4e8febe4846336f602c87b6ec90ba023c92903129dc431de44e38c8d0a",
  "358392d6f148f4b2017a53d3d9dc7f1ef35e9ef16d7b5530eba23f05e1d42814",
  "73503d5a79c885fedc61b673585f1ecfc5db964cc5d61e1d53fc433bc61b62a2",
  "48bd5a5d9b2b35c0898762100f2fddf98d6bacedc0f5e16dde2f59c9c9affe32",
  "6edda4856ce8e99ae06d448d278941556a677f9089a3422c851be2545385bd56",
  "6199a876ccdf421b435395a0a967f5276160af89f9e32bb6fa1792a87fca57e4",
  "68878bc81c8e111558b690270ade28337680dfd61ba0b42417b87253777da435",
  "5e8313b592b1dbb1b028f830a2c34255e30853ae2ccf9592e5c01f27cb1763a2",
  "2b9137d49f93148a5351691132d2eb69a0ad2a9c5bf88cd5664d3af9140d4eeb",
  "100481cd4e9f11a4767fe2facc5d9f1ade8aeb15bf4e8782d84ec8ab9492703b",
  "3cb1e32c67b744123cc062785599ec91c9a753dba2ef53c4abe4f61f7b507484",
  "7a1db5ea2899284aef76785e9a57fa51b6d1529ceef95586bb7a106e5e4a98d9",
  "03b6d4f7c700359209cb8356a0495ca800623144e5baaaf60b81b2a571a89a7f",
  "039b435155f59528a28ee6cb008f893639a9c328098f5c5f1a862f4f72ea3a0b",
  "2ed816e52851e5519dcab9eab7710806bcbd2d95ecf525a0f44d736f7bab4b35",
  "4254ff361ea2a10417e32e7979a5e4a90c4e3e19e1b6de296bb7ac5823752faa",
  "317a8bca8c8c66d774824193fde5d9b314a52a9e59426be76ab89e281d06e0e1",
  "7dddd60238ae58bc902dd85d72ea38052935e70b5879ae49022848c2a979bb94",
  "4a3c947a7726a11ca7faf2e3d85c1943edea6b06a6d19d54f8b5525987dad308",
  "6f73381fd40278490ca60c9f9ab536e4eb262819c5c8f47de563ae0bb185464e",
  "75536a7f396299eda8d301ae6db92339ca3cf11afa5f0ebc65611ce7b047e3f0",
  "2f5522ea63bd402ca3365f154e229b5913132ea3a22663fca8de8e90c3dd637e",
  "2f34f8fef0f45ce38dbea96e5b6c31ec2045d30250186cf51dc5d4f4c26619eb",
  "4b398274d906d1c6a4948a53bb52de39978a7da45247d4f0e1950eb4f174424d",
  "12819188a9f6afadc4fc22da9a247bd873e9a5dc0207f1c16bcf0f733d7126ee",
  "260a1746172a43ce51f0a24361a017871fb649d4a189c021facd671af43ef3ba",
  "269b3e877b3079de22723aa08bd55f4ff35a2b54f2e45f7bb95eb44c5fb40f62",
  "233008e6626e6299e42f06973964075fc608247020aca8eae9bee8b6fd63b97e",
  "52636d90124ea9019a06348a2d648a147ad9c212e9dfc3f3be58ff9b87a8b97e",
  "7d4920d6dddf9f002f05c4803afc3f59469a056e6a2f8b42782896b2ced16907",
  "54086b2d2ff1c64d9e3d6fcaf3f75cd1148fa9216aa9e4156cb3a952a04d8bfc",
  "319bfe2381b67e663bf02a64b4c4dfa26ebf25fffefd19675e790851df3561ed",
  "4fbd5bec43c9b0479608012eb7eb1394aced154d5ae50d7e40948a051666f6ab",
  "2348e94d1e18a5debd45785d33052465ec22fd1c315013362003bae5a9ebd666",
  "1eba038456a1372389cd26906c01c1ed721d18174555a8a439cafdde979abc74",
  "6df5c83623ae76835841d0df2d37a955e86c696308fd8e111428a8e1034a78d6",
  "24e9a40160642c6667e5daddfedda10da86ffaedf58e66124e122a75c736c81f",
  "551c32f15e0d4463c214cdfbe55414a0e0f160df826cd0c7dd4d3974a9f590f4",
  "40563b1edafae457c118d0533f36354bd9a1a3fe29ac8e4d9a7ca569418a2e4b",
  "2ca99b046c3fbf3e51028d05178530aa8603c317895f8e5b78ef7b9cba581cbe",
  "2cfe73f98cfba088424e37ea4db47396d3e6d3cb22b2df2432cc04bceb281d6f",
  "2b4d7f38da96b5bfbd34643b6ded7220685e3bb07e04b5492af2a5c104de55df",
  "0638469aa6a2f8640aa0511b32d76e8de3c57ad4ff2aa85c8a23418df99c6a93",
  "1ed14697b5831023d90fc48190aa76649bcb0b215c12ac9402bb0531c5f94dec",
  "1e4286f6c8cfdababdb88b3cc34b48a0a4efeecaca0b80408ab4e16827b31a70",
  "56750ac125a3dc2545cfcb0b5490a397fb78a78b9b256416e1f4c196f687db17",
  "033742f774ce6caaaba763df9cc46843d7823e54a7c547183192ff1bf95fb805",
  "6cbfd7ffd962331ebe2fbb8249b22d170422c8a8504d8e50e28b7e899884e498",
  "3cf04156851aef94e8ef916861b9903ca9d94308812ca00687d3d211c45efe1e",
  "251d85f805632a4335ad47828084016b73d589be20cf7c830cb809c28b75a258",
  "32dfe0019496734f0d3482dfb4f0de3b5cc7d4ad2d49e9b32c36181aaa789f08",
  "4bc7fa02a551479e1ac702b00db98834a015d283a8125350fdf139714b3ca5b4",
  "5da2c572b10f9400d6cd31e3cae88450ef59ba783613142dad23b57ea2183ad8",
  "084ca0457d867236d247204e8bbffe7230c17dfe52e6bd50ac1f7b464586a36f",
  "655881dcbf3b3db0deb1941bed47577a701f87c2838d44b87542fa89d42a88fa",
  "4ad1fe477ce6864412156794085bf87bb5ee02ec126219d58598a9deaa2b91fb",
  "7f900178aeeb89b541230bbf57d25eb35c1e5910d2a60bc9bff648122176fcd8",
  "0aa9467a77bbdd5bde157fa4ab20f637d4635fa45eb3a6e39d6d49d4f1da47a1",
  "69b5b3665081bf6f2a242ab717f224dd6fb1244ba974499eb32867211f98d6da",
  "07f6e56d589fb05e1399f94771e37f130ded9c8eb9c94613ce72ecd5d74d6725",
  "19537b501429b0f1832911069e79f7fb0f8d1c6058ca7ce30977e048c6b1cc6f",
  "6e9a9b3d5a7088555c8961d5731ffa723d99659823fa60e6a3f7aba1155013e3",
  "4e854ccc919e9cf3c93025cc6e8f95ba8e0291b679e3998337be809d06da012a",
  "27d172487513a7300795c8fb214f8d640315efdc0eaaea48dfb157377fc4cbf7",
  "39d10b5d517131145a594d209b92ba1be4fd4ccd1e8d5fbde59abcab1b8a2996",
  "3b2f0da4decf8d56c94d3fc06b0a9d669d2dcbd38f90ab1444f6cadca1d496d9",
  "04c8d77d8e2c4152db4ab95c59a22d50ab988a5f7a412e049354fefc3e8d7059",
  "13bb7a0c13519ed186cd9f33b2d28995e224db470ef8ab2ef40f41e8e2057180",
  "0f4302129a929eb809457e89cc9cbe85807ce476c76f41b40bbb2a2cad642fbc"
]


[command]
bitcoin-cli -regtest -rpcwallet=alice getbalance
[output]
5079.99997180 (Depends on what you have been workin on)

# Exersice 5 - Send transactions
- get Bob's address
[commands]
BOB=$(bitcoin-cli -regtest -rpcwallet=bob getnewaddress)

- send 10 btc from alice to bob
[command]
bitcoin-cli -regtest -rpcwallet=alice sendtoaddress "$BOB" 10

# confirm transaction

- check mempool
[commnd]
bitcoin-cli -regtest getmempoolinfo

[output]

{
  "loaded": true,
  "size": 0,
  "bytes": 0,
  "usage": 0,
  "total_fee": 0.00000000,
  "maxmempool": 300000000,
  "mempoolminfee": 0.00000100,
  "minrelaytxfee": 0.00000100,
  "incrementalrelayfee": 0.00000100,
  "unbroadcastcount": 0,
  "fullrbf": true,
  "permitbaremultisig": true,
  "maxdatacarriersize": 100000,
  "limitclustercount": 64,
  "limitclustersize": 101000,
  "optimal": true
}
# Exersice 6 - Explore transaction
- list alice's transactions

[command]
bitcoin-cli -regtest -rpcwallet=alice listtransactions

- getrawtransactions
bitcoin-cli -regtest getrawtransaction "TXID" true

[output]
{
  "txid": "89bd9f41651cbdb9d905981600387af50841d8ac6ee286dc87a756fde56fe8c5",
  "hash": "26ac23f216949d6a053e70c62e74d5acf5eb39cfc011013bcec181dbf0a6a62b",
  "version": 2,
  "size": 169,
  "vsize": 142,
  "weight": 568,
  "locktime": 204,
  "vin": [
    {
      "coinbase": "02cd0000",
      "txinwitness": [
        "0000000000000000000000000000000000000000000000000000000000000000"
      ],
      "sequence": 4294967294
    }
  ],
  "vout": [
    {
      "value": 25.00000000,
      "n": 0,
      "scriptPubKey": {
        "asm": "0 679f8c84d25354db8f63c3423ff2b0edb9829c5f",
        "desc": "addr(bcrt1qv70cepxj2d2dhrmrcdprlu4sakuc98zltd0wuf)#2nyhec0k",
        "hex": "0014679f8c84d25354db8f63c3423ff2b0edb9829c5f",
        "address": "bcrt1qv70cepxj2d2dhrmrcdprlu4sakuc98zltd0wuf",
        "type": "witness_v0_keyhash"
      }
    },
    {
      "value": 0.00000000,
      "n": 1,
      "scriptPubKey": {
        "asm": "OP_RETURN aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf9",
        "desc": "raw(6a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf9)#cav96mf3",
        "hex": "6a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf9",
        "type": "nulldata"
      }
    }
  ],
  "hex": "020000000001010000000000000000000000000000000000000000000000000000000000000000ffffffff0402cd0000feffffff0200f9029500000000160014679f8c84d25354db8f63c3423ff2b0edb9829c5f0000000000000000266a24aa21a9ede2f61c3f71d1defd3fa999dfa36953755c690689799962b48bebd836974e8cf901200000000000000000000000000000000000000000000000000000000000000000cc000000",
  "blockhash": "0f4302129a929eb809457e89cc9cbe85807ce476c76f41b40bbb2a2cad642fbc",
  "confirmations": 1,
  "time": 1781019914,
  "blocktime": 1781019914
}

# Exersice 7 - Explore a block
[command]
bitcoin-cli -regtest getbestblockhash
[output]
0f4302129a929eb809457e89cc9cbe85807ce476c76f41b40bbb2a2cad642fbc

[command]
bitcoin-cli -regtest getblock "0f4302129a929eb809457e89cc9cbe85807ce476c76f41b40bbb2a2cad642fbc"

# Exersice 8 - ADDRESS TYPES
- legacy(oldest)
bitcoin-cli -regtest -rpcwallet=bob getnewaddress "" "legacy"

- native(segwit32 recommended)
bitcoin-cli -regtest -rpcwallet=bob getnewaddress "" "bench32"

- taproot(newest)
bitcoin-cli -regtest -rpcwallet=bob getnewaddress "" "bench32m"







