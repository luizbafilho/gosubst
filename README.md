# gosubst

`gosubst` is an `envsubst` on steroids. Receives a template file in the `stdin` replaces all variables and sends it to `stdout`.

## Installing
Find the latest gosubst for your platform on the [releases](https://github.com/luizbafilho/gosubst/releases) page
```
curl -o /usr/local/bin/gosubst -sSL https://github.com/luizbafilho/gosubst/releases/download/<version>/gosubst_<os>-<arch>
chmod +x /usr/local/bin/gosubst
```

## Usage
`gosubst` copies the `stdin` to `stdout` replacing all variables.

```shell
gosubst copies stardard input to standard output replacing all variables present in values file

Usage:
  gosubst [flags]

Flags:
      --type string     values type (toml, yaml or json) (default "yaml")
  -v, --values string   values file
```

Values file sample:
```yaml
foo: "bar"
baz:
 boo: "bla"
 zoo: "mee"
```

Template file sample:

```
Sample access: {{.foo}}
Accessing variables: {{.baz.boo}} | {{.baz.zoo}}
```

```shell
$ gosubst -f values.yaml < template.conf

Sample access: bar
Accessing variables: bla | mee
```

### Values Files
The values file provides all the values to be replaced when the template is processed. `gosubst` supports `yaml`, `json` or `toml` files.

### Environment variables

`gosubst` makes all environment variables available in the template.

```shell
$ echo "Home path: {{.Env.HOME}}" | gosubst
Home path: /home/vagrant
```

### Template functions
Given that `gosubst` uses `Go Templates` you can leverage all the power `Go` provide. Here is a small sample:

```
{{with .Account -}}
Dear {{.FirstName}} {{.LastName}},  
{{- end}}

Below are your account statement details for period from {{.FromDate}} to {{.ToDate}}.

{{if .Purchases -}}
  Your purchases:
  {{- range .Purchases }}
      {{ .Date}} {{ printf "%-20s" .Description }} {{.AmountInCents -}}
  {{- end}}
{{- else}}
You didn't make any purchases during the period.  
{{- end}}

{{$note := .Account.Note -}}
{{if $note -}}
Note: {{$note}}  
{{- end}}

Best Wishes,  
Customer Service  
```

For more details: https://golang.org/pkg/text/template/#hdr-Text_and_spaces

### Helper functions
`gosubst` makes available a lot of usefull functions provided by [Sprig](https://github.com/Masterminds/sprig)


```shell
$ echo "generated-id: {{uuidv4}}" | gosubst
generated-id: 00a315b7-c846-4751-b2ee-b50eb5359bac

$ echo "generated-sha: {{sha256sum .Env.USER | trunc 7}}" | gosubst
generated-sha: 4b569e80caa6ba1d7416f4aa2177a85e316afe1142d606bcfd6b70af0c0bf666
```

Check it out all functions available: https://github.com/Masterminds/sprig#functions
