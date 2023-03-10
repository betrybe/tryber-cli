const { program } = require('commander');
const axios = require('axios');
const moment = require('moment');
const spawnSync = require('child_process').spawnSync;

program
  .command('test')
  .argument('[args...]')
  .description('Test command')
  .action( async(args) => {
    const command =  (`npm test ${args}`).trim()
    console.log(`🎹 ${command}`)
    const shell = spawnSync(command, {shell: true, stdio:'inherit'})

    if(shell.status === 1) {
      // console.log(shell.stderr.toString())
    }

    const date = moment().format('YYYY-MM-DD HH:mm:ss');
    const data = JSON.stringify({ command, date });

    try {
      console.log(`📦 Payload: ${data}`)
      const response =  await axios.post('https://example.com/api/test', data, {
        headers: {
          'Content-Type': 'application/json',
        },
      });
      console.log(response.data);
    } catch (error) {
      console.error("🔥 Erro ao enviar dados para API");
    }
  });

program.parse();