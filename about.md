# PasteBox

A simple, modern almost drop-in replacement for Hastebin. This was made quickly as a drop-in for us, since the original hastebin repo that we used disappeared from GitHub.

Most of the config options remain exactly the same, as do the API endpoints. However, there are some feature differences.

The frontend is written in Vue, with TailwindCSS. GO powers the backend

Source Code: https://github.com/firstdarkdev/pastebox/

---

## Development Overview

### Implemented Features:

- [X] Document Pasting
- [X] Document Copying
- [X] Raw Document Viewing
- [X] Shortcut Keys

### Storage Backends:

- [X] File-Based Storage
- [ ] S3 Storage
- [ ] Google Data-Store
- [ ] Memcached
- [ ] Mongo
- [ ] PostgreSQL
- [ ] Redis
- [ ] RethinkDB

### Key Generators:

- [X] Phonetic
- [X] Alphabetic (random)
- [X] Dictionary (this doesn't work yet)

### Other:

- [ ] Documentation

---

## License

This application is licensed under the MIT license. It does NOT use any code from the original haste server, and was completely written from scratch. The features are just similar, because we needed a 1:1 drop in for our own use.
