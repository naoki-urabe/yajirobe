var express = require('express')
var path = require('path')
var serveStatic = require('serve-static')
const app = express();
app.use(serveStatic(__dirname + '/dist'));
var host = process.env.VUE_APP_HOST
var port = process.env.VUE_APP_FRONT_PORT;
app.route('/*')
    .get(function(req, res) {
          res.sendFile(path.join(__dirname + '/dist/index.html'));
});
app.listen(port)
console.log('server started');
console.log(`http://${host}:${port}`);
