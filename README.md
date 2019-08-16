[![Go Report Card](https://goreportcard.com/badge/github.com/Luzifer/sri)](https://goreportcard.com/report/github.com/Luzifer/sri)
![](https://badges.fyi/github/license/Luzifer/sri)
![](https://badges.fyi/github/downloads/Luzifer/sri)
![](https://badges.fyi/github/latest-release/Luzifer/sri)
![](https://knut.in/project-status/sri)

# Luzifer / sri

`sri` is a very small helper to calculate [SRI](https://www.w3.org/TR/SRI/) information for `<link>` or `<script>` tags.

Files are downloaded as they currently are and a checksum is calculated. This checksum then is printed inside the desired HTML tag for embedding into an HTML page.

Please be aware you should use this with non-changing URLs as the browser will no longer load the file as soon as the hash does no longer match.

## Usage

```console
# sri --help
Usage of sri:
      --html               Print HTML tags with SRI information (If disabled just prints the hashes) (default true)
      --html-tag string    Tag to use for HTML mode (supported: link, script) (default "link")
      --log-level string   Log level (debug, info, warn, error, fatal) (default "info")
      --version            Prints current version and exits

# sri --html-tag link https://use.fontawesome.com/releases/v5.10.1/css/all.css
<link href="https://use.fontawesome.com/releases/v5.10.1/css/all.css" integrity="sha512-9my9Mb2+0YO+I4PUCSwUYO7sEK21Y0STBAiFEYoWtd2VzLEZZ4QARDrZ30hdM1GlioHJ8o8cWQiy8IAb1hy/Hg==" crossorigin="anonymous">
```
