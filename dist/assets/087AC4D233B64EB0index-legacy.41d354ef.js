/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
System.register(["./087AC4D233B64EB0index-legacy.e02f1eb2.js","./087AC4D233B64EB0upload-legacy.c1cf5fb4.js","./087AC4D233B64EB0progress-legacy.09cb36e5.js","./087AC4D233B64EB0cloneDeep-legacy.1e61118e.js","./087AC4D233B64EB0_baseClone-legacy.1583cc72.js","./087AC4D233B64EB0_Uint8Array-legacy.792a8164.js","./087AC4D233B64EB0_initCloneObject-legacy.f845501a.js","./087AC4D233B64EB0isEqual-legacy.e931ef95.js"],(function(n,e){"use strict";var l,t,a,u,c,o,s,i,r,d,f,B;return{setters:[function(n){l=n.u,t=n.r,a=n.M,u=n.o,c=n.c,o=n.d,s=n.w,i=n.b,r=n.i,d=n.E,f=n.q},function(n){B=n.E},null,null,null,null,null,null],execute:function(){var e=i("div",{class:"el-upload__text"},[r(" 拖拽或"),i("em",null,"点击上传")],-1),g=i("div",{class:"el-upload__tip"}," 请把安装包的zip拖拽至此处上传 ",-1);n("default",{__name:"index",setup:function(n){l();var i=t("https://www.sztv.com.cn/ysztv/web/"),r=function(n){if(0===n.code){var e="";n.data&&n.data.forEach((function(n,l){e+="".concat(l+1,".").concat(n.msg,"\n")})),alert(e)}else d.error(n.msg)};return function(n,l){var t=a("upload-filled"),d=f,p=B;return u(),c("div",null,[o(p,{drag:"",action:"".concat(i.value,"/autoCode/installPlugin"),"show-file-list":!1,"on-success":r,"on-error":r,name:"plug"},{tip:s((function(){return[g]})),default:s((function(){return[o(d,{class:"el-icon--upload"},{default:s((function(){return[o(t)]})),_:1}),e]})),_:1},8,["action"])])}}})}}}));