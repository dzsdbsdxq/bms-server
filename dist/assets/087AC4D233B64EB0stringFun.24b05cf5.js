/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
const e=e=>e[0]?e.replace(e[0],e[0].toUpperCase()):"",a=e=>e[0]?e.replace(e[0],e[0].toLowerCase()):"",r=e=>"ID"===e?"ID":e.replace(/([A-Z])/g,"_$1").toLowerCase(),t=e=>e.replace(/\_(\w)/g,(function(e,a){return a.toUpperCase()})),o=e=>{const a="ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789";let r="";for(let t=0;t<e;t++)r+=a.charAt(Math.floor(62*Math.random()));return r};export{a,e as b,t as c,o as g,r as t};
