# Be careful

Unfortunately openapi spec from supabase auth is bugged

You must change
```
web_authn_credential:
        web_authn_credential:
          type: jsonb
```

to
  
```

        web_authn_credential:
          type: object
          additionalProperties: true
```

Because type jsonb doesn't exist

Then you must change in client.go (after generation) and add following somewhere

```

type N200Type string

const (
	N200TypePhone N200Type = "phone"	
	N200TypeTotp N200Type = "totp"
	N200TypeWebauthn N200Type = "webauthn"
)
```