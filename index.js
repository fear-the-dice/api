var drakov = require('drakov');

console.log(__dirname + '/server.md');

var argv = {
    sourceFiles: __dirname + '/server.md',
    serverPort: 3000,
    stealthmode: true,
    disableCORS: false,
    public: true
};

drakov.run(argv);
