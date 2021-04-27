from financefeast import FinanceFeast, Environments

client = FinanceFeast(client_id="your-client-id", client_secret="your-client-secret")
print(client.sma(ticker='air.nz'))