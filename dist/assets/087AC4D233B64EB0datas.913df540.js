/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{r as a,o as t,c as e,d as l,b as o,w as r,i,F as s,m as u,g as n,t as d,E as h,k as c}from"./087AC4D233B64EB0index.39f1eb31.js";import{E as y,a as m}from"./087AC4D233B64EB0checkbox.a3853383.js";/* empty css                               */import{s as f}from"./087AC4D233B64EB0authority.e5789bef.js";import{_ as p}from"./087AC4D233B64EB0warningBar.beb718a3.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";const v={class:"sticky top-0.5 z-10 bg-white my-4"},I={class:"clear-both pt-4"},A=Object.assign({name:"Datas"},{__name:"datas",props:{row:{default:function(){return{}},type:Object},authority:{default:function(){return[]},type:Array}},emits:["changeRow"],setup(A,{expose:B,emit:g}){const w=A,C=a([]),E=a(!1),j=a=>{a&&a.forEach((a=>{const t={};t.authorityId=a.authorityId,t.authorityName=a.authorityName,C.value.push(t),a.children&&a.children.length&&j(a.children)}))},_=a([]);j(w.authority),w.row.dataAuthorityId&&w.row.dataAuthorityId.forEach((a=>{const t=C.value&&C.value.filter((t=>t.authorityId===a.authorityId))&&C.value.filter((t=>t.authorityId===a.authorityId))[0];_.value.push(t)}));const b=g,k=()=>{_.value=[...C.value],b("changeRow","dataAuthorityId",_.value),E.value=!0},D=()=>{_.value=C.value.filter((a=>a.authorityId===w.row.authorityId)),b("changeRow","dataAuthorityId",_.value),E.value=!0},x=()=>{const a=[];z(w.row,a),_.value=C.value.filter((t=>a.indexOf(t.authorityId)>-1)),b("changeRow","dataAuthorityId",_.value),E.value=!0},z=(a,t)=>{t.push(a.authorityId),a.children&&a.children.forEach((a=>{z(a,t)}))},R=async()=>{0===(await f(w.row)).code&&h({type:"success",message:"资源设置成功"})},N=()=>{b("changeRow","dataAuthorityId",_.value),E.value=!0};return B({enterAndNext:()=>{R()},needConfirm:E}),(a,h)=>{const f=c,A=y,B=m;return t(),e("div",null,[l(p,{title:"此功能仅用于创建角色和角色的many2many关系表，具体使用还须自己结合表实现业务，详情参考示例代码（客户示例）。此功能不建议使用，建议使用插件市场【组织管理功能（点击前往）】来管理资源权限。",href:"https://plugin.gin-vue-admin.com/#/layout/newPluginInfo?id=36"}),o("div",v,[l(f,{class:"float-left",type:"primary",size:"small",onClick:k},{default:r((()=>[i("全选")])),_:1}),l(f,{class:"float-left",type:"primary",size:"small",onClick:D},{default:r((()=>[i("本角色")])),_:1}),l(f,{class:"float-left",type:"primary",size:"small",onClick:x},{default:r((()=>[i("本角色及子角色")])),_:1}),l(f,{class:"float-right",type:"primary",size:"small",onClick:R},{default:r((()=>[i("确 定")])),_:1})]),o("div",I,[l(B,{modelValue:_.value,"onUpdate:modelValue":h[0]||(h[0]=a=>_.value=a),onChange:N},{default:r((()=>[(t(!0),e(s,null,u(C.value,((a,e)=>(t(),n(A,{key:e,label:a},{default:r((()=>[i(d(a.authorityName),1)])),_:2},1032,["label"])))),128))])),_:1},8,["modelValue"])])])}}});export{A as default};
