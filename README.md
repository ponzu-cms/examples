# Ponzu Examples
Examples of full Ponzu systems, scripts, data files, and more to demonstrate and test features

## Contents

### [News Article Demo](https://github.com/ponzu-cms/examples/tree/master/news)
News includes a demo featuring a news/publication with articles and authors as content types.

The types implement `item.Pushable` and `db.Searchable` to demonstrate the HTTP/2 Server Push and full-text search support. To test the HTTP/2 Server Push feature, you must run ponzu with the `--devhttps` flag like so:
```bash
$ ponzu --devhttps run
```
  - This generates self-signed certificates and starts up a TLS server, which is required by HTTP/2.

---


## Contributions

Please feel free to contribute your own examples. They help the community and save people time! Please follow the format specified above, and keep your example up to date with new versions. Thanks!
