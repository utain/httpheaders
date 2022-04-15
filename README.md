# HttpHeaders

Constant of [HTTP headers](https://developer.mozilla.org/en-US/docs/Web/HTTP/Headers)

```bash
go get github.com/utain/httpheaders
```

### Example

```go
import (
	"testing"
	"net/http/httptest"

	"github.com/utain/httpheaders"
	"github.com/utain/httpheaders/mediatypes"
	"github.com/utain/httpheaders/methods"
)

func TestExample(t *testing.T) {
  // ...
  req := httptest.NewRequest(methods.GET, "/api/xyz", nil)
  req.Header.Add(httpheaders.ContentType, mediatypes.ApplicationJsonUtf8)
  // ...
}
```