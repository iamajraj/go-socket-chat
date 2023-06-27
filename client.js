const net = require('net');
const { stdin, stdout } = require('process');
const readline = require('readline');

const read = readline.createInterface(stdin, stdout);
const socket = net.createConnection({
  host: '0.0.0.0',
  port: 8000,
});

socket.on('connect', () => {
  console.log('Connected');
  sendMsg();
});

socket.on('data', (data) => {
  readline.moveCursor(stdout, null, -1);
  readline.clearLine(stdout, 0);
  console.log(`${data.toString()}`);
});

function sendMsg() {
  read.question('-> ', (ans) => {
    socket.write(ans, (err) => {
      if (err) {
        console.log(err);
      }
    });

    sendMsg();
  });
}

socket.on('close', () => {
  console.log('Connection end');
});
