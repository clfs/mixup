# mixup
Shuffle standard input.

Install the binary to your current directory:

```text
GOBIN="$(pwd)" go install github.com/clfs/mixup@latest
```

## Examples

Sample input:

```text
$ head main.go
package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand/v2"
	"os"
)
```

Random shuffle:

```text
$ head main.go | mixup
import (
	"os"
	"flag"

	"log"
	"fmt"
	"math/rand/v2"
package main
)
	"bufio"

```

Deterministic shuffle:

```text
$ head main.go | mixup -d
	"flag"
	"fmt"
	"os"

package main
)
	"math/rand/v2"
	"bufio"
	"log"
import (
```
