/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{r as a,o as s,c as e,d as t,b as l,w as o,i as r,E as u,j as i,k as m}from"./087AC4D233B64EB0index.39f1eb31.js";/* empty css                               *//* empty css                              */import{_ as n}from"./087AC4D233B64EB0warningBar.beb718a3.js";import{l as p}from"./087AC4D233B64EB0autoCode.2e981990.js";const c={class:"p-5 bg-white"},d={class:"flex items-center gap-3"},B={__name:"pubPlug",setup(B){const C=a(""),f=async()=>{const a=await p({plugName:C.value});0===a.code&&u.success(a.msg)};return(a,u)=>{const p=i,B=m;return s(),e("div",c,[t(n,{title:"目前只支持标准插件（通过插件模板生成的标准目录插件），非标准插件请自行打包"}),l("div",d,[t(p,{modelValue:C.value,"onUpdate:modelValue":u[0]||(u[0]=a=>C.value=a),placeholder:"插件模板处填写的【插件名】"},null,8,["modelValue"]),t(B,{type:"primary",size:"small",onClick:f},{default:o((()=>[r("打包插件")])),_:1})])])}}};export{B as default};
