/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{b as e,u as t,E as o}from"./087AC4D233B64EB0index.d8bfc0de.js";import{x as a,b5 as r,z as s,B as p,A as i,r as l,D as n,aa as d,o as f,g as b,w as c,c as u,C as v,t as m,h,G as g,i as w,ae as y,H as B,d0 as x,I as A}from"./087AC4D233B64EB0index.39f1eb31.js";import{d as C}from"./087AC4D233B64EB0dropdown.9d3edc17.js";const S=a({trigger:e.trigger,placement:C.placement,disabled:e.disabled,visible:t.visible,transition:t.transition,popperOptions:C.popperOptions,tabindex:C.tabindex,content:t.content,popperStyle:t.popperStyle,popperClass:t.popperClass,enterable:{...t.enterable,default:!0},effect:{...t.effect,default:"light"},teleported:t.teleported,title:String,width:{type:[String,Number],default:150},offset:{type:Number,default:void 0},showAfter:{type:Number,default:0},hideAfter:{type:Number,default:200},autoClose:{type:Number,default:0},showArrow:{type:Boolean,default:!0},persistent:{type:Boolean,default:!0},"onUpdate:visible":{type:Function}}),E={"update:visible":e=>r(e),"before-enter":()=>!0,"before-leave":()=>!0,"after-enter":()=>!0,"after-leave":()=>!0},N=s({name:"ElPopover"}),R=s({...N,props:S,emits:E,setup(e,{expose:t,emit:a}){const r=e,s=p((()=>r["onUpdate:visible"])),B=i("popover"),x=l(),A=p((()=>{var e;return null==(e=n(x))?void 0:e.popperRef})),C=p((()=>[{width:d(r.width)},r.popperStyle])),S=p((()=>[B.b(),r.popperClass,{[B.m("plain")]:!!r.content}])),E=p((()=>r.transition==="".concat(B.namespace.value,"-fade-in-linear"))),N=()=>{a("before-enter")},R=()=>{a("before-leave")},D=()=>{a("after-enter")},U=()=>{a("update:visible",!1),a("after-leave")};return t({popperRef:A,hide:()=>{var e;null==(e=x.value)||e.hide()}}),(e,t)=>(f(),b(n(o),y({ref_key:"tooltipRef",ref:x},e.$attrs,{trigger:e.trigger,placement:e.placement,disabled:e.disabled,visible:e.visible,transition:e.transition,"popper-options":e.popperOptions,tabindex:e.tabindex,content:e.content,offset:e.offset,"show-after":e.showAfter,"hide-after":e.hideAfter,"auto-close":e.autoClose,"show-arrow":e.showArrow,"aria-label":e.title,effect:e.effect,enterable:e.enterable,"popper-class":n(S),"popper-style":n(C),teleported:e.teleported,persistent:e.persistent,"gpu-acceleration":n(E),"onUpdate:visible":n(s),onBeforeShow:N,onBeforeHide:R,onShow:D,onHide:U}),{content:c((()=>[e.title?(f(),u("div",{key:0,class:v(n(B).e("title")),role:"title"},m(e.title),3)):h("v-if",!0),g(e.$slots,"default",{},(()=>[w(m(e.content),1)]))])),default:c((()=>[e.$slots.reference?g(e.$slots,"reference",{key:0}):h("v-if",!0)])),_:3},16,["trigger","placement","disabled","visible","transition","popper-options","tabindex","content","offset","show-after","hide-after","auto-close","show-arrow","aria-label","effect","enterable","popper-class","popper-style","teleported","persistent","gpu-acceleration","onUpdate:visible"]))}});const D=(e,t)=>{const o=t.arg||t.value,a=null==o?void 0:o.popperRef;a&&(a.triggerRef=e)};const U=A(B(R,[["__file","popover.vue"]]),{directive:x({mounted(e,t){D(e,t)},updated(e,t){D(e,t)}},"popover")});export{U as E};
