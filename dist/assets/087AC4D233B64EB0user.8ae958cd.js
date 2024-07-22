/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{r as e,L as a,o as l,c as t,b as i,d as s,w as o,i as r,g as u,h as d,d1 as n,P as m,d2 as p,E as c,d3 as v,d4 as g,d5 as h,a8 as y,d6 as B,k as b,j as f}from"./087AC4D233B64EB0index.39f1eb31.js";import{E as w}from"./087AC4D233B64EB0drawer.ae3d541f.js";import"./087AC4D233B64EB0overlay.3905897c.js";import{a as C,E as I}from"./087AC4D233B64EB0form-item.a86c41d2.js";/* empty css                              */import{E as D}from"./087AC4D233B64EB0pagination.ce43b352.js";import"./087AC4D233B64EB0tag.d3c1c3e4.js";import"./087AC4D233B64EB0select.7f5b5538.js";import"./087AC4D233B64EB0scrollbar.3894915f.js";/* empty css                               */import{E,a as j}from"./087AC4D233B64EB0table-column.e9b4a870.js";import"./087AC4D233B64EB0checkbox.a3853383.js";/* empty css                                */import{E as A}from"./087AC4D233B64EB0switch.5bfb00d1.js";import{E as k}from"./087AC4D233B64EB0radio.6c10d629.js";/* empty css                               */import{g as V}from"./087AC4D233B64EB0authority.e5789bef.js";import{C as _}from"./087AC4D233B64EB0index.142f6335.js";import{_ as z}from"./087AC4D233B64EB0index.0dabe5ab.js";import"./087AC4D233B64EB0isUndefined.d281c39a.js";import"./087AC4D233B64EB0castArray.209ac14b.js";import"./087AC4D233B64EB0_baseClone.a73122b0.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0_initCloneObject.0c0ca3ea.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";import"./087AC4D233B64EB0index.d8bfc0de.js";import"./087AC4D233B64EB0index.5d58b83c.js";import"./087AC4D233B64EB0strings.2df7e9e1.js";import"./087AC4D233B64EB0debounce.0aab18a9.js";import"./087AC4D233B64EB0index.8fd65b4f.js";import"./087AC4D233B64EB0cloneDeep.5b0b6a51.js";import"./087AC4D233B64EB0image-viewer.9c6934e5.js";import"./087AC4D233B64EB0common.c541c571.js";import"./087AC4D233B64EB0upload.099811f5.js";import"./087AC4D233B64EB0progress.65ec04e5.js";const x={class:"gva-table-box"},N={class:"gva-btn-list"},U={class:"gva-pagination"},S={class:"flex items-center justify-between"},O=i("span",{class:"text-lg"},"用户",-1),T=["src"],q={key:1,class:"header-img-box"},J=Object.assign({name:"User"},{__name:"user",setup(J){const F=e("https://www.sztv.com.cn/ysztv/web//"),P=(e,a)=>{e&&e.forEach((e=>{if(e.children&&e.children.length){const l={authorityId:e.authorityId,authorityName:e.authorityName,children:[]};P(e.children,l.children),a.push(l)}else{const l={authorityId:e.authorityId,authorityName:e.authorityName};a.push(l)}}))},$=e(1),R=e(0),G=e(10),H=e([]),K=e=>{G.value=e,Z()},L=e=>{$.value=e,Z()},Z=async()=>{const e=await n({page:$.value,pageSize:G.value});0===e.code&&(H.value=e.data.list,R.value=e.data.total,$.value=e.data.page,G.value=e.data.pageSize)};a((()=>H.value),(()=>{M()}));(async()=>{Z();const e=await V({page:1,pageSize:999});Y(e.data.list)})();const M=()=>{H.value&&H.value.forEach((e=>{e.authorityIds=e.authorities&&e.authorities.map((e=>e.authorityId))}))},Q=e(null),W=()=>{Q.value.open()},X=e([]),Y=e=>{X.value=[],P(e,X.value)},ee=e({username:"",password:"",nickName:"",headerImg:"",authorityId:"",authorityIds:[],enable:1}),ae=e({userName:[{required:!0,message:"请输入用户名",trigger:"blur"},{min:5,message:"最低5位字符",trigger:"blur"}],password:[{required:!0,message:"请输入用户密码",trigger:"blur"},{min:6,message:"最低6位字符",trigger:"blur"}],nickName:[{required:!0,message:"请输入用户昵称",trigger:"blur"}],phone:[{pattern:/^1([38][0-9]|4[014-9]|[59][0-35-9]|6[2567]|7[0-8])\d{8}$/,message:"请输入合法手机号",trigger:"blur"}],email:[{pattern:/^([0-9A-Za-z\-_.]+)@([0-9a-z]+\.[a-z]{2,3}(\.[a-z]{2})?)$/g,message:"请输入正确的邮箱",trigger:"blur"}],authorityId:[{required:!0,message:"请选择用户角色",trigger:"blur"}]}),le=e(null),te=async()=>{ee.value.authorityId=ee.value.authorityIds[0],le.value.validate((async e=>{if(e){const e={...ee.value};if("add"===oe.value){0===(await g(e)).code&&(c({type:"success",message:"创建成功"}),await Z(),se())}if("edit"===oe.value){0===(await h(e)).code&&(c({type:"success",message:"编辑成功"}),await Z(),se())}}}))},ie=e(!1),se=()=>{le.value.resetFields(),ee.value.headerImg="",ee.value.authorityIds=[],ie.value=!1},oe=e("add"),re=()=>{oe.value="add",ie.value=!0},ue={},de=async(e,a,l)=>{if(a)return void(l||(ue[e.ID]=[...e.authorityIds]));await y();0===(await B({ID:e.ID,authorityIds:e.authorityIds})).code?c({type:"success",message:"角色设置成功"}):l?e.authorityIds=[l,...e.authorityIds]:(e.authorityIds=[...ue[e.ID]],delete ue[e.ID])};return(e,a)=>{const n=b,g=E,B=k,V=A,J=j,P=D,M=f,Y=C,ue=I,ne=w;return l(),t("div",null,[i("div",x,[i("div",N,[s(n,{type:"primary",icon:"plus",size:"small",onClick:re},{default:o((()=>[r("新增用户")])),_:1})]),s(J,{data:H.value,border:!1,size:"small","row-key":"ID"},{default:o((()=>[s(g,{align:"left",label:"头像","min-width":"75"},{default:o((e=>[s(_,{style:{"margin-top":"8px"},"pic-src":e.row.headerImg},null,8,["pic-src"])])),_:1}),s(g,{align:"left",label:"ID","min-width":"50",prop:"ID"}),s(g,{align:"left",label:"用户名","min-width":"150",prop:"userName"}),s(g,{align:"left",label:"昵称","min-width":"150",prop:"nickName"}),s(g,{align:"left",label:"手机号","min-width":"180",prop:"phone"}),s(g,{align:"left",label:"邮箱","min-width":"180",prop:"email"}),s(g,{align:"left",label:"用户角色","min-width":"200"},{default:o((e=>[s(B,{modelValue:e.row.authorityIds,"onUpdate:modelValue":a=>e.row.authorityIds=a,options:X.value,size:"small","show-all-levels":!1,"collapse-tags":"",props:{multiple:!0,checkStrictly:!0,label:"authorityName",value:"authorityId",disabled:"disabled",emitPath:!1},clearable:!1,onVisibleChange:a=>{de(e.row,a,0)},onRemoveTag:a=>{de(e.row,!1,a)}},null,8,["modelValue","onUpdate:modelValue","options","onVisibleChange","onRemoveTag"])])),_:1}),s(g,{align:"left",label:"启用","min-width":"150"},{default:o((e=>[s(V,{modelValue:e.row.enable,"onUpdate:modelValue":a=>e.row.enable=a,"inline-prompt":"","active-value":1,"inactive-value":2,onChange:()=>{(async e=>{ee.value=JSON.parse(JSON.stringify(e)),await y();const a={...ee.value};0===(await h(a)).code&&(c({type:"success",message:"".concat(2===a.enable?"禁用":"启用","成功")}),await Z(),ee.value.headerImg="",ee.value.authorityIds=[])})(e.row)}},null,8,["modelValue","onUpdate:modelValue","onChange"])])),_:1}),s(g,{label:"操作","min-width":"250",fixed:"right"},{default:o((e=>[s(n,{type:"primary",link:"",size:"small",icon:"delete",onClick:a=>(async e=>{m.confirm("确定要删除吗?","提示",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{0===(await v({id:e.ID})).code&&(c.success("删除成功"),await Z())}))})(e.row)},{default:o((()=>[r("删除")])),_:2},1032,["onClick"]),s(n,{type:"primary",link:"",icon:"edit",size:"small",onClick:a=>{return l=e.row,oe.value="edit",ee.value=JSON.parse(JSON.stringify(l)),void(ie.value=!0);var l}},{default:o((()=>[r("编辑")])),_:2},1032,["onClick"]),s(n,{type:"primary",link:"",icon:"magic-stick",size:"small",onClick:a=>{return l=e.row,void m.confirm("是否将此用户密码重置为123456?","警告",{confirmButtonText:"确定",cancelButtonText:"取消",type:"warning"}).then((async()=>{const e=await p({ID:l.ID});0===e.code?c({type:"success",message:e.msg}):c({type:"error",message:e.msg})}));var l}},{default:o((()=>[r("重置密码")])),_:2},1032,["onClick"])])),_:1})])),_:1},8,["data"]),i("div",U,[s(P,{"current-page":$.value,"page-size":G.value,"page-sizes":[10,30,50,100],total:R.value,small:"",layout:"total, sizes, prev, pager, next",onCurrentChange:L,onSizeChange:K},null,8,["current-page","page-size","total"])])]),s(ne,{modelValue:ie.value,"onUpdate:modelValue":a[7]||(a[7]=e=>ie.value=e),size:"60%","show-close":!1,"close-on-press-escape":!1,"close-on-click-modal":!1},{header:o((()=>[i("div",S,[O,i("div",null,[s(n,{size:"small",onClick:se},{default:o((()=>[r("取 消")])),_:1}),s(n,{size:"small",type:"primary",onClick:te},{default:o((()=>[r("确 定")])),_:1})])])])),default:o((()=>[s(ue,{ref_key:"userForm",ref:le,rules:ae.value,model:ee.value,"label-width":"80px"},{default:o((()=>["add"===oe.value?(l(),u(Y,{key:0,label:"用户名",prop:"userName"},{default:o((()=>[s(M,{modelValue:ee.value.userName,"onUpdate:modelValue":a[0]||(a[0]=e=>ee.value.userName=e)},null,8,["modelValue"])])),_:1})):d("",!0),"add"===oe.value?(l(),u(Y,{key:1,label:"密码",prop:"password"},{default:o((()=>[s(M,{modelValue:ee.value.password,"onUpdate:modelValue":a[1]||(a[1]=e=>ee.value.password=e)},null,8,["modelValue"])])),_:1})):d("",!0),s(Y,{label:"昵称",prop:"nickName"},{default:o((()=>[s(M,{modelValue:ee.value.nickName,"onUpdate:modelValue":a[2]||(a[2]=e=>ee.value.nickName=e)},null,8,["modelValue"])])),_:1}),s(Y,{label:"手机号",prop:"phone"},{default:o((()=>[s(M,{modelValue:ee.value.phone,"onUpdate:modelValue":a[3]||(a[3]=e=>ee.value.phone=e)},null,8,["modelValue"])])),_:1}),s(Y,{label:"邮箱",prop:"email"},{default:o((()=>[s(M,{modelValue:ee.value.email,"onUpdate:modelValue":a[4]||(a[4]=e=>ee.value.email=e)},null,8,["modelValue"])])),_:1}),s(Y,{label:"用户角色",prop:"authorityId"},{default:o((()=>[s(B,{modelValue:ee.value.authorityIds,"onUpdate:modelValue":a[5]||(a[5]=e=>ee.value.authorityIds=e),style:{width:"100%"},size:"small",options:X.value,"show-all-levels":!1,props:{multiple:!0,checkStrictly:!0,label:"authorityName",value:"authorityId",disabled:"disabled",emitPath:!1},clearable:!1},null,8,["modelValue","options"])])),_:1}),s(Y,{label:"启用",prop:"disabled"},{default:o((()=>[s(V,{modelValue:ee.value.enable,"onUpdate:modelValue":a[6]||(a[6]=e=>ee.value.enable=e),"inline-prompt":"","active-value":1,"inactive-value":2},null,8,["modelValue"])])),_:1}),s(Y,{label:"头像","label-width":"80px"},{default:o((()=>[i("div",{style:{display:"inline-block"},onClick:W},[ee.value.headerImg?(l(),t("img",{key:0,alt:"头像",class:"header-img-box",src:ee.value.headerImg&&"http"!==ee.value.headerImg.slice(0,4)?F.value+ee.value.headerImg:ee.value.headerImg},null,8,T)):(l(),t("div",q,"从媒体库选择")),s(z,{ref_key:"chooseImg",ref:Q,target:ee.value,"target-key":"headerImg"},null,8,["target"])])])),_:1})])),_:1},8,["rules","model"])])),_:1},8,["modelValue"])])}}});export{J as default};