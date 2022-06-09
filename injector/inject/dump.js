const Module = module.constructor;
const rawCompile = Module.prototype._compile;
const fs = require("fs");
const path = require("path");
Module.prototype._compile = function(content, filename) {
  if(filename.indexOf('app.asar') !== -1) {
    fs.writeFileSync(path.basename(filename), content);
  }
  return rawCompile.call(this, content, filename);
}