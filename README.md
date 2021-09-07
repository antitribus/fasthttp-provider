# fasthttp-provider

## Example:

```go
import ...

var fhp = FastHTTPProvider{
    client: &fasthttp.Client{MaxConnsPerHost: 300,
        MaxConnDuration:     60 * time.Second,
        MaxIdleConnDuration: 1 * time.Second,
        WriteTimeout:        1 * time.Second,
        ReadTimeout:         1 * time.Second,
    },
    beforeRequest: before,
    afterRequest:  after,
}

var duration = time.Millisecond * 500

type Pokemon struct {
    Name  string `json:"name"`
    Order int    `json:"order"`
}

func main() {
    req := &fasthttp.Request{}
    req.SetRequestURI("https://pokeapi.co/api/v2/pokemon/ditto")
    
    pokemon := Pokemon{}
    
    completeResponse, _ := fhp.JSON(ctx, req, &pokemon, &duration)
    
    fmt.Println(pokemon.Name) // ditto
    fmt.Println(completeResponse.statusCode()) // 200
}
```
