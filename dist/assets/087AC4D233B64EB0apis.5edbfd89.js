/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{K as a,r as e,L as t,o as s,c as o,b as r,d as l,w as i,i as n,E as p,j as c,k as d}from"./087AC4D233B64EB0index.39f1eb31.js";import{E as u}from"./087AC4D233B64EB0scrollbar.3894915f.js";import{E as h}from"./087AC4D233B64EB0tree.2b9a65c6.js";import"./087AC4D233B64EB0checkbox.a3853383.js";/* empty css                               *//* empty css                              */import{e as m}from"./087AC4D233B64EB0api.db29bd05.js";import"./087AC4D233B64EB0index.bfb39b2a.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";const f={class:"sticky top-0.5 z-10 bg-white"},y={class:"tree-content"},B=Object.assign({name:"Apis"},{__name:"apis",props:{row:{default:function(){return{}},type:Object}},setup(B,{expose:v}){const E=B,j=e({children:"children",label:"description"}),A=e(""),C=e([]),b=e([]),k=e("");(async()=>{const e=(await m()).data.apis;C.value=I(e);const t=await(s={authorityId:E.row.authorityId},a({url:"/casbin/getPolicyPathByAuthorityId",method:"post",data:s}));var s;k.value=E.row.authorityId,b.value=[],t.data.paths&&t.data.paths.forEach((a=>{b.value.push("p:"+a.path+"m:"+a.method)}))})();const D=e(!1),w=()=>{D.value=!0},I=a=>{const e={};a&&a.forEach((a=>{a.onlyId="p:"+a.path+"m:"+a.method,Object.prototype.hasOwnProperty.call(e,a.apiGroup)?e[a.apiGroup].push(a):Object.assign(e,{[a.apiGroup]:[a]})}));const t=[];for(const s in e){const a={ID:s,description:s+"组",children:e[s]};t.push(a)}return t},g=e(null),x=async()=>{const e=g.value.getCheckedNodes(!0);var t=[];e&&e.forEach((a=>{var e={path:a.path,method:a.method};t.push(e)}));var s;0===(await(s={authorityId:k.value,casbinInfos:t},a({url:"/casbin/updateCasbin",method:"post",data:s}))).code&&p({type:"success",message:"api设置成功"})};v({needConfirm:D,enterAndNext:()=>{x()}});const O=(a,e)=>!a||-1!==e.description.indexOf(a);return t(A,(a=>{g.value.filter(a)})),(a,e)=>{const t=c,p=d,m=h,B=u;return s(),o("div",null,[r("div",f,[l(t,{modelValue:A.value,"onUpdate:modelValue":e[0]||(e[0]=a=>A.value=a),class:"w-3/5",size:"small",placeholder:"筛选"},null,8,["modelValue"]),l(p,{class:"float-right",type:"primary",size:"small",onClick:x},{default:i((()=>[n("确 定")])),_:1})]),r("div",y,[l(B,null,{default:i((()=>[l(m,{ref_key:"apiTree",ref:g,data:C.value,"default-checked-keys":b.value,props:j.value,"default-expand-all":"","highlight-current":"","node-key":"onlyId","show-checkbox":"","filter-node-method":O,onCheck:w},null,8,["data","default-checked-keys","props"])])),_:1})])])}}});export{B as default};
