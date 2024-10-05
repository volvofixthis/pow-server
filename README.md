### Description
This is an example of PoW.  
Client receives a challege and exchanges a solution on a passage.

### PoW algorithm selection
I liked the idea of memory-bound function as a pow algorithm. Argon2 has been selected because memory-hardness makes it more resistant to attackers using specialized hardware like GPUs, FPGAs, or ASICs, which are commonly used to accelerate brute-force attacks.

### Possible improvements
It will be nice to create UDP server. ConnAdapter is protocol agnostic already.  
I tried to create tcp server with limited number of workers, but it could be not enough for high load.  
It should be replaced with tcp server which will use event-loop for network.
Json should be replaced by protobuf or custom marshaling.

### Running in Containers
Test server and client should be launched with this command:  
```bash
docker-compose up
```
Ater introducing some modifications:  
```bash
docker-compose up --force-recreate --build
```

### Building locally
Server and client should be built with this command:  
```bash
make build
```

### Running test/linters
They should be run with this commands:  
```bash
make linters
make test
```

### Example of envs
```bash
# Bind address
export TCP_ADDRESS=127.0.0.1:8201
# Timeout of PoW task
export POW_TIMEOUT=2s
# Argon2 iterations
export POW_ITERATION=2
# Argon2 memory in MB
export POW_MEMORY=524288
# Number of precomputed PoW tasks
export EMISSION_SIZE=10
# Task emission delay
export EMISSION_DELAY=1s
# Connection read timeout
export CON_READ_TIMEOUT=3s
# Connection write timeout
export CON_WRITE_TIMEOUT=3s
# Number of workers
export WORKER_POOL=12
# Size of connection queue
export CON_QUEUE=256
```
