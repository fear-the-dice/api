var drakov = require('drakov');

var argv = {
    sourceFiles: __dirname + '/server.md',
    serverPort: 3000,
    stealthmode: true,
    disableCORS: false,
    public: true
};

drakov.run(argv);
