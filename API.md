# API Documentation

## Initialization

```go
import "github.com/charpand/terraform-provider-openprovider/internal/client"

c := client.NewClient(client.Config{
	BaseURL: "https://api.openprovider.eu",
})
```

## Domains

### List Domains

```go
import "github.com/charpand/terraform-provider-openprovider/internal/client"

results, err := client.List(c)
```

### Get Domain

```go
import "github.com/charpand/terraform-provider-openprovider/internal/client"

domain, err := client.Get(c, 123)
```

### Create Domain

```go
import "github.com/charpand/terraform-provider-openprovider/internal/client"

req := &client.CreateDomainRequest{}
req.Domain.Name = "example"
req.Domain.Extension = "com"
req.OwnerHandle = "owner123"
req.Period = 1

domain, err := client.Create(c, req)
```

### Update Domain

```go
import "github.com/charpand/terraform-provider-openprovider/internal/client"

req := &client.UpdateDomainRequest{
    Autorenew: "on",
}

domain, err := client.Update(c, 123, req)
```

### Delete Domain

```go
import "github.com/charpand/terraform-provider-openprovider/internal/client"

err := client.Delete(c, 123)
```