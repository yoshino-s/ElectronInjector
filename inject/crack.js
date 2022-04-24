const Module = module.constructor;
const rawCompile = Module.prototype._compile;
const replaceTable = [
  [/\.getHasLicense=\(\)=>\w+,/, ".getHasLicense=()=>true,"],
  [/get\("SLicense"\);/, `get("SLicense");return{deviceId:'1',fingerprint:'1',email:'1',license:'1',version:'1',date:'1',failCounts:0, lastRetry: new Date()};`]
];
const oldReplaceTable = [
  ['exports.shouldShowNoLicenseHint=shouldShowNoLicenseHint', 'exports.shouldShowNoLicenseHint=()=>false'],
  ['exports.getHasLicense=getHasLicense', 'exports.getHasLicense=()=>true'],
  ['return fingerPrint', 'return "1"'],
  [`readLicenseInfo=()=>{const e=getLicenseLocalStore().get("SLicense");if(!e)return null;const[n,t,i]=e.split("#"),a=decrypt(n);return a.fingerprint!=fingerPrint?null:(Object.assign(a,{failCounts:t,lastRetry:new Date(i)}),a)}`,
    `readLicenseInfo=()=>{return{deviceId:'1',fingerprint:'1',email:'1',license:'1',version:'1',date:'1',failCounts:0, lastRetry: new Date()}}`]
]
Module.prototype._compile = function (content, filename) {
  if (filename.indexOf('atom.js') !== -1) { // over 1.2.4
    for (const [regex, replace] of replaceTable) {
      content = content.replace(regex, replace);
    }
  }
  if (filename.indexOf("License.js") !== -1) { // before 1.2.4
    for (const [regex, replace] of oldReplaceTable) {
      content = content.replace(regex, replace);
    }
  }
  return rawCompile.call(this, content, filename);
}