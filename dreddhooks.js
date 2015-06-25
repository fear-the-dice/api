var hooks = require('hooks');
var spawn = require('child_process').spawn;

hooks.beforeAll(function (transactions, done) {
  hooks.log("executing 'mongorestore --drop'");

  ls = spawn('mongorestore', ['--drop']);

  ls.stdout.on('data', function (data) {
    hooks.log('' + data);
  });

  ls.stderr.on('data', function (data) {
    hooks.log('' + data);
  });

  ls.on('close', function (code) {
    hooks.log('mongorestore exited with code ' + code);
    done();
  });
});
