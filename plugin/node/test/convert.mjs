import fs from "fs"
import yaml from 'js-yaml';
const config =  fs.readFileSync("../../../../../../.drone.js");

const wrappedScript = `
    (function() {
      ${config}
      return { Pipeline,pipelines };
    })();
  `;

const res = eval(wrappedScript);
const yamlStr = await yaml.dump(res.Pipeline(res.pipelines[0]));
console.log(yamlStr)