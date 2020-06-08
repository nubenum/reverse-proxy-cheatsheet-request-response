# Reverse Proxy Cheatsheet Experimental Setup

Some live examples matching the setup of the cheatsheet.

1. Check out the repo.
2. Set the example you want to test (e.g. `1-nothing`, `2-rewrite`, ..., see apache or nginx directories) in the `.env` file.
3. Start the backend and the two reverse proxies with  `docker-compose up -d`.
4. Access the reverse proxies at http://localhost:8081 and http://localhost:8082