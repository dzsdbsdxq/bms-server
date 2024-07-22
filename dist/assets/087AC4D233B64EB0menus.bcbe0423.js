/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{_ as e,r as a,L as t,o,c as s,b as l,d as r,w as n,i,t as u,n as d,h as c,aM as m,aN as p,E as B,aO as y,a8 as f,j as h,k as C}from"./087AC4D233B64EB0index.39f1eb31.js";import{E}from"./087AC4D233B64EB0dialog.c75a9340.js";import"./087AC4D233B64EB0overlay.3905897c.js";import{E as D,a as j}from"./087AC4D233B64EB0table-column.e9b4a870.js";import"./087AC4D233B64EB0checkbox.a3853383.js";/* empty css                                *//* empty css                               */import{E as v}from"./087AC4D233B64EB0scrollbar.3894915f.js";import"./087AC4D233B64EB0tag.d3c1c3e4.js";import{E as A}from"./087AC4D233B64EB0tree.2b9a65c6.js";/* empty css                               *//* empty css                              */import{u as k}from"./087AC4D233B64EB0authority.e5789bef.js";import{g as w,s as I}from"./087AC4D233B64EB0authorityBtn.0890d580.js";import"./087AC4D233B64EB0isUndefined.d281c39a.js";import"./087AC4D233B64EB0index.d8bfc0de.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0debounce.0aab18a9.js";import"./087AC4D233B64EB0_initCloneObject.0c0ca3ea.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0index.bfb39b2a.js";const b={class:"sticky top-0.5 z-10 bg-white"},g={class:"clear-both tree-content"},_={class:"custom-tree-node"},x={key:0},R={class:"dialog-footer"},z=e(Object.assign({name:"Menus"},{__name:"menus",props:{row:{default:function(){return{}},type:Object}},emits:["changeRow"],setup(e,{expose:z,emit:N}){const V=e,O=N,U=a(""),M=a([]),S=a([]),T=a(!1),q=a({children:"children",label:function(e){return e.meta.title}});(async()=>{const e=await m();M.value=e.data.menus;const a=(await p({authorityId:V.row.authorityId})).data.menus,t=[];a.forEach((e=>{a.some((a=>a.parentId===e.menuId))||t.push(Number(e.menuId))})),S.value=t})();const G=()=>{T.value=!0},J=a(null),L=async()=>{const e=J.value.getCheckedNodes(!1,!0);0===(await y({menus:e,authorityId:V.row.authorityId})).code&&B({type:"success",message:"菜单设置成功!"})};z({enterAndNext:()=>{L()},needConfirm:T});const $=a(!1),F=a([]),H=a([]),K=a();let P="";const Q=e=>{H.value=e},W=e=>{$.value=!0,F.value=e.menuBtn},X=()=>{$.value=!1},Y=async()=>{const e=H.value.map((e=>e.ID));0===(await I({menuID:P,selected:e,authorityId:V.row.authorityId})).code&&(B({type:"success",message:"设置成功"}),$.value=!1)},Z=(e,a)=>!e||-1!==a.meta.title.indexOf(e);return t(U,(e=>{J.value.filter(e)})),(a,t)=>{const m=h,p=C,y=A,I=v,z=D,N=j,T=E;return o(),s("div",null,[l("div",b,[r(m,{modelValue:U.value,"onUpdate:modelValue":t[0]||(t[0]=e=>U.value=e),class:"w-3/5",placeholder:"筛选"},null,8,["modelValue"]),r(p,{class:"float-right",type:"primary",size:"small",onClick:L},{default:n((()=>[i("确 定")])),_:1})]),l("div",g,[r(I,null,{default:n((()=>[r(y,{ref_key:"menuTree",ref:J,data:M.value,"default-checked-keys":S.value,props:q.value,"default-expand-all":"","highlight-current":"","node-key":"ID","show-checkbox":"","filter-node-method":Z,onCheck:G},{default:n((({node:a,data:t})=>[l("span",_,[l("span",null,u(a.label),1),l("span",null,[r(p,{type:"primary",link:"",size:"small",style:d({color:e.row.defaultRouter===t.name?"#E6A23C":"#85ce61"}),disabled:!a.checked,onClick:()=>(async e=>{const a=await k({authorityId:V.row.authorityId,AuthorityName:V.row.authorityName,parentId:V.row.parentId,defaultRouter:e.name});0===a.code&&(B({type:"success",message:"设置成功"}),O("changeRow","defaultRouter",a.data.authority.defaultRouter))})(t)},{default:n((()=>[i(u(e.row.defaultRouter===t.name?"首页":"设为首页"),1)])),_:2},1032,["style","disabled","onClick"])]),t.menuBtn.length?(o(),s("span",x,[r(p,{type:"primary",link:"",size:"small",onClick:()=>(async e=>{P=e.ID;const a=await w({menuID:P,authorityId:V.row.authorityId});0===a.code&&(W(e),await f(),a.data.selected&&a.data.selected.forEach((e=>{F.value.some((a=>{a.ID===e&&K.value.toggleRowSelection(a,!0)}))})))})(t)},{default:n((()=>[i(" 分配按钮 ")])),_:2},1032,["onClick"])])):c("",!0)])])),_:1},8,["data","default-checked-keys","props"])])),_:1})]),r(T,{modelValue:$.value,"onUpdate:modelValue":t[1]||(t[1]=e=>$.value=e),title:"分配按钮","destroy-on-close":""},{footer:n((()=>[l("div",R,[r(p,{size:"small",onClick:X},{default:n((()=>[i("取 消")])),_:1}),r(p,{size:"small",type:"primary",onClick:Y},{default:n((()=>[i("确 定")])),_:1})])])),default:n((()=>[r(N,{ref_key:"btnTableRef",ref:K,data:F.value,"row-key":"ID",onSelectionChange:Q},{default:n((()=>[r(z,{type:"selection",width:"55"}),r(z,{label:"按钮名称",prop:"name"}),r(z,{label:"按钮备注",prop:"desc"})])),_:1},8,["data"])])),_:1},8,["modelValue"])])}}}),[["__scopeId","data-v-889b2610"]]);export{z as default};