/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{ai as e,l as a,r as l,a as s,o as t,c as o,b as u,d,w as r,i,E as p,j as m,k as n}from"./087AC4D233B64EB0index.39f1eb31.js";import{a as c,E as v}from"./087AC4D233B64EB0form-item.a86c41d2.js";/* empty css                               */import{E as b}from"./087AC4D233B64EB0switch.5bfb00d1.js";/* empty css                              */import{f,c as y,u as B}from"./087AC4D233B64EB0sysExtendsDetails.3d67246c.js";import"./087AC4D233B64EB0date.225eaf4f.js";import"./087AC4D233B64EB0dictionary.2d6433e1.js";import"./087AC4D233B64EB0castArray.209ac14b.js";import"./087AC4D233B64EB0_baseClone.a73122b0.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0_initCloneObject.0c0ca3ea.js";import"./087AC4D233B64EB0sysDictionary.e0472ec6.js";const E={class:"gva-form-box"},V=Object.assign({name:"SysExtendsDetailsForm"},{__name:"sysExtendsDetailsForm",setup(V){const D=e(),_=a(),A=l(""),C=l({label:"",key:"",value:"",status:!1,sort:0,sysExtendsId:0,appid:""}),j=s({}),x=l();(async()=>{if(D.query.id){const e=await f({ID:D.query.id});0===e.code&&(C.value=e.data.resysExtendsDetails,A.value="update")}else A.value="create"})();const h=async()=>{var e;null==(e=x.value)||e.validate((async e=>{if(!e)return;let a;switch(A.value){case"create":default:a=await y(C.value);break;case"update":a=await B(C.value)}0===a.code&&p({type:"success",message:"创建/更改成功"})}))},k=()=>{_.go(-1)};return(e,a)=>{const l=m,s=c,p=b,f=n,y=v;return t(),o("div",null,[u("div",E,[d(y,{model:C.value,ref_key:"elFormRef",ref:x,"label-position":"right",rules:j,"label-width":"80px"},{default:r((()=>[d(s,{label:"展示值:",prop:"label"},{default:r((()=>[d(l,{modelValue:C.value.label,"onUpdate:modelValue":a[0]||(a[0]=e=>C.value.label=e),clearable:!0,placeholder:"请输入展示值"},null,8,["modelValue"])])),_:1}),d(s,{label:"字典值:",prop:"key"},{default:r((()=>[d(l,{modelValue:C.value.key,"onUpdate:modelValue":a[1]||(a[1]=e=>C.value.key=e),clearable:!0,placeholder:"请输入字典值"},null,8,["modelValue"])])),_:1}),d(s,{label:"扩展值:",prop:"value"},{default:r((()=>[d(l,{modelValue:C.value.value,"onUpdate:modelValue":a[2]||(a[2]=e=>C.value.value=e),clearable:!0,placeholder:"请输入扩展值"},null,8,["modelValue"])])),_:1}),d(s,{label:"启用状态:",prop:"status"},{default:r((()=>[d(p,{modelValue:C.value.status,"onUpdate:modelValue":a[3]||(a[3]=e=>C.value.status=e),"active-color":"#13ce66","inactive-color":"#ff4949","active-text":"是","inactive-text":"否",clearable:""},null,8,["modelValue"])])),_:1}),d(s,{label:"排序标记:",prop:"sort"},{default:r((()=>[d(l,{modelValue:C.value.sort,"onUpdate:modelValue":a[4]||(a[4]=e=>C.value.sort=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),d(s,{label:"关联标记ID:",prop:"sysExtendsId"},{default:r((()=>[d(l,{modelValue:C.value.sysExtendsId,"onUpdate:modelValue":a[5]||(a[5]=e=>C.value.sysExtendsId=e),modelModifiers:{number:!0},clearable:!0,placeholder:"请输入"},null,8,["modelValue"])])),_:1}),d(s,{label:"关联标记Appid:",prop:"appid"},{default:r((()=>[d(l,{modelValue:C.value.appid,"onUpdate:modelValue":a[6]||(a[6]=e=>C.value.appid=e),clearable:!0,placeholder:"请输入关联标记Appid"},null,8,["modelValue"])])),_:1}),d(s,null,{default:r((()=>[d(f,{type:"primary",onClick:h},{default:r((()=>[i("保存")])),_:1}),d(f,{type:"primary",onClick:k},{default:r((()=>[i("返回")])),_:1})])),_:1})])),_:1},8,["model","rules"])])])}}});export{V as default};
