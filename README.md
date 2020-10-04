# Open Graph Protocol

![Go Version](https://img.shields.io/github/go-mod/go-version/go-ogp/ogp?style=flat-square)
![License](https://img.shields.io/github/license/go-ogp/ogp?style=flat-square)
[![PkgGoDev](https://pkg.go.dev/badge/gopkg.in/ogp.v1?tab=doc)](https://pkg.go.dev/gopkg.in/ogp.v1?tab=doc)

If you have no time for reading and implementing bullshit description of [Open
Graph Protocol](https://ogp.me) for your awesome website, then this library is
for you. OGP provides fluent APIs to build standard Open Graph objects that can
be included in Go templates, with less headache.

_Note_. This is not an Open Graph metadata parser (yet)!

## Getting Started

Use `go get` to install OGP library into your project:

```sh
$ go get -u gopkg.in/ogp.v1
```

Then, you can build OpenGraph objects:

```go
// main.go
func Handler(w http.ResponseWriter, r *http.Request) {
    ...
    ogpProfile := ogp.Profile().
        Title("John Smith").
        URL("http://jsmith.me").
        Image(ogp.Image().URL("http://jsmith.me/avatar.jpg")).
        FirstName("John").
        LastName("Smith").
        Username("jsmith").
        HTML()
    tmpl.Execute("", &struct{
        OGP: ogpProfile,
        ...
    })
}
```

And, express them in Go templates:

```html
<head prefix="og: https://ogp.me/ns#">
  <!-- ... -->
  {{ .OGP }}
  <!-- ... -->
</head>
```

## License

OGP is published under MIT license.
