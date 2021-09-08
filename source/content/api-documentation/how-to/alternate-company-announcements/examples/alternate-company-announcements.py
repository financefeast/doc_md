from financefeast import Rest

client = Rest(token="API_TOKEN")
print(client.announcement(ticker="air.nz", data_from="2021-01-01").data)