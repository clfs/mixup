# mixup
Shuffle standard input.

Install the binary to your current directory:

```text
GOBIN="$(pwd)" go install github.com/clfs/mixup@latest
```

Example:

```text
$ head main.go
package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)

$ head main.go | mixup
	"fmt"
	"bufio"
import (
	"os"

)
	"log"
	"math/rand/v2"
package main

```
