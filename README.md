Bring your own runtime to AWS Lambda!

This project offers a C compatible library for applications to
import and run on AWS Lambda. Currently supported is NodeJS,
but bindings could be written for any language.

We do NOT use a wrapper or loader, but instead offer
library methods that emulate a Golang 1.x application.

Supported language runtimes:
- NodeJS

# Usage - NodeJS

The npm module currently ships a pre-compiled binary.
Changes in the future will properly compile this or
utilize node-pre-gyp. Until then, if the npm module does
not work, you may need to build this project locally
using `npm run build`.

1. `npm install -S @ewindisch/goomba` or checkout this package,
run `npm run build` and install it into your project with `npm install`.
2. Copy a NodeJS binary into your project as file `node`.
3. Include the following as the file `main` with mode `755`:

```
#!./node
var lambda = require('@ewindisch/goomba')
function handler (context, event) {
  console.log("Hello world");
}
lambda.start(handler);
```

Finally, package as a zip and deploy to Lambda!

License: Apache/2.0
