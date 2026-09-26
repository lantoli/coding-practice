#!/usr/bin/env node
/*
Tuenti Challenge 4, Apr 2014, Challenge 6 - Man in the middle

Code based on client.js and server.js provided in the challenge.
*/

if ((process.version.split('.')[1]|0) < 10) {
	console.log('Please, upgrade your node version to 0.10+');
	process.exit();
}

var net = require('net');
var util = require('util');
var crypto = require('crypto');


var options = {
	'port': 6969,
	'host': '54.83.207.90'
}

var dh, secret, keyServer, KEYPHRASE;

process.stdin.resume();
process.stdin.once('data', function(data) {
    KEYPHRASE = data.toString().trim();
    var socket = net.connect(options, function() {
    });

    socket.on('data', function(data) {

        data = data.toString()
        clientToServer = data[0] == 'C'; // can start with CLIENT->SERVER or SERVER->CLIENT, both lengths are 15
        data = data.substring(15);
        dataParams = data.trim().split('|');

        if (clientToServer) {
            if (data.indexOf("key|") == 0) {
                dh = crypto.createDiffieHellman(256);
                dh.generateKeys();
                socket.write(util.format('key|%s|%s\n', dh.getPrime('hex'), dh.getPublicKey('hex')));

            } else if (data.indexOf("keyphrase|") == 0) {
                secret = dh.computeSecret(keyServer, 'hex');
                var cipher = crypto.createCipheriv('aes-256-ecb', secret, '');
                var keyphrase = cipher.update(KEYPHRASE, 'utf8', 'hex') + cipher.final('hex');
                socket.write(util.format('keyphrase|%s\n', keyphrase));

            } else {
                socket.write(data);
            }
        } else {
            if (data.indexOf("key|") == 0) {
                keyServer = dataParams[1]
                socket.write(data);
            } else if (data.indexOf("result|") == 0) {
                result = dataParams[1]
                var decipher = crypto.createDecipheriv('aes-256-ecb', secret, '');
                var message = decipher.update(result, 'hex', 'utf8') + decipher.final('utf8');
                console.log(message);
                socket.end();
            } else {
                socket.write(data);
            }

        }
    });
});
