# go-templates

Several Golang templates to be used with `gonew`.

## Helpful Links

- [Golang Blog about gonew](https://go.dev/blog/gonew)

## Using a template

As stated in the official blog post, start by installing `gonew` itself:

```bash
$ go install golang.org/x/tools/cmd/gonew@latest
```

And afterwards provide a link to the specific example directory you want to use:

```bash
$ gonew github.com/torbendury/go-templates/http-server example.com/myserver
$ cd ./myserver
```
