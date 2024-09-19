## Cinedata

Cinedata is a JSON-based API designed to manage and retrieve movie information, similar to the functionality of the [Open Movie Database API](https://www.omdbapi.com/). It supports various endpoints that allow users to perform CRUD (Create, Read, Update, Delete) operations on movie data, manage user authentication, and reset passwords. The API is built for scalability and security, ensuring reliable access and interaction with movie data.

* [__API Documentation__](https://cinedata.mahmud.pro)
* [__System Design Document__](https://mahmud.pro/projects/cinedata/system-design-document)

## Technology Stack

* __Backend__: Written in __Go__ (with httprouter for RESTful routing)
* __Database__: PostgreSQL for data storage and retrieval
* __Security__: HTTPS (TLS encryption), bcrypt for password hashing, and IP-based rate limiting
* __Server & Proxy__: Hosted on a __DigitalOcean Linux server__, using __Caddy__ for reverse proxy and SSL management