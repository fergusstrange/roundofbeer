const fs = require('fs');

const templateFile = process.argv[2];
const domain = process.argv[3];

if (!templateFile || !domain) {
    process.exit(1);
}

let templateJson = JSON.parse(fs.readFileSync(templateFile, 'utf-8'));
let outputJson = JSON.stringify(Object.assign({}, templateJson, {
    stages: {
        default: {
            domain: domain,
        },
    },
}));

process.stdout.write(outputJson);
