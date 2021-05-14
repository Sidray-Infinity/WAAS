import requests
import tqdm

url = "http://127.0.0.1:8080/wallet/balance/3"
params = {"amount":0.01, "type": True}

for i in tqdm.tqdm(range(9000000)):
    r = requests.patch(url=url, json=params)