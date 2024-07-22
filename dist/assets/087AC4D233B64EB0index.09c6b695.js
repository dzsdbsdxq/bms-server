/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{a1 as e,ao as t,H as n,z as a,A as l,o,g as s,w as u,G as r,ae as i,T as c,ap as d,aq as p,ar as v,B as m,as as h,x as f,U as b,Y as x,at as M,r as g,a9 as y,au as I,a as C,L as S,W as k,X as E,ab as O,av as w,aw as T,q as N,N as A,O as P,F as B,Z as z,ax as L,y as $,ay as j,az as W,aA as _,aB as q,aC as D,aD as R,aE as F,a8 as G,a3 as H,M as V,c as Z,b as U,C as X,i as Y,t as J,I as K,$ as Q}from"./087AC4D233B64EB0index.39f1eb31.js";import{_ as ee}from"./087AC4D233B64EB0index.bfb39b2a.js";import{E as te,C as ne}from"./087AC4D233B64EB0index.d8bfc0de.js";let ae=class{constructor(e,t){this.parent=e,this.domNode=t,this.subIndex=0,this.subIndex=0,this.init()}init(){this.subMenuItems=this.domNode.querySelectorAll("li"),this.addListeners()}gotoSubIndex(e){e===this.subMenuItems.length?e=0:e<0&&(e=this.subMenuItems.length-1),this.subMenuItems[e].focus(),this.subIndex=e}addListeners(){const n=this.parent.domNode;Array.prototype.forEach.call(this.subMenuItems,(a=>{a.addEventListener("keydown",(a=>{let l=!1;switch(a.code){case e.down:this.gotoSubIndex(this.subIndex+1),l=!0;break;case e.up:this.gotoSubIndex(this.subIndex-1),l=!0;break;case e.tab:t(n,"mouseleave");break;case e.enter:case e.space:l=!0,a.currentTarget.click()}return l&&(a.preventDefault(),a.stopPropagation()),!1}))}))}},le=class{constructor(e,t){this.domNode=e,this.submenu=null,this.submenu=null,this.init(t)}init(e){this.domNode.setAttribute("tabindex","0");const t=this.domNode.querySelector(".".concat(e,"-menu"));t&&(this.submenu=new ae(this,t)),this.addListeners()}addListeners(){this.domNode.addEventListener("keydown",(n=>{let a=!1;switch(n.code){case e.down:t(n.currentTarget,"mouseenter"),this.submenu&&this.submenu.gotoSubIndex(0),a=!0;break;case e.up:t(n.currentTarget,"mouseenter"),this.submenu&&this.submenu.gotoSubIndex(this.submenu.subMenuItems.length-1),a=!0;break;case e.tab:t(n.currentTarget,"mouseleave");break;case e.enter:case e.space:a=!0,n.currentTarget.click()}a&&n.preventDefault()}))}},oe=class{constructor(e,t){this.domNode=e,this.init(t)}init(e){const t=this.domNode.childNodes;Array.from(t).forEach((t=>{1===t.nodeType&&new le(t,e)}))}};var se=n(a({name:"ElMenuCollapseTransition",setup(){const e=l("menu");return{listeners:{onBeforeEnter:e=>e.style.opacity="0.2",onEnter(t,n){d(t,"".concat(e.namespace.value,"-opacity-transition")),t.style.opacity="1",n()},onAfterEnter(t){p(t,"".concat(e.namespace.value,"-opacity-transition")),t.style.opacity=""},onBeforeLeave(t){t.dataset||(t.dataset={}),v(t,e.m("collapse"))?(p(t,e.m("collapse")),t.dataset.oldOverflow=t.style.overflow,t.dataset.scrollWidth=t.clientWidth.toString(),d(t,e.m("collapse"))):(d(t,e.m("collapse")),t.dataset.oldOverflow=t.style.overflow,t.dataset.scrollWidth=t.clientWidth.toString(),p(t,e.m("collapse"))),t.style.width="".concat(t.scrollWidth,"px"),t.style.overflow="hidden"},onLeave(e){d(e,"horizontal-collapse-transition"),e.style.width="".concat(e.dataset.scrollWidth,"px")}}}}}),[["render",function(e,t,n,a,l,d){return o(),s(c,i({mode:"out-in"},e.listeners),{default:u((()=>[r(e.$slots,"default")])),_:3},16)}],["__file","menu-collapse-transition.vue"]]);function ue(e,t){const n=m((()=>{let n=e.parent;const a=[t.value];for(;"ElMenu"!==n.type.name;)n.props.index&&a.unshift(n.props.index),n=n.parent;return a}));return{parentMenu:m((()=>{let t=e.parent;for(;t&&!["ElMenu","ElSubMenu"].includes(t.type.name);)t=t.parent;return t})),indexPath:n}}function re(e){return m((()=>{const t=e.backgroundColor;return t?new h(t).shade(20).toString():""}))}const ie=(e,t)=>{const n=l("menu");return m((()=>n.cssVarBlock({"text-color":e.textColor||"","hover-text-color":e.textColor||"","bg-color":e.backgroundColor||"","hover-bg-color":re(e).value||"","active-color":e.activeTextColor||"",level:"".concat(t)})))},ce=f({index:{type:String,required:!0},showTimeout:Number,hideTimeout:Number,popperClass:String,disabled:Boolean,teleported:{type:Boolean,default:void 0},popperOffset:Number,expandCloseIcon:{type:b},expandOpenIcon:{type:b},collapseCloseIcon:{type:b},collapseOpenIcon:{type:b}}),de="ElSubMenu";var pe=a({name:de,props:ce,setup(e,{slots:t,expose:n}){const a=z(),{indexPath:o,parentMenu:s}=ue(a,m((()=>e.index))),u=l("menu"),r=l("sub-menu"),i=x("rootMenu");i||M(de,"can not inject root menu");const c=x("subMenu:".concat(s.value.uid));c||M(de,"can not inject sub menu");const d=g({}),p=g({});let v;const h=g(!1),f=g(),b=g(null),$=m((()=>"horizontal"===G.value&&W.value?"bottom-start":"right-start")),j=m((()=>"horizontal"===G.value&&W.value||"vertical"===G.value&&!i.props.collapse?e.expandCloseIcon&&e.expandOpenIcon?R.value?e.expandOpenIcon:e.expandCloseIcon:y:e.collapseCloseIcon&&e.collapseOpenIcon?R.value?e.collapseOpenIcon:e.collapseCloseIcon:I)),W=m((()=>0===c.level)),_=m((()=>{const t=e.teleported;return void 0===t?W.value:t})),q=m((()=>i.props.collapse?"".concat(u.namespace.value,"-zoom-in-left"):"".concat(u.namespace.value,"-zoom-in-top"))),D=m((()=>"horizontal"===G.value&&W.value?["bottom-start","bottom-end","top-start","top-end","right-start","left-start"]:["right-start","right","right-end","left-start","bottom-start","bottom-end","top-start","top-end"])),R=m((()=>i.openedMenus.includes(e.index))),F=m((()=>{let e=!1;return Object.values(d.value).forEach((t=>{t.active&&(e=!0)})),Object.values(p.value).forEach((t=>{t.active&&(e=!0)})),e})),G=m((()=>i.props.mode)),H=C({index:e.index,indexPath:o,active:F}),V=ie(i.props,c.level+1),Z=m((()=>{var t;return null!=(t=e.popperOffset)?t:i.props.popperOffset})),U=m((()=>{var t;return null!=(t=e.popperClass)?t:i.props.popperClass})),X=m((()=>{var t;return null!=(t=e.showTimeout)?t:i.props.showTimeout})),Y=m((()=>{var t;return null!=(t=e.hideTimeout)?t:i.props.hideTimeout})),J=e=>{var t,n,a;e||null==(a=null==(n=null==(t=b.value)?void 0:t.popperRef)?void 0:n.popperInstanceRef)||a.destroy()},K=()=>{"hover"===i.props.menuTrigger&&"horizontal"===i.props.mode||i.props.collapse&&"vertical"===i.props.mode||e.disabled||i.handleSubMenuClick({index:e.index,indexPath:o.value,active:F.value})},Q=(t,n=X.value)=>{var a;"focus"!==t.type&&("click"===i.props.menuTrigger&&"horizontal"===i.props.mode||!i.props.collapse&&"vertical"===i.props.mode||e.disabled?c.mouseInChild.value=!0:(c.mouseInChild.value=!0,null==v||v(),({stop:v}=L((()=>{i.openMenu(e.index,o.value)}),n)),_.value&&(null==(a=s.value.vnode.el)||a.dispatchEvent(new MouseEvent("mouseenter")))))},ne=(t=!1)=>{var n;"click"===i.props.menuTrigger&&"horizontal"===i.props.mode||!i.props.collapse&&"vertical"===i.props.mode?c.mouseInChild.value=!1:(null==v||v(),c.mouseInChild.value=!1,({stop:v}=L((()=>!h.value&&i.closeMenu(e.index,o.value)),Y.value)),_.value&&t&&(null==(n=c.handleMouseleave)||n.call(c,!0)))};S((()=>i.props.collapse),(e=>J(Boolean(e))));{const e=e=>{p.value[e.index]=e},t=e=>{delete p.value[e.index]};k("subMenu:".concat(a.uid),{addSubMenu:e,removeSubMenu:t,handleMouseleave:ne,mouseInChild:h,level:c.level+1})}return n({opened:R}),E((()=>{i.addSubMenu(H),c.addSubMenu(H)})),O((()=>{c.removeSubMenu(H),i.removeSubMenu(H)})),()=>{var n;const l=[null==(n=t.title)?void 0:n.call(t),w(N,{class:r.e("icon-arrow"),style:{transform:R.value?e.expandCloseIcon&&e.expandOpenIcon||e.collapseCloseIcon&&e.collapseOpenIcon&&i.props.collapse?"none":"rotateZ(180deg)":"none"}},{default:()=>T(j.value)?w(a.appContext.components[j.value]):w(j.value)})],o=i.isMenuPopup?w(te,{ref:b,visible:R.value,effect:"light",pure:!0,offset:Z.value,showArrow:!1,persistent:!0,popperClass:U.value,placement:$.value,teleported:_.value,fallbackPlacements:D.value,transition:q.value,gpuAcceleration:!1},{content:()=>{var e;return w("div",{class:[u.m(G.value),u.m("popup-container"),U.value],onMouseenter:e=>Q(e,100),onMouseleave:()=>ne(!0),onFocus:e=>Q(e,100)},[w("ul",{class:[u.b(),u.m("popup"),u.m("popup-".concat($.value))],style:V.value},[null==(e=t.default)?void 0:e.call(t)])])},default:()=>w("div",{class:r.e("title"),onClick:K},l)}):w(B,{},[w("div",{class:r.e("title"),ref:f,onClick:K},l),w(ee,{},{default:()=>{var e;return A(w("ul",{role:"menu",class:[u.b(),u.m("inline")],style:V.value},[null==(e=t.default)?void 0:e.call(t)]),[[P,R.value]])}})]);return w("li",{class:[r.b(),r.is("active",F.value),r.is("opened",R.value),r.is("disabled",e.disabled)],role:"menuitem",ariaHaspopup:!0,ariaExpanded:R.value,onMouseenter:Q,onMouseleave:()=>ne(),onFocus:Q},[o])}}});const ve=f({mode:{type:String,values:["horizontal","vertical"],default:"vertical"},defaultActive:{type:String,default:""},defaultOpeneds:{type:$(Array),default:()=>j([])},uniqueOpened:Boolean,router:Boolean,menuTrigger:{type:String,values:["hover","click"],default:"hover"},collapse:Boolean,backgroundColor:String,textColor:String,activeTextColor:String,closeOnClickOutside:Boolean,collapseTransition:{type:Boolean,default:!0},ellipsis:{type:Boolean,default:!0},popperOffset:{type:Number,default:6},ellipsisIcon:{type:b,default:()=>W},popperEffect:{type:String,values:["dark","light"],default:"dark"},popperClass:String,showTimeout:{type:Number,default:300},hideTimeout:{type:Number,default:300}}),me=e=>Array.isArray(e)&&e.every((e=>T(e)));var he=a({name:"ElMenu",props:ve,emits:{close:(e,t)=>T(e)&&me(t),open:(e,t)=>T(e)&&me(t),select:(e,t,n,a)=>T(e)&&me(t)&&R(n)&&(void 0===a||a instanceof Promise)},setup(e,{emit:t,slots:n,expose:a}){const o=z(),s=o.appContext.config.globalProperties.$router,u=g(),r=l("menu"),i=l("sub-menu"),c=g(-1),d=g(e.defaultOpeneds&&!e.collapse?e.defaultOpeneds.slice(0):[]),p=g(e.defaultActive),v=g({}),h=g({}),f=m((()=>"horizontal"===e.mode||"vertical"===e.mode&&e.collapse)),b=(n,a)=>{d.value.includes(n)||(e.uniqueOpened&&(d.value=d.value.filter((e=>a.includes(e)))),d.value.push(n),t("open",n,a))},x=e=>{const t=d.value.indexOf(e);-1!==t&&d.value.splice(t,1)},M=(e,n)=>{x(e),t("close",e,n)},y=({index:e,indexPath:t})=>{d.value.includes(e)?M(e,t):b(e,t)},I=n=>{("horizontal"===e.mode||e.collapse)&&(d.value=[]);const{index:a,indexPath:l}=n;if(!F(a)&&!F(l))if(e.router&&s){const e=n.route||a,o=s.push(e).then((e=>(e||(p.value=a),e)));t("select",a,l,{index:a,indexPath:l,route:e},o)}else p.value=a,t("select",a,l,{index:a,indexPath:l})},O=()=>{var e,t;if(!u.value)return-1;const n=Array.from(null!=(t=null==(e=u.value)?void 0:e.childNodes)?t:[]).filter((e=>"#comment"!==e.nodeName&&("#text"!==e.nodeName||e.nodeValue))),a=getComputedStyle(u.value),l=Number.parseInt(a.paddingLeft,10),o=Number.parseInt(a.paddingRight,10),s=u.value.clientWidth-l-o;let r=0,i=0;return n.forEach(((e,t)=>{r+=(e=>{const t=getComputedStyle(e),n=Number.parseInt(t.marginLeft,10),a=Number.parseInt(t.marginRight,10);return e.offsetWidth+n+a||0})(e),r<=s-64&&(i=t+1)})),i===n.length?-1:i};let T=!0;const P=()=>{if(c.value===O())return;const e=()=>{c.value=-1,G((()=>{c.value=O()}))};T?e():((e,t=33.34)=>{let n;return()=>{n&&clearTimeout(n),n=setTimeout((()=>{e()}),t)}})(e)(),T=!1};let B;S((()=>e.defaultActive),(t=>{v.value[t]||(p.value=""),(t=>{const n=v.value,a=n[t]||p.value&&n[p.value]||n[e.defaultActive];p.value=a?a.index:t})(t)})),S((()=>e.collapse),(e=>{e&&(d.value=[])})),S(v.value,(()=>{const t=p.value&&v.value[p.value];if(!t||"horizontal"===e.mode||e.collapse)return;t.indexPath.forEach((e=>{const t=h.value[e];t&&b(e,t.indexPath)}))})),_((()=>{"horizontal"===e.mode&&e.ellipsis?B=q(u,P).stop:null==B||B()}));const L=g(!1);{const t=e=>{h.value[e.index]=e},n=e=>{delete h.value[e.index]},a=e=>{v.value[e.index]=e},l=e=>{delete v.value[e.index]};k("rootMenu",C({props:e,openedMenus:d,items:v,subMenus:h,activeIndex:p,isMenuPopup:f,addMenuItem:a,removeMenuItem:l,addSubMenu:t,removeSubMenu:n,openMenu:b,closeMenu:M,handleMenuItemClick:I,handleSubMenuClick:y})),k("subMenu:".concat(o.uid),{addSubMenu:t,removeSubMenu:n,mouseInChild:L,level:0})}E((()=>{"horizontal"===e.mode&&new oe(o.vnode.el,r.namespace.value)}));a({open:e=>{const{indexPath:t}=h.value[e];t.forEach((e=>b(e,t)))},close:x,handleResize:P});return()=>{var a,l;let o=null!=(l=null==(a=n.default)?void 0:a.call(n))?l:[];const s=[];if("horizontal"===e.mode&&u.value){const t=D(o),n=-1===c.value?t:t.slice(0,c.value),a=-1===c.value?[]:t.slice(c.value);(null==a?void 0:a.length)&&e.ellipsis&&(o=n,s.push(w(pe,{index:"sub-menu-more",class:i.e("hide-arrow"),popperOffset:e.popperOffset},{title:()=>w(N,{class:i.e("icon-more")},{default:()=>w(e.ellipsisIcon)}),default:()=>a})))}const p=ie(e,0),v=e.closeOnClickOutside?[[ne,()=>{d.value.length&&(L.value||(d.value.forEach((e=>{return t("close",e,(n=e,h.value[n].indexPath));var n})),d.value=[]))}]]:[],m=A(w("ul",{key:String(e.collapse),role:"menubar",ref:u,style:p.value,class:{[r.b()]:!0,[r.m(e.mode)]:!0,[r.m("collapse")]:e.collapse}},[...o,...s]),v);return e.collapseTransition&&"vertical"===e.mode?w(se,(()=>m)):m}}});const fe=f({index:{type:$([String,null]),default:null},route:{type:$([String,Object])},disabled:Boolean}),be="ElMenuItem";var xe=n(a({name:be,components:{ElTooltip:te},props:fe,emits:{click:e=>T(e.index)&&Array.isArray(e.indexPath)},setup(e,{emit:t}){const n=z(),a=x("rootMenu"),o=l("menu"),s=l("menu-item");a||M(be,"can not inject root menu");const{parentMenu:u,indexPath:r}=ue(n,H(e,"index")),i=x("subMenu:".concat(u.value.uid));i||M(be,"can not inject sub menu");const c=m((()=>e.index===a.activeIndex)),d=C({index:e.index,indexPath:r,active:c});return E((()=>{i.addSubMenu(d),a.addMenuItem(d)})),O((()=>{i.removeSubMenu(d),a.removeMenuItem(d)})),{parentMenu:u,rootMenu:a,active:c,nsMenu:o,nsMenuItem:s,handleClick:()=>{e.disabled||(a.handleMenuItemClick({index:e.index,indexPath:r.value,route:e.route}),t("click",d))}}}}),[["render",function(e,t,n,a,l,i){const c=V("el-tooltip");return o(),Z("li",{class:X([e.nsMenuItem.b(),e.nsMenuItem.is("active",e.active),e.nsMenuItem.is("disabled",e.disabled)]),role:"menuitem",tabindex:"-1",onClick:t[0]||(t[0]=(...t)=>e.handleClick&&e.handleClick(...t))},["ElMenu"===e.parentMenu.type.name&&e.rootMenu.props.collapse&&e.$slots.title?(o(),s(c,{key:0,effect:e.rootMenu.props.popperEffect,placement:"right","fallback-placements":["left"],persistent:""},{content:u((()=>[r(e.$slots,"title")])),default:u((()=>[U("div",{class:X(e.nsMenu.be("tooltip","trigger"))},[r(e.$slots,"default")],2)])),_:3},8,["effect"])):(o(),Z(B,{key:1},[r(e.$slots,"default"),r(e.$slots,"title")],64))],2)}],["__file","menu-item.vue"]]);var Me=n(a({name:"ElMenuItemGroup",props:{title:String},setup:()=>({ns:l("menu-item-group")})}),[["render",function(e,t,n,a,l,s){return o(),Z("li",{class:X(e.ns.b())},[U("div",{class:X(e.ns.e("title"))},[e.$slots.title?r(e.$slots,"title",{key:1}):(o(),Z(B,{key:0},[Y(J(e.title),1)],64))],2),U("ul",null,[r(e.$slots,"default")])],2)}],["__file","menu-item-group.vue"]]);const ge=K(he,{MenuItem:xe,MenuItemGroup:Me,SubMenu:pe}),ye=Q(xe);Q(Me);const Ie=Q(pe);export{ge as E,Ie as a,ye as b};