var drakov = require('drakov');

var argv = {
    sourceFiles: __dirname + '/apiary.apib',
    serverPort: 3000,
    stealthmode: true,
    disableCORS: false,
    public: true
};

drakov.run(argv);
