# Barlus API - Dev version

[![GitHub license](https://img.shields.io/github/license/barlus-engineer/barlus-api)](https://github.com/barlus-engineer/barlus-api/blob/main/LICENSE)
[![GitHub stars](https://img.shields.io/github/stars/barlus-engineer/barlus-api)](https://github.com/barlus-engineer/barlus-api/stargazers)

Barlus API is a backend service for barlus website

## Installation

### Prerequisites
- **Go**: version **1.18**+
- **PostgreSQL**: version **10**+
- **Redis**: version **9**

### Clone the repository
```sh
git clone https://github.com/barlus-engineer/barlus-api.git
cd barlus-api
```

### Environment Variables
```
SERVER_NAME=[server name] # default (Barlus API)
SERVER_VERSION=[server version] # default (v0.1.1)
SERVER_RELEASE=[true or false] # default (false)

HTTP_HOST=[http host]
HTTP_PORT=[http port]

POSTGRES_URL=[Postgres url]

REDIS_URL=[your redis url]
```