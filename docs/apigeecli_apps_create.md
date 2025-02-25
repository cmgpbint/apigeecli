## apigeecli apps create

Create a Developer App

### Synopsis

Create a Developer App

```
apigeecli apps create [flags]
```

### Options

```
      --attrs stringToString   Custom attributes (default [])
  -c, --callback string        The callbackUrl is used by OAuth
  -e, --email string           The developer's email or id
  -x, --expires string         A setting, in milliseconds, for the lifetime of the consumer key
  -h, --help                   help for create
  -n, --name string            Name of the developer app
  -p, --prods stringArray      A list of api products
  -s, --scopes stringArray     OAuth scopes
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
      --no-output        Disable printing API responses from the control plane
  -o, --org string       Apigee organization name
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli apps](apigeecli_apps.md)	 - Manage Apigee Developer Applications

###### Auto generated by spf13/cobra on 3-Nov-2022
