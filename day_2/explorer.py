import requests, json

def rpc(method, params=None, wallet=None):
    url = "http://localhost:8332/"
    if wallet:
        wallet = f"{url}wallet/{wallet}/"
    data = json.dumps({
        "method": method,
        "params": params or [],
        "jsonrpc": "2.0",
        "id": "1"
    })
    response = requests.post(url, data=data, auth=('user', 'password'))
    return response.json()

info = rpc("getblockchaininfo")
print(f"Chain: {info['chain']}")
print(f"Blocks: {info['blocks']}")

balance = rpc("getbalance",[], wallet="alice")
print(f"Alice's balance: {balance} BTC")