/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{bf as r,br as t,bg as e,bi as n}from"./087AC4D233B64EB0index.39f1eb31.js";import{f as o,h as u,j as a,o as f,U as c}from"./087AC4D233B64EB0_Uint8Array.02b83471.js";var s=Object.create;const i=function(){function t(){}return function(e){if(!r(e))return{};if(s)return s(e);t.prototype=e;var n=new t;return t.prototype=void 0,n}}();function p(r,t){var e=-1,n=r.length;for(t||(t=Array(n));++e<n;)t[e]=r[e];return t}function v(r,n,o,u){var a=!o;o||(o={});for(var f=-1,c=n.length;++f<c;){var s=n[f],i=u?u(o[s],r[s],s,o,r):void 0;void 0===i&&(i=r[s]),a?t(o,s,i):e(o,s,i)}return o}var b=Object.prototype.hasOwnProperty;function l(t){if(!r(t))return function(r){var t=[];if(null!=r)for(var e in Object(r))t.push(e);return t}(t);var e=o(t),n=[];for(var u in t)("constructor"!=u||!e&&b.call(t,u))&&n.push(u);return n}function y(r){return u(r)?a(r,!0):l(r)}const d=f(Object.getPrototypeOf,Object);var j="object"==typeof exports&&exports&&!exports.nodeType&&exports,h=j&&"object"==typeof module&&module&&!module.nodeType&&module,g=h&&h.exports===j?n.Buffer:void 0,m=g?g.allocUnsafe:void 0;function O(r,t){if(t)return r.slice();var e=r.length,n=m?m(e):new r.constructor(e);return r.copy(n),n}function w(r){var t=new r.constructor(r.byteLength);return new c(t).set(new c(r)),t}function x(r,t){var e=t?w(r.buffer):r.buffer;return new r.constructor(e,r.byteOffset,r.length)}function B(r){return"function"!=typeof r.constructor||o(r)?{}:i(d(r))}export{w as a,x as b,v as c,p as d,O as e,d as g,B as i,y as k};
