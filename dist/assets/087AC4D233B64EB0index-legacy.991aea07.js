/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
!function(){function e(e,r){var t=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);r&&(a=a.filter((function(r){return Object.getOwnPropertyDescriptor(e,r).enumerable}))),t.push.apply(t,a)}return t}function r(r){for(var a=1;a<arguments.length;a++){var i=null!=arguments[a]?arguments[a]:{};a%2?e(Object(i),!0).forEach((function(e){t(r,e,i[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(r,Object.getOwnPropertyDescriptors(i)):e(Object(i)).forEach((function(e){Object.defineProperty(r,e,Object.getOwnPropertyDescriptor(i,e))}))}return r}function t(e,r,t){var a;return(r="symbol"==typeof(a=function(e,r){if("object"!=typeof e||!e)return e;var t=e[Symbol.toPrimitive];if(void 0!==t){var a=t.call(e,r||"default");if("object"!=typeof a)return a;throw new TypeError("@@toPrimitive must return a primitive value.")}return("string"===r?String:Number)(e)}(r,"string"))?a:a+"")in e?Object.defineProperty(e,r,{value:t,enumerable:!0,configurable:!0,writable:!0}):e[r]=t,e}System.register(["./087AC4D233B64EB0index-legacy.e02f1eb2.js","./087AC4D233B64EB0image-viewer-legacy.5cb70cb3.js"],(function(e,t){"use strict";var a,i,n,c,l,s,o,u,v,p,f,d,y,g,b,m,h,w,x,S,z,j,O,k,I,P,E,B,A;return{setters:[function(e){a=e.x,i=e.b3,n=e.bs,c=e.U,l=e.y,s=e.z,o=e.A,u=e.r,v=e.B,p=e.aw,f=e.aa,d=e.L,y=e.o,g=e.c,b=e.n,m=e.D,h=e.g,w=e.w,x=e.p,S=e.q,z=e.G,j=e.C,O=e.H,k=e.I,I=e._,P=e.u,E=e.F,B=e.h},function(e){A=e.E}],execute:function(){var C=document.createElement("style");C.textContent=".el-avatar{--el-avatar-text-color: var(--el-color-white);--el-avatar-bg-color: var(--el-text-color-disabled);--el-avatar-text-size: 14px;--el-avatar-icon-size: 18px;--el-avatar-border-radius: var(--el-border-radius-base);--el-avatar-size-large: 56px;--el-avatar-size-small: 24px;--el-avatar-size: 40px;display:inline-flex;justify-content:center;align-items:center;box-sizing:border-box;text-align:center;overflow:hidden;color:var(--el-avatar-text-color);background:var(--el-avatar-bg-color);width:var(--el-avatar-size);height:var(--el-avatar-size);font-size:var(--el-avatar-text-size)}.el-avatar>img{display:block;width:100%;height:100%}.el-avatar--circle{border-radius:50%}.el-avatar--square{border-radius:var(--el-avatar-border-radius)}.el-avatar--icon{font-size:var(--el-avatar-icon-size)}.el-avatar--small{--el-avatar-size: 24px}.el-avatar--large{--el-avatar-size: 56px}.headerAvatar[data-v-33de2ed2]{display:flex;justify-content:center;align-items:center;margin-right:8px}.headerAvatar .avatar[data-v-33de2ed2]{width:100%;height:100%}.file[data-v-33de2ed2]{width:80px;height:80px;position:relative}\n",document.head.appendChild(C);var D=a({size:{type:[Number,String],values:i,default:"",validator:function(e){return n(e)}},shape:{type:String,values:["circle","square"],default:"circle"},icon:{type:c},src:{type:String,default:""},alt:String,srcSet:String,fit:{type:l(String),default:"cover"}}),_={error:function(e){return e instanceof Event}},q=["src","alt","srcset"],T=s({name:"ElAvatar"}),F=s(r(r({},T),{},{props:D,emits:_,setup:function(e,r){var t=r.emit,a=e,i=o("avatar"),c=u(!1),l=v((function(){var e=a.size,r=a.icon,t=a.shape,n=[i.b()];return p(e)&&n.push(i.m(e)),r&&n.push(i.m("icon")),t&&n.push(i.m(t)),n})),s=v((function(){var e=a.size;return n(e)?i.cssVarBlock({size:f(e)||""}):void 0})),O=v((function(){return{objectFit:a.fit}}));function k(e){c.value=!0,t("error",e)}return d((function(){return a.src}),(function(){return c.value=!1})),function(e,r){return y(),g("span",{class:j(m(l)),style:b(m(s))},[!e.src&&!e.srcSet||c.value?e.icon?(y(),h(m(S),{key:1},{default:w((function(){return[(y(),h(x(e.icon)))]})),_:1})):z(e.$slots,"default",{key:2}):(y(),g("img",{key:0,src:e.src,alt:e.alt,srcset:e.srcSet,style:b(m(O)),onError:k},null,44,q))],6)}}})),L=k(O(F,[["__file","avatar.vue"]])),N=""+new URL("087AC4D233B64EB0noBody.745c3d16.png",t.meta.url).href,U={class:"headerAvatar"},G=["src"],H=["src"],R=Object.assign({name:"CustomPic"},{__name:"index",props:{picType:{type:String,required:!1,default:"avatar"},picSrc:{type:String,required:!1,default:""},preview:{type:Boolean,default:!1}},setup:function(e){var r=e,t=u("https://www.sztv.com.cn/ysztv/web//"),a=u(N),i=P(),n=v((function(){return""===r.picSrc?""!==i.userInfo.headerImg&&"http"===i.userInfo.headerImg.slice(0,4)?i.userInfo.headerImg:t.value+i.userInfo.headerImg:""!==r.picSrc&&"http"===r.picSrc.slice(0,4)?r.picSrc:t.value+r.picSrc})),c=v((function(){return r.picSrc&&"http"!==r.picSrc.slice(0,4)?t.value+r.picSrc:r.picSrc})),l=v((function(){return r.preview?[c.value]:[]}));return function(r,t){var s=L,o=A;return y(),g("span",U,["avatar"===e.picType?(y(),g(E,{key:0},[m(i).userInfo.headerImg?(y(),h(s,{key:0,size:30,src:n.value},null,8,["src"])):(y(),h(s,{key:1,size:30,src:a.value},null,8,["src"]))],64)):B("",!0),"img"===e.picType?(y(),g(E,{key:1},[m(i).userInfo.headerImg?(y(),g("img",{key:0,src:n.value,class:"avatar"},null,8,G)):(y(),g("img",{key:1,src:a.value,class:"avatar"},null,8,H))],64)):B("",!0),"file"===e.picType?(y(),h(o,{key:2,src:c.value,class:"file","preview-src-list":l.value,"preview-teleported":!0},null,8,["src","preview-src-list"])):B("",!0)])}}});e("C",I(R,[["__scopeId","data-v-33de2ed2"]]))}}}))}();