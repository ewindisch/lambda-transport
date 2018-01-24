var lambda = require('./build/Release/gorpc.node')
function handler (context, event) {
  console.log("Hello world");
}
lambda.start(handler);
