Bring your own runtime to AWS Lambda!

This project offers a C compatible library for applications to
import and run on AWS Lambda. Currently supported is NodeJS,
but bindings could be written for any language.

We do NOT use a wrapper or loader, but instead offer
library methods that emulate a Golang 1.x application.

Supported language runtimes:
- NodeJS

# Usage

*WORK IN PROGRESS* installation via npm will make these instructions functional

1. Copy a NodeJS binary into your project as file `node`.
2. Include the following as the file `main` with mode `755`:

```
#!./node
var lambda = require('@ewindisch/goomba')
function handler (context, event) {
  console.log("Hello world");
}
lambda.start(handler);
```

License: Apache/2.0
