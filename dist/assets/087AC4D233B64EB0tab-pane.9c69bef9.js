/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{aU as e,aC as a,cN as t,x as l,y as s,ay as o,z as n,Y as r,at as i,A as u,r as c,L as d,a8 as v,aB as b,o as p,c as f,C as m,D as h,n as y,H as g,Z as C,cO as x,cP as P,B as w,X as B,aV as T,d as k,q as N,bO as S,au as A,bx as E,a1 as R,W as F,G as L,aI as _,bF as K,bZ as O,aw as V,bs as q,a0 as z,cQ as D,a as j,am as H,N as I,O as M,h as W,I as X,$ as Y}from"./087AC4D233B64EB0index.39f1eb31.js";import{c as Z}from"./087AC4D233B64EB0strings.2df7e9e1.js";const $=(l,s)=>{const o={},n=e([]);return{children:n,addChild:e=>{o[e.uid]=e,n.value=((e,l,s)=>a(e.subTree).filter((e=>{var a;return t(e)&&(null==(a=e.type)?void 0:a.name)===l&&!!e.component})).map((e=>e.component.uid)).map((e=>s[e])).filter((e=>!!e)))(l,s,o)},removeChild:e=>{delete o[e],n.value=n.value.filter((a=>a.uid!==e))}}},G=Symbol("tabsRootContextKey"),Q=l({tabs:{type:s(Array),default:()=>o([])}}),U="ElTabBar",J=n({name:U});var ee=g(n({...J,props:Q,setup(e,{expose:a}){const t=e,l=C(),s=r(G);s||i(U,"<el-tabs><el-tab-bar /></el-tabs>");const o=u("tabs"),n=c(),g=c(),x=()=>g.value=(()=>{let e=0,a=0;const o=["top","bottom"].includes(s.props.tabPosition)?"width":"height",n="width"===o?"x":"y",r="x"===n?"left":"top";return t.tabs.every((s=>{var n,i;const u=null==(i=null==(n=l.parent)?void 0:n.refs)?void 0:i["tab-".concat(s.uid)];if(!u)return!1;if(!s.active)return!0;e=u["offset".concat(Z(r))],a=u["client".concat(Z(o))];const c=window.getComputedStyle(u);return"width"===o&&(t.tabs.length>1&&(a-=Number.parseFloat(c.paddingLeft)+Number.parseFloat(c.paddingRight)),e+=Number.parseFloat(c.paddingLeft)),!1})),{[o]:"".concat(a,"px"),transform:"translate".concat(Z(n),"(").concat(e,"px)")}})();return d((()=>t.tabs),(async()=>{await v(),x()}),{immediate:!0}),b(n,(()=>x())),a({ref:n,update:x}),(e,a)=>(p(),f("div",{ref_key:"barRef",ref:n,class:m([h(o).e("active-bar"),h(o).is(h(s).props.tabPosition)]),style:y(g.value)},null,6))}}),[["__file","tab-bar.vue"]]);const ae=l({panes:{type:s(Array),default:()=>o([])},currentName:{type:[String,Number],default:""},editable:Boolean,type:{type:String,values:["card","border-card",""],default:""},stretch:Boolean}),te="ElTabNav",le=n({name:te,props:ae,emits:{tabClick:(e,a,t)=>t instanceof Event,tabRemove:(e,a)=>a instanceof Event},setup(e,{expose:a,emit:t}){const l=C(),s=r(G);s||i(te,"<el-tabs><tab-nav /></el-tabs>");const o=u("tabs"),n=x(),p=P(),f=c(),m=c(),h=c(),y=c(),g=c(!1),F=c(0),L=c(!1),_=c(!0),K=w((()=>["top","bottom"].includes(s.props.tabPosition)?"width":"height")),O=w((()=>{const e="width"===K.value?"X":"Y";return{transform:"translate".concat(e,"(-").concat(F.value,"px)")}})),V=()=>{if(!f.value)return;const e=f.value["offset".concat(Z(K.value))],a=F.value;if(!a)return;const t=a>e?a-e:0;F.value=t},q=()=>{if(!f.value||!m.value)return;const e=m.value["offset".concat(Z(K.value))],a=f.value["offset".concat(Z(K.value))],t=F.value;if(e-t<=a)return;const l=e-t>2*a?t+a:e-a;F.value=l},z=async()=>{const e=m.value;if(!(g.value&&h.value&&f.value&&e))return;await v();const a=h.value.querySelector(".is-active");if(!a)return;const t=f.value,l=["top","bottom"].includes(s.props.tabPosition),o=a.getBoundingClientRect(),n=t.getBoundingClientRect(),r=l?e.offsetWidth-n.width:e.offsetHeight-n.height,i=F.value;let u=i;l?(o.left<n.left&&(u=i-(n.left-o.left)),o.right>n.right&&(u=i+o.right-n.right)):(o.top<n.top&&(u=i-(n.top-o.top)),o.bottom>n.bottom&&(u=i+(o.bottom-n.bottom))),u=Math.max(u,0),F.value=Math.min(u,r)},D=()=>{var a;if(!m.value||!f.value)return;e.stretch&&(null==(a=y.value)||a.update());const t=m.value["offset".concat(Z(K.value))],l=f.value["offset".concat(Z(K.value))],s=F.value;l<t?(g.value=g.value||{},g.value.prev=s,g.value.next=s+l<t,t-s<l&&(F.value=t-l)):(g.value=!1,s>0&&(F.value=0))},j=e=>{const a=e.code,{up:t,down:l,left:s,right:o}=R;if(![t,l,s,o].includes(a))return;const n=Array.from(e.currentTarget.querySelectorAll("[role=tab]:not(.is-disabled)")),r=n.indexOf(e.target);let i;i=a===s||a===t?0===r?n.length-1:r-1:r<n.length-1?r+1:0,n[i].focus({preventScroll:!0}),n[i].click(),H()},H=()=>{_.value&&(L.value=!0)},I=()=>L.value=!1;return d(n,(e=>{"hidden"===e?_.value=!1:"visible"===e&&setTimeout((()=>_.value=!0),50)})),d(p,(e=>{e?setTimeout((()=>_.value=!0),50):_.value=!1})),b(h,D),B((()=>setTimeout((()=>z()),0))),T((()=>D())),a({scrollToActiveTab:z,removeFocus:I}),d((()=>e.panes),(()=>l.update()),{flush:"post",deep:!0}),()=>{const a=g.value?[k("span",{class:[o.e("nav-prev"),o.is("disabled",!g.value.prev)],onClick:V},[k(N,null,{default:()=>[k(S,null,null)]})]),k("span",{class:[o.e("nav-next"),o.is("disabled",!g.value.next)],onClick:q},[k(N,null,{default:()=>[k(A,null,null)]})])]:null,l=e.panes.map(((a,l)=>{var n,r,i,u;const c=a.uid,d=a.props.disabled,v=null!=(r=null!=(n=a.props.name)?n:a.index)?r:"".concat(l),b=!d&&(a.isClosable||e.editable);a.index="".concat(l);const p=b?k(N,{class:"is-icon-close",onClick:e=>t("tabRemove",a,e)},{default:()=>[k(E,null,null)]}):null,f=(null==(u=(i=a.slots).label)?void 0:u.call(i))||a.props.label,m=!d&&a.active?0:-1;return k("div",{ref:"tab-".concat(c),class:[o.e("item"),o.is(s.props.tabPosition),o.is("active",a.active),o.is("disabled",d),o.is("closable",b),o.is("focus",L.value)],id:"tab-".concat(v),key:"tab-".concat(c),"aria-controls":"pane-".concat(v),role:"tab","aria-selected":a.active,tabindex:m,onFocus:()=>H(),onBlur:()=>I(),onClick:e=>{I(),t("tabClick",a,v,e)},onKeydown:e=>{!b||e.code!==R.delete&&e.code!==R.backspace||t("tabRemove",a,e)}},[f,p])}));return k("div",{ref:h,class:[o.e("nav-wrap"),o.is("scrollable",!!g.value),o.is(s.props.tabPosition)]},[a,k("div",{class:o.e("nav-scroll"),ref:f},[k("div",{class:[o.e("nav"),o.is(s.props.tabPosition),o.is("stretch",e.stretch&&["top","bottom"].includes(s.props.tabPosition))],ref:m,style:O.value,role:"tablist",onKeydown:j},[e.type?null:k(ee,{ref:y,tabs:[...e.panes]},null),l])])])}}}),se=l({type:{type:String,values:["card","border-card",""],default:""},closable:Boolean,addable:Boolean,modelValue:{type:[String,Number]},editable:Boolean,tabPosition:{type:String,values:["top","right","bottom","left"],default:"top"},beforeLeave:{type:s(Function),default:()=>!0},stretch:Boolean}),oe=e=>V(e)||q(e),ne=n({name:"ElTabs",props:se,emits:{[K]:e=>oe(e),tabClick:(e,a)=>a instanceof Event,tabChange:e=>oe(e),edit:(e,a)=>["remove","add"].includes(a),tabRemove:e=>oe(e),tabAdd:()=>!0},setup(e,{emit:a,slots:t,expose:l}){var s;const o=u("tabs"),{children:n,addChild:r,removeChild:i}=$(C(),"ElTabPane"),b=c(),p=c(null!=(s=e.modelValue)?s:"0"),f=async(t,l=!1)=>{var s,o,n;if(p.value!==t&&!O(t))try{!1!==await(null==(s=e.beforeLeave)?void 0:s.call(e,t,p.value))&&(p.value=t,l&&(a(K,t),a("tabChange",t)),null==(n=null==(o=b.value)?void 0:o.removeFocus)||n.call(o))}catch(r){}},m=(e,t,l)=>{e.props.disabled||(f(t,!0),a("tabClick",e,l))},h=(e,t)=>{e.props.disabled||O(e.props.name)||(t.stopPropagation(),a("edit",e.props.name,"remove"),a("tabRemove",e.props.name))},y=()=>{a("edit",void 0,"add"),a("tabAdd")};return d((()=>e.modelValue),(e=>f(e))),d(p,(async()=>{var e;await v(),null==(e=b.value)||e.scrollToActiveTab()})),F(G,{props:e,currentName:p,registerPane:r,unregisterPane:i}),l({currentName:p}),()=>{const a=t["add-icon"],l=e.editable||e.addable?k("span",{class:o.e("new-tab"),tabindex:"0",onClick:y,onKeydown:e=>{e.code===R.enter&&y()}},[a?L(t,"add-icon"):k(N,{class:o.is("icon-plus")},{default:()=>[k(_,null,null)]})]):null,s=k("div",{class:[o.e("header"),o.is(e.tabPosition)]},[l,k(le,{ref:b,currentName:p.value,editable:e.editable,type:e.type,panes:n.value,stretch:e.stretch,onTabClick:m,onTabRemove:h},null)]),r=k("div",{class:o.e("content")},[L(t,"default")]);return k("div",{class:[o.b(),o.m(e.tabPosition),{[o.m("card")]:"card"===e.type,[o.m("border-card")]:"border-card"===e.type}]},[..."bottom"!==e.tabPosition?[s,r]:[r,s]])}}}),re=l({label:{type:String,default:""},name:{type:[String,Number]},closable:Boolean,disabled:Boolean,lazy:Boolean}),ie=["id","aria-hidden","aria-labelledby"],ue="ElTabPane",ce=n({name:ue});var de=g(n({...ce,props:re,setup(e){const a=e,t=C(),l=z(),s=r(G);s||i(ue,"usage: <el-tabs><el-tab-pane /></el-tabs/>");const o=u("tab-pane"),n=c(),v=w((()=>a.closable||s.props.closable)),b=D((()=>{var e;return s.currentName.value===(null!=(e=a.name)?e:n.value)})),y=c(b.value),g=w((()=>{var e;return null!=(e=a.name)?e:n.value})),x=D((()=>!a.lazy||y.value||b.value));d(b,(e=>{e&&(y.value=!0)}));const P=j({uid:t.uid,slots:l,props:a,paneName:g,active:b,index:n,isClosable:v});return B((()=>{s.registerPane(P)})),H((()=>{s.unregisterPane(P.uid)})),(e,a)=>h(x)?I((p(),f("div",{key:0,id:"pane-".concat(h(g)),class:m(h(o).b()),role:"tabpanel","aria-hidden":!h(b),"aria-labelledby":"tab-".concat(h(g))},[L(e.$slots,"default")],10,ie)),[[M,h(b)]]):W("v-if",!0)}}),[["__file","tab-pane.vue"]]);const ve=X(ne,{TabPane:de}),be=Y(de);export{be as E,ve as a};
