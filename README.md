Eigenlayer Popularity Datalayer#

Here is the flow for handling real time notifications for both on-chain and off-chain data. Distribution is done by Kafka channels in each step.

1-Starting point of the flow is obtaining the the transactions via blockchain scrapers where navigator is going to forward the transaction + transaction receipt.

2-When a transaction comes, the processor checks if the log contains OperatorRegistered event by parsing the logs. 

3-Then if the log is parsed without error, next step is getting metadata information from Dune Api. 

4-The data that is gathered from Dune Api, is sent to the smart contract in Movement Chain by calling registerEvent function.



SideNode: Due to the privacy of the project, the all code wasn't pushed. If it is requested, we can show it privately. 
