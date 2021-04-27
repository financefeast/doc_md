from financefeast import FinanceFeast
from financefeast import Environments

client = FinanceFeast(client_id="your-client-id", client_secret="your-client-secret", environment=Environments.prod.value)
print(client.alive())