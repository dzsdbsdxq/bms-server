/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{_ as l,a as e,o as a,c as o,b as s,d as u,w as n,g as t,F as d,m as p,D as r,aI as i,c$ as m,h as c,i as b,E as V,P as f,j as y,k as B}from"./087AC4D233B64EB0index.39f1eb31.js";import{a as k,E as C}from"./087AC4D233B64EB0form-item.a86c41d2.js";/* empty css                               *//* empty css                              */import"./087AC4D233B64EB0tag.d3c1c3e4.js";import{E as h,a as g}from"./087AC4D233B64EB0select.7f5b5538.js";import"./087AC4D233B64EB0scrollbar.3894915f.js";/* empty css                               */import{E}from"./087AC4D233B64EB0checkbox.a3853383.js";import{b as v}from"./087AC4D233B64EB0stringFun.24b05cf5.js";import{k as U}from"./087AC4D233B64EB0autoCode.2e981990.js";import"./087AC4D233B64EB0castArray.209ac14b.js";import"./087AC4D233B64EB0_baseClone.a73122b0.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0_initCloneObject.0c0ca3ea.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";import"./087AC4D233B64EB0index.d8bfc0de.js";import"./087AC4D233B64EB0isUndefined.d281c39a.js";import"./087AC4D233B64EB0index.5d58b83c.js";import"./087AC4D233B64EB0strings.2df7e9e1.js";import"./087AC4D233B64EB0debounce.0aab18a9.js";const j={class:"gva-table-box"},A=l({__name:"autoPlug",setup(l){const A=e({plugName:"",routerGroup:"",hasGlobal:!0,hasRequest:!0,hasResponse:!0,global:[{key:"",type:"",desc:""}],request:[{key:"",type:"",desc:""}],response:[{key:"",type:"",desc:""}]}),D=()=>{A.plugName=v(A.plugName)},_=async()=>{if(!A.plugName||!A.routerGroup)return void V.error("插件名称和插件路由组为必填项");if(A.hasGlobal){if(A.global.some((l=>{if(!l.key||!l.type)return!0})))return void V.error("全局属性的key和type为必填项")}if(A.hasRequest){if(A.request.some((l=>{if(!l.key||!l.type)return!0})))return void V.error("请求属性的key和type为必填项")}if(A.hasResponse){if(A.response.some((l=>{if(!l.key||!l.type)return!0})))return void V.error("响应属性的key和type为必填项")}0===(await U(A)).code&&f("创建成功，插件已自动写入后端plugin目录下，请按照自己的逻辑进行创造")},R=l=>{l.push({key:"",value:""})},q=(l,e)=>{1!==l.length?l.splice(e,1):V.warning("至少有一个全局属性")};return(l,e)=>{const V=y,f=k,v=E,U=h,G=g,w=B,x=C;return a(),o("div",null,[s("div",j,[u(x,{"label-width":"140px",class:"w-[680px]"},{default:n((()=>[u(f,{label:"插件名"},{default:n((()=>[u(V,{modelValue:A.plugName,"onUpdate:modelValue":e[0]||(e[0]=l=>A.plugName=l),placeholder:"必填（英文大写字母开头）",onBlur:D},null,8,["modelValue"])])),_:1}),u(f,{label:"路由组"},{default:n((()=>[u(V,{modelValue:A.routerGroup,"onUpdate:modelValue":e[1]||(e[1]=l=>A.routerGroup=l),placeholder:"将会作为插件路由组使用"},null,8,["modelValue"])])),_:1}),u(f,{label:"使用全局属性"},{default:n((()=>[u(v,{modelValue:A.hasGlobal,"onUpdate:modelValue":e[2]||(e[2]=l=>A.hasGlobal=l)},null,8,["modelValue"])])),_:1}),A.hasGlobal?(a(),t(f,{key:0,label:"全局属性"},{default:n((()=>[(a(!0),o(d,null,p(A.global,((l,t)=>(a(),o("div",{key:t,class:"plug-row"},[s("span",null,[u(V,{modelValue:l.key,"onUpdate:modelValue":e=>l.key=e,placeholder:"key 必填"},null,8,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(G,{modelValue:l.type,"onUpdate:modelValue":e=>l.type=e,placeholder:"type 必填"},{default:n((()=>[u(U,{label:"string",value:"string"}),u(U,{label:"int",value:"int"}),u(U,{label:"float32",value:"float32"}),u(U,{label:"float64",value:"float64"}),u(U,{label:"bool",value:"bool"})])),_:2},1032,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(V,{modelValue:l.desc,"onUpdate:modelValue":e=>l.desc=e,placeholder:"备注 必填"},null,8,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(w,{icon:r(i),circle:"",size:"small",onClick:e[3]||(e[3]=l=>R(A.global))},null,8,["icon"])]),s("span",null,[u(w,{icon:r(m),circle:"",size:"small",onClick:l=>q(A.global,t)},null,8,["icon","onClick"])])])))),128))])),_:1})):c("",!0),u(f,{label:"使用Request"},{default:n((()=>[u(v,{modelValue:A.hasRequest,"onUpdate:modelValue":e[4]||(e[4]=l=>A.hasRequest=l)},null,8,["modelValue"])])),_:1}),A.hasRequest?(a(),t(f,{key:1,label:"Request"},{default:n((()=>[(a(!0),o(d,null,p(A.request,((l,t)=>(a(),o("div",{key:t,class:"plug-row"},[s("span",null,[u(V,{modelValue:l.key,"onUpdate:modelValue":e=>l.key=e,placeholder:"key 必填"},null,8,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(G,{modelValue:l.type,"onUpdate:modelValue":e=>l.type=e,placeholder:"type 必填"},{default:n((()=>[u(U,{label:"string",value:"string"}),u(U,{label:"int",value:"int"}),u(U,{label:"float32",value:"float32"}),u(U,{label:"float64",value:"float64"}),u(U,{label:"bool",value:"bool"})])),_:2},1032,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(V,{modelValue:l.desc,"onUpdate:modelValue":e=>l.desc=e,placeholder:"备注 必填"},null,8,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(w,{icon:r(i),circle:"",size:"small",onClick:e[5]||(e[5]=l=>R(A.request))},null,8,["icon"])]),s("span",null,[u(w,{icon:r(m),circle:"",size:"small",onClick:l=>q(A.request,t)},null,8,["icon","onClick"])])])))),128))])),_:1})):c("",!0),u(f,{label:"使用Response"},{default:n((()=>[u(v,{modelValue:A.hasResponse,"onUpdate:modelValue":e[6]||(e[6]=l=>A.hasResponse=l)},null,8,["modelValue"])])),_:1}),A.hasResponse?(a(),t(f,{key:2,label:"Response"},{default:n((()=>[(a(!0),o(d,null,p(A.response,((l,t)=>(a(),o("div",{key:t,class:"plug-row"},[s("span",null,[u(V,{modelValue:l.key,"onUpdate:modelValue":e=>l.key=e,placeholder:"key 必填"},null,8,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(G,{modelValue:l.type,"onUpdate:modelValue":e=>l.type=e,placeholder:"type 必填"},{default:n((()=>[u(U,{label:"string",value:"string"}),u(U,{label:"int",value:"int"}),u(U,{label:"float32",value:"float32"}),u(U,{label:"float64",value:"float64"}),u(U,{label:"bool",value:"bool"})])),_:2},1032,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(V,{modelValue:l.desc,"onUpdate:modelValue":e=>l.desc=e,placeholder:"备注 必填"},null,8,["modelValue","onUpdate:modelValue"])]),s("span",null,[u(w,{icon:r(i),circle:"",size:"small",onClick:e[7]||(e[7]=l=>R(A.response))},null,8,["icon"])]),s("span",null,[u(w,{icon:r(m),circle:"",size:"small",onClick:l=>q(A.response,t)},null,8,["icon","onClick"])])])))),128))])),_:1})):c("",!0),u(f,null,{default:n((()=>[u(w,{type:"primary",size:"small",onClick:_},{default:n((()=>[b("创建")])),_:1})])),_:1})])),_:1})])])}}},[["__scopeId","data-v-fc9f4086"]]);export{A as default};
