## apigeecli keyaliases create-self-signed

Create a Key Alias from self-seigned cert

### Synopsis

Create a Key Alias by generating a self-signed cert

```
apigeecli keyaliases create-self-signed [flags]
```

### Options

```
  -s, --alias string   Name of the key alias
  -c, --cert string    Certificate in JSON format
  -x, --exp            Ignore expiry validation
  -h, --help           help for create-self-signed
  -k, --key string     Name of the key store
  -w, --nl             Ignore new line in cert chain
```

### Options inherited from parent commands

```
  -a, --account string   Path Service Account private key in JSON
      --disable-check    Disable check for newer versions
  -e, --env string       Apigee environment name
      --no-output        Disable printing API responses from the control plane
  -o, --org string       Apigee organization name
  -t, --token string     Google OAuth Token
```

### SEE ALSO

* [apigeecli keyaliases](apigeecli_keyaliases.md)	 - Manage Key Aliases

###### Auto generated by spf13/cobra on 3-Nov-2022
