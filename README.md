## How to start

Start the service server to serve HTTP requests and a worker to synchronize blocks from Ronin.

```bash
./scripts/bin.sh service start
```

How worker works

```text
The worker will run at a configured interval, so at each interval the worker will call Ronin to get the next block.
For first block, the worker calls Ronin to get the most recent block and store it
For next block, the workers calls Ronin to get next block using the number of the last synced block.
```

Eviction strategy

```text
To avoid a memory leak, I have a configuration for maximum capacity. After reaching maximum capacity, the heading block is poped.
```

Data structure

```text
I used a doubly linked list to store blocks. It is convenient for implementing an eviction strategy.
```

## Worker synchronizes with Ronin

![Diagram](docs/image/worker.drawio.png)

## API Transaction by hash

![Diagram](docs/image/getTransactionByHash.drawio.png)

```bash
curl --location 'localhost:8090/v1/transactions?hash=0x500e83afa0f4e85ff031a0e245cae12ca212b27720d63f5201ea7cb45d65e0ef'
```

## API List transactions by block number

![Diagram](docs/image/getTransactionsByBlockNumber.drawio.png)

```bash
curl --location 'localhost:8090/v1/blocks/transactions?block_number=0x2132959'
```

## API list transaction in range

![Diagram](docs/image/getTransactionsInRange.drawio.png)

```bash
curl --location 'localhost:8090/v1/transactions/list?from=0&to=9'
```

## API get percentage of transactions which have gas fee lees than

![Diagram](docs/image/getPercentageOfTransactionGasFee.drawio.png)

```bash
curl --location 'localhost:8090/v1/transactions/gas?gas_fee=10'
```

