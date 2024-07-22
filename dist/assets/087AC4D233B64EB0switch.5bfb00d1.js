/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{x as e,cZ as a,U as i,y as t,bF as n,b5 as c,aw as l,bs as s,c0 as o,c_ as v,z as d,bW as u,ac as r,A as f,bX as p,cf as y,B as b,r as h,aa as m,L as k,b6 as I,X as x,o as g,c as V,b as S,D as T,C as w,e as _,g as B,w as C,p as A,q as E,h as N,t as P,d as z,aR as D,G as F,n as K,af as W,H as X,a8 as $,cW as j,at as q,I as G}from"./087AC4D233B64EB0index.39f1eb31.js";const H=e({modelValue:{type:[Boolean,String,Number],default:!1},disabled:{type:Boolean,default:!1},loading:{type:Boolean,default:!1},size:{type:String,validator:a},width:{type:[String,Number],default:""},inlinePrompt:{type:Boolean,default:!1},inactiveActionIcon:{type:i},activeActionIcon:{type:i},activeIcon:{type:i},inactiveIcon:{type:i},activeText:{type:String,default:""},inactiveText:{type:String,default:""},activeValue:{type:[Boolean,String,Number],default:!0},inactiveValue:{type:[Boolean,String,Number],default:!1},name:{type:String,default:""},validateEvent:{type:Boolean,default:!0},beforeChange:{type:t(Function)},id:String,tabindex:{type:[String,Number]},label:{type:String,default:void 0}}),L={[n]:e=>c(e)||l(e)||s(e),[o]:e=>c(e)||l(e)||s(e),[v]:e=>c(e)||l(e)||s(e)},R=["onClick"],U=["id","aria-checked","aria-disabled","aria-label","name","true-value","false-value","disabled","tabindex","onKeydown"],Z=["aria-hidden"],J=["aria-hidden"],M=["aria-hidden"],O="ElSwitch",Q=d({name:O});const Y=G(X(d({...Q,props:H,emits:L,setup(e,{expose:a,emit:i}){const t=e,{formItem:l}=u(),s=r(),d=f("switch"),{inputId:X}=p(t,{formItemContext:l}),G=y(b((()=>t.loading))),H=h(!1!==t.modelValue),L=h(),Q=h(),Y=b((()=>[d.b(),d.m(s.value),d.is("disabled",G.value),d.is("checked",ne.value)])),ee=b((()=>[d.e("label"),d.em("label","left"),d.is("active",!ne.value)])),ae=b((()=>[d.e("label"),d.em("label","right"),d.is("active",ne.value)])),ie=b((()=>({width:m(t.width)})));k((()=>t.modelValue),(()=>{H.value=!0}));const te=b((()=>!!H.value&&t.modelValue)),ne=b((()=>te.value===t.activeValue));[t.activeValue,t.inactiveValue].includes(te.value)||(i(n,t.inactiveValue),i(o,t.inactiveValue),i(v,t.inactiveValue)),k(ne,(e=>{var a;L.value.checked=e,t.validateEvent&&(null==(a=null==l?void 0:l.validate)||a.call(l,"change").catch((e=>I())))}));const ce=()=>{const e=ne.value?t.inactiveValue:t.activeValue;i(n,e),i(o,e),i(v,e),$((()=>{L.value.checked=ne.value}))},le=()=>{if(G.value)return;const{beforeChange:e}=t;if(!e)return void ce();const a=e();[j(a),c(a)].includes(!0)||q(O,"beforeChange must return type `Promise<boolean>` or `boolean`"),j(a)?a.then((e=>{e&&ce()})).catch((e=>{})):a&&ce()};return x((()=>{L.value.checked=ne.value})),a({focus:()=>{var e,a;null==(a=null==(e=L.value)?void 0:e.focus)||a.call(e)},checked:ne}),(e,a)=>(g(),V("div",{class:w(T(Y)),onClick:W(le,["prevent"])},[S("input",{id:T(X),ref_key:"input",ref:L,class:w(T(d).e("input")),type:"checkbox",role:"switch","aria-checked":T(ne),"aria-disabled":T(G),"aria-label":e.label,name:e.name,"true-value":e.activeValue,"false-value":e.inactiveValue,disabled:T(G),tabindex:e.tabindex,onChange:ce,onKeydown:_(le,["enter"])},null,42,U),e.inlinePrompt||!e.inactiveIcon&&!e.inactiveText?N("v-if",!0):(g(),V("span",{key:0,class:w(T(ee))},[e.inactiveIcon?(g(),B(T(E),{key:0},{default:C((()=>[(g(),B(A(e.inactiveIcon)))])),_:1})):N("v-if",!0),!e.inactiveIcon&&e.inactiveText?(g(),V("span",{key:1,"aria-hidden":T(ne)},P(e.inactiveText),9,Z)):N("v-if",!0)],2)),S("span",{ref_key:"core",ref:Q,class:w(T(d).e("core")),style:K(T(ie))},[e.inlinePrompt?(g(),V("div",{key:0,class:w(T(d).e("inner"))},[e.activeIcon||e.inactiveIcon?(g(),B(T(E),{key:0,class:w(T(d).is("icon"))},{default:C((()=>[(g(),B(A(T(ne)?e.activeIcon:e.inactiveIcon)))])),_:1},8,["class"])):e.activeText||e.inactiveText?(g(),V("span",{key:1,class:w(T(d).is("text")),"aria-hidden":!T(ne)},P(T(ne)?e.activeText:e.inactiveText),11,J)):N("v-if",!0)],2)):N("v-if",!0),S("div",{class:w(T(d).e("action"))},[e.loading?(g(),B(T(E),{key:0,class:w(T(d).is("loading"))},{default:C((()=>[z(T(D))])),_:1},8,["class"])):T(ne)?F(e.$slots,"active-action",{key:1},(()=>[e.activeActionIcon?(g(),B(T(E),{key:0},{default:C((()=>[(g(),B(A(e.activeActionIcon)))])),_:1})):N("v-if",!0)])):T(ne)?N("v-if",!0):F(e.$slots,"inactive-action",{key:2},(()=>[e.inactiveActionIcon?(g(),B(T(E),{key:0},{default:C((()=>[(g(),B(A(e.inactiveActionIcon)))])),_:1})):N("v-if",!0)]))],2)],6),e.inlinePrompt||!e.activeIcon&&!e.activeText?N("v-if",!0):(g(),V("span",{key:1,class:w(T(ae))},[e.activeIcon?(g(),B(T(E),{key:0},{default:C((()=>[(g(),B(A(e.activeIcon)))])),_:1})):N("v-if",!0),!e.activeIcon&&e.activeText?(g(),V("span",{key:1,"aria-hidden":!T(ne)},P(e.activeText),9,M)):N("v-if",!0)],2))],10,R))}}),[["__file","switch.vue"]]));export{Y as E};