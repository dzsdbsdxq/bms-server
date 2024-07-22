/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{K as e,a as l,r as a,o as u,c as o,d,w as t,i as m,F as s,m as n,g as i,h as V,b as p,E as v,j as r,k as c}from"./087AC4D233B64EB0index.39f1eb31.js";import{a as b,E as f}from"./087AC4D233B64EB0form-item.a86c41d2.js";import{E as _,a as y}from"./087AC4D233B64EB0collapse-item.58413777.js";import{E as U}from"./087AC4D233B64EB0switch.5bfb00d1.js";/* empty css                               *//* empty css                                *//* empty css                               */import{E as g}from"./087AC4D233B64EB0input-number.c9c9df06.js";/* empty css                              */import{E as q}from"./087AC4D233B64EB0checkbox.a3853383.js";import"./087AC4D233B64EB0tag.d3c1c3e4.js";import{E as h,a as w}from"./087AC4D233B64EB0select.7f5b5538.js";import"./087AC4D233B64EB0scrollbar.3894915f.js";import{a as k,s as B}from"./087AC4D233B64EB0system.285f9e44.js";import{E as z}from"./087AC4D233B64EB0index.d8bfc0de.js";import"./087AC4D233B64EB0castArray.209ac14b.js";import"./087AC4D233B64EB0_baseClone.a73122b0.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0_initCloneObject.0c0ca3ea.js";import"./087AC4D233B64EB0index.bfb39b2a.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";import"./087AC4D233B64EB0index.5d58b83c.js";import"./087AC4D233B64EB0strings.2df7e9e1.js";import"./087AC4D233B64EB0debounce.0aab18a9.js";import"./087AC4D233B64EB0isUndefined.d281c39a.js";const j={class:"system"},x=p("h2",null,"本地文件配置",-1),C=p("h2",null,"qiniu上传配置",-1),E=p("h2",null,"腾讯云COS上传配置",-1),A=p("h2",null,"阿里云OSS上传配置",-1),D=p("h2",null,"华为云Obs上传配置",-1),M={class:"mt-4"},S=Object.assign({name:"Config"},{__name:"system",setup(S){const K=l([]),I=a({system:{"iplimit-count":0,"iplimit-time":0},jwt:{},mysql:{},pgsql:{},excel:{},autocode:{},redis:{},mongo:{coll:"",options:"",database:"",username:"",password:"","min-pool-size":"","max-pool-size":"","socket-timeout-ms":"","connect-timeout-ms":"","is-zap":"",hosts:[{host:"",port:""}]},qiniu:{},"tencent-cos":{},"aliyun-oss":{},"hua-wei-obs":{},captcha:{},zap:{},local:{},email:{},timer:{detail:{}}}),O=async()=>{const e=await k();0===e.code&&(I.value=e.data.config)};O();const W=()=>{},T=async()=>{0===(await B({config:I.value})).code&&(v({type:"success",message:"配置文件设置成功"}),await O())},R=async()=>{var l;0===(await e({url:"/email/emailTest",method:"post",data:l})).code?(v({type:"success",message:"邮件发送成功"}),await O()):v({type:"error",message:"邮件发送失败"})};return(e,l)=>{const a=r,v=b,k=h,B=w,S=q,O=g,F=z,N=_,H=c,L=U,P=y,Z=f;return u(),o("div",j,[d(Z,{ref:"form",model:I.value,"label-width":"240px"},{default:t((()=>[d(P,{modelValue:K,"onUpdate:modelValue":l[104]||(l[104]=e=>K=e)},{default:t((()=>[d(N,{title:"系统配置",name:"1"},{default:t((()=>[d(v,{label:"端口值"},{default:t((()=>[d(a,{modelValue:I.value.system.addr,"onUpdate:modelValue":l[0]||(l[0]=e=>I.value.system.addr=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"数据库类型"},{default:t((()=>[d(B,{modelValue:I.value.system["db-type"],"onUpdate:modelValue":l[1]||(l[1]=e=>I.value.system["db-type"]=e),style:{width:"100%"}},{default:t((()=>[d(k,{value:"mysql"}),d(k,{value:"pgsql"})])),_:1},8,["modelValue"])])),_:1}),d(v,{label:"Oss类型"},{default:t((()=>[d(B,{modelValue:I.value.system["oss-type"],"onUpdate:modelValue":l[2]||(l[2]=e=>I.value.system["oss-type"]=e),style:{width:"100%"}},{default:t((()=>[d(k,{value:"local"}),d(k,{value:"qiniu"}),d(k,{value:"tencent-cos"}),d(k,{value:"aliyun-oss"}),d(k,{value:"huawei-obs"})])),_:1},8,["modelValue"])])),_:1}),d(v,{label:"多点登录拦截"},{default:t((()=>[d(S,{modelValue:I.value.system["use-multipoint"],"onUpdate:modelValue":l[3]||(l[3]=e=>I.value.system["use-multipoint"]=e)},{default:t((()=>[m("开启")])),_:1},8,["modelValue"])])),_:1}),d(v,{label:"开启redis"},{default:t((()=>[d(S,{modelValue:I.value.system["use-redis"],"onUpdate:modelValue":l[4]||(l[4]=e=>I.value.system["use-redis"]=e)},{default:t((()=>[m("开启")])),_:1},8,["modelValue"])])),_:1}),d(v,{label:"限流次数"},{default:t((()=>[d(O,{modelValue:I.value.system["iplimit-count"],"onUpdate:modelValue":l[5]||(l[5]=e=>I.value.system["iplimit-count"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"限流时间"},{default:t((()=>[d(O,{modelValue:I.value.system["iplimit-time"],"onUpdate:modelValue":l[6]||(l[6]=e=>I.value.system["iplimit-time"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(F,{content:"请修改完成后，注意一并修改前端env环境下的VITE_BASE_PATH",placement:"top-start"},{default:t((()=>[d(v,{label:"全局路由前缀"},{default:t((()=>[d(a,{modelValue:I.value.system["router-prefix"],"onUpdate:modelValue":l[7]||(l[7]=e=>I.value.system["router-prefix"]=e)},null,8,["modelValue"])])),_:1})])),_:1})])),_:1}),d(N,{title:"jwt签名",name:"2"},{default:t((()=>[d(v,{label:"jwt签名"},{default:t((()=>[d(a,{modelValue:I.value.jwt["signing-key"],"onUpdate:modelValue":l[8]||(l[8]=e=>I.value.jwt["signing-key"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"有效期"},{default:t((()=>[d(a,{modelValue:I.value.jwt["expires-time"],"onUpdate:modelValue":l[9]||(l[9]=e=>I.value.jwt["expires-time"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"缓冲期"},{default:t((()=>[d(a,{modelValue:I.value.jwt["buffer-time"],"onUpdate:modelValue":l[10]||(l[10]=e=>I.value.jwt["buffer-time"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"签发者"},{default:t((()=>[d(a,{modelValue:I.value.jwt.issuer,"onUpdate:modelValue":l[11]||(l[11]=e=>I.value.jwt.issuer=e)},null,8,["modelValue"])])),_:1})])),_:1}),d(N,{title:"Zap日志配置",name:"3"},{default:t((()=>[d(v,{label:"级别"},{default:t((()=>[d(a,{modelValue:I.value.zap.level,"onUpdate:modelValue":l[12]||(l[12]=e=>I.value.zap.level=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"输出"},{default:t((()=>[d(a,{modelValue:I.value.zap.format,"onUpdate:modelValue":l[13]||(l[13]=e=>I.value.zap.format=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"日志前缀"},{default:t((()=>[d(a,{modelValue:I.value.zap.prefix,"onUpdate:modelValue":l[14]||(l[14]=e=>I.value.zap.prefix=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"日志文件夹"},{default:t((()=>[d(a,{modelValue:I.value.zap.director,"onUpdate:modelValue":l[15]||(l[15]=e=>I.value.zap.director=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"编码级"},{default:t((()=>[d(a,{modelValue:I.value.zap["encode-level"],"onUpdate:modelValue":l[16]||(l[16]=e=>I.value.zap["encode-level"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"栈名"},{default:t((()=>[d(a,{modelValue:I.value.zap["stacktrace-key"],"onUpdate:modelValue":l[17]||(l[17]=e=>I.value.zap["stacktrace-key"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"日志留存时间(默认以天为单位)"},{default:t((()=>[d(a,{modelValue:I.value.zap["max-age"],"onUpdate:modelValue":l[18]||(l[18]=e=>I.value.zap["max-age"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"显示行"},{default:t((()=>[d(S,{modelValue:I.value.zap["show-line"],"onUpdate:modelValue":l[19]||(l[19]=e=>I.value.zap["show-line"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"输出控制台"},{default:t((()=>[d(S,{modelValue:I.value.zap["log-in-console"],"onUpdate:modelValue":l[20]||(l[20]=e=>I.value.zap["log-in-console"]=e)},null,8,["modelValue"])])),_:1})])),_:1}),d(N,{title:"Redis admin数据库配置",name:"4"},{default:t((()=>[d(v,{label:"库"},{default:t((()=>[d(a,{modelValue:I.value.redis.db,"onUpdate:modelValue":l[21]||(l[21]=e=>I.value.redis.db=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"地址"},{default:t((()=>[d(a,{modelValue:I.value.redis.addr,"onUpdate:modelValue":l[22]||(l[22]=e=>I.value.redis.addr=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"密码"},{default:t((()=>[d(a,{modelValue:I.value.redis.password,"onUpdate:modelValue":l[23]||(l[23]=e=>I.value.redis.password=e)},null,8,["modelValue"])])),_:1})])),_:1}),d(N,{title:"Mongo 数据库配置",name:"14"},{default:t((()=>[d(v,{label:"collection name(表名,一般不写)"},{default:t((()=>[d(a,{modelValue:I.value.mongo.coll,"onUpdate:modelValue":l[24]||(l[24]=e=>I.value.mongo.coll=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"mongodb options"},{default:t((()=>[d(a,{modelValue:I.value.mongo.options,"onUpdate:modelValue":l[25]||(l[25]=e=>I.value.mongo.options=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"database name(数据库名)"},{default:t((()=>[d(a,{modelValue:I.value.mongo.database,"onUpdate:modelValue":l[26]||(l[26]=e=>I.value.mongo.database=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"用户名"},{default:t((()=>[d(a,{modelValue:I.value.mongo.username,"onUpdate:modelValue":l[27]||(l[27]=e=>I.value.mongo.username=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"密码"},{default:t((()=>[d(a,{modelValue:I.value.mongo.password,"onUpdate:modelValue":l[28]||(l[28]=e=>I.value.mongo.password=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"最小连接池"},{default:t((()=>[d(a,{modelValue:I.value.mongo["min-pool-size"],"onUpdate:modelValue":l[29]||(l[29]=e=>I.value.mongo["min-pool-size"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"最大连接池"},{default:t((()=>[d(a,{modelValue:I.value.mongo["max-pool-size"],"onUpdate:modelValue":l[30]||(l[30]=e=>I.value.mongo["max-pool-size"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"socket超时时间"},{default:t((()=>[d(a,{modelValue:I.value.mongo["socket-timeout-ms"],"onUpdate:modelValue":l[31]||(l[31]=e=>I.value.mongo["socket-timeout-ms"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"连接超时时间"},{default:t((()=>[d(a,{modelValue:I.value.mongo["socket-timeout-ms"],"onUpdate:modelValue":l[32]||(l[32]=e=>I.value.mongo["socket-timeout-ms"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"是否开启zap日志"},{default:t((()=>[d(S,{modelValue:I.value.mongo["is-zap"],"onUpdate:modelValue":l[33]||(l[33]=e=>I.value.mongo["is-zap"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"hosts"},{default:t((()=>[(u(!0),o(s,null,n(I.value.mongo.hosts,((e,l)=>(u(),o(s,null,[(u(!0),o(s,null,n(e,((m,s)=>(u(),o("div",{key:s},[(u(),i(v,{key:l+s,label:s},{default:t((()=>[d(a,{modelValue:e[s],"onUpdate:modelValue":l=>e[s]=l},null,8,["modelValue","onUpdate:modelValue"])])),_:2},1032,["label"]))])))),128))],64)))),256))])),_:1})])),_:1}),d(N,{title:"邮箱配置",name:"5"},{default:t((()=>[d(v,{label:"接收者邮箱"},{default:t((()=>[d(a,{modelValue:I.value.email.to,"onUpdate:modelValue":l[34]||(l[34]=e=>I.value.email.to=e),placeholder:"可多个，以逗号分隔"},null,8,["modelValue"])])),_:1}),d(v,{label:"端口"},{default:t((()=>[d(a,{modelValue:I.value.email.port,"onUpdate:modelValue":l[35]||(l[35]=e=>I.value.email.port=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"发送者邮箱"},{default:t((()=>[d(a,{modelValue:I.value.email.from,"onUpdate:modelValue":l[36]||(l[36]=e=>I.value.email.from=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"host"},{default:t((()=>[d(a,{modelValue:I.value.email.host,"onUpdate:modelValue":l[37]||(l[37]=e=>I.value.email.host=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"是否为ssl"},{default:t((()=>[d(S,{modelValue:I.value.email["is-ssl"],"onUpdate:modelValue":l[38]||(l[38]=e=>I.value.email["is-ssl"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"secret"},{default:t((()=>[d(a,{modelValue:I.value.email.secret,"onUpdate:modelValue":l[39]||(l[39]=e=>I.value.email.secret=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"测试邮件"},{default:t((()=>[d(H,{size:"small",onClick:R},{default:t((()=>[m("测试邮件")])),_:1})])),_:1})])),_:1}),d(N,{title:"验证码配置",name:"7"},{default:t((()=>[d(v,{label:"字符长度"},{default:t((()=>[d(a,{modelValue:I.value.captcha["key-long"],"onUpdate:modelValue":l[40]||(l[40]=e=>I.value.captcha["key-long"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"图片宽度"},{default:t((()=>[d(a,{modelValue:I.value.captcha["img-width"],"onUpdate:modelValue":l[41]||(l[41]=e=>I.value.captcha["img-width"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"图片高度"},{default:t((()=>[d(a,{modelValue:I.value.captcha["img-height"],"onUpdate:modelValue":l[42]||(l[42]=e=>I.value.captcha["img-height"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1})])),_:1}),d(N,{title:"数据库配置",name:"9"},{default:t((()=>["mysql"===I.value.system["db-type"]?(u(),o(s,{key:0},[d(v,{label:"用户名"},{default:t((()=>[d(a,{modelValue:I.value.mysql.username,"onUpdate:modelValue":l[43]||(l[43]=e=>I.value.mysql.username=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"密码"},{default:t((()=>[d(a,{modelValue:I.value.mysql.password,"onUpdate:modelValue":l[44]||(l[44]=e=>I.value.mysql.password=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"地址"},{default:t((()=>[d(a,{modelValue:I.value.mysql.path,"onUpdate:modelValue":l[45]||(l[45]=e=>I.value.mysql.path=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"数据库"},{default:t((()=>[d(a,{modelValue:I.value.mysql["db-name"],"onUpdate:modelValue":l[46]||(l[46]=e=>I.value.mysql["db-name"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"前缀"},{default:t((()=>[d(a,{modelValue:I.value.mysql.refix,"onUpdate:modelValue":l[47]||(l[47]=e=>I.value.mysql.refix=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"复数表"},{default:t((()=>[d(L,{modelValue:I.value.mysql.singular,"onUpdate:modelValue":l[48]||(l[48]=e=>I.value.mysql.singular=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"引擎"},{default:t((()=>[d(a,{modelValue:I.value.mysql.engine,"onUpdate:modelValue":l[49]||(l[49]=e=>I.value.mysql.engine=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"maxIdleConns"},{default:t((()=>[d(a,{modelValue:I.value.mysql["max-idle-conns"],"onUpdate:modelValue":l[50]||(l[50]=e=>I.value.mysql["max-idle-conns"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"maxOpenConns"},{default:t((()=>[d(a,{modelValue:I.value.mysql["max-open-conns"],"onUpdate:modelValue":l[51]||(l[51]=e=>I.value.mysql["max-open-conns"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"写入日志"},{default:t((()=>[d(S,{modelValue:I.value.mysql["log-zap"],"onUpdate:modelValue":l[52]||(l[52]=e=>I.value.mysql["log-zap"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"日志模式"},{default:t((()=>[d(a,{modelValue:I.value.mysql["log-mode"],"onUpdate:modelValue":l[53]||(l[53]=e=>I.value.mysql["log-mode"]=e)},null,8,["modelValue"])])),_:1})],64)):V("",!0),"pgsql"===I.value.system["db-type"]?(u(),o(s,{key:1},[d(v,{label:"用户名"},{default:t((()=>[d(a,{modelValue:I.value.pgsql.username,"onUpdate:modelValue":l[54]||(l[54]=e=>I.value.pgsql.username=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"密码"},{default:t((()=>[d(a,{modelValue:I.value.pgsql.password,"onUpdate:modelValue":l[55]||(l[55]=e=>I.value.pgsql.password=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"地址"},{default:t((()=>[d(a,{modelValue:I.value.pgsql.path,"onUpdate:modelValue":l[56]||(l[56]=e=>I.value.pgsql.path=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"数据库"},{default:t((()=>[d(a,{modelValue:I.value.pgsql.dbname,"onUpdate:modelValue":l[57]||(l[57]=e=>I.value.pgsql.dbname=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"前缀"},{default:t((()=>[d(a,{modelValue:I.value.pgsql.refix,"onUpdate:modelValue":l[58]||(l[58]=e=>I.value.pgsql.refix=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"复数表"},{default:t((()=>[d(L,{modelValue:I.value.pgsql.singular,"onUpdate:modelValue":l[59]||(l[59]=e=>I.value.pgsql.singular=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"引擎"},{default:t((()=>[d(a,{modelValue:I.value.pgsql.engine,"onUpdate:modelValue":l[60]||(l[60]=e=>I.value.pgsql.engine=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"maxIdleConns"},{default:t((()=>[d(a,{modelValue:I.value.pgsql["max-idle-conns"],"onUpdate:modelValue":l[61]||(l[61]=e=>I.value.pgsql["max-idle-conns"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"maxOpenConns"},{default:t((()=>[d(a,{modelValue:I.value.pgsql["max-open-conns"],"onUpdate:modelValue":l[62]||(l[62]=e=>I.value.pgsql["max-open-conns"]=e),modelModifiers:{number:!0}},null,8,["modelValue"])])),_:1}),d(v,{label:"写入日志"},{default:t((()=>[d(S,{modelValue:I.value.pgsql["log-zap"],"onUpdate:modelValue":l[63]||(l[63]=e=>I.value.pgsql["log-zap"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"日志模式"},{default:t((()=>[d(a,{modelValue:I.value.pgsql["log-mode"],"onUpdate:modelValue":l[64]||(l[64]=e=>I.value.pgsql["log-mode"]=e)},null,8,["modelValue"])])),_:1})],64)):V("",!0)])),_:1}),d(N,{title:"oss配置",name:"10"},{default:t((()=>["local"===I.value.system["oss-type"]?(u(),o(s,{key:0},[x,d(v,{label:"本地文件访问路径"},{default:t((()=>[d(a,{modelValue:I.value.local.path,"onUpdate:modelValue":l[65]||(l[65]=e=>I.value.local.path=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"本地文件存储路径"},{default:t((()=>[d(a,{modelValue:I.value.local["store-path"],"onUpdate:modelValue":l[66]||(l[66]=e=>I.value.local["store-path"]=e)},null,8,["modelValue"])])),_:1})],64)):V("",!0),"qiniu"===I.value.system["oss-type"]?(u(),o(s,{key:1},[C,d(v,{label:"存储区域"},{default:t((()=>[d(a,{modelValue:I.value.qiniu.zone,"onUpdate:modelValue":l[67]||(l[67]=e=>I.value.qiniu.zone=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"空间名称"},{default:t((()=>[d(a,{modelValue:I.value.qiniu.bucket,"onUpdate:modelValue":l[68]||(l[68]=e=>I.value.qiniu.bucket=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"CDN加速域名"},{default:t((()=>[d(a,{modelValue:I.value.qiniu["img-path"],"onUpdate:modelValue":l[69]||(l[69]=e=>I.value.qiniu["img-path"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"是否使用https"},{default:t((()=>[d(S,{modelValue:I.value.qiniu["use-https"],"onUpdate:modelValue":l[70]||(l[70]=e=>I.value.qiniu["use-https"]=e)},{default:t((()=>[m("开启")])),_:1},8,["modelValue"])])),_:1}),d(v,{label:"accessKey"},{default:t((()=>[d(a,{modelValue:I.value.qiniu["access-key"],"onUpdate:modelValue":l[71]||(l[71]=e=>I.value.qiniu["access-key"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"secretKey"},{default:t((()=>[d(a,{modelValue:I.value.qiniu["secret-key"],"onUpdate:modelValue":l[72]||(l[72]=e=>I.value.qiniu["secret-key"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"上传是否使用CDN上传加速"},{default:t((()=>[d(S,{modelValue:I.value.qiniu["use-cdn-domains"],"onUpdate:modelValue":l[73]||(l[73]=e=>I.value.qiniu["use-cdn-domains"]=e)},{default:t((()=>[m("开启")])),_:1},8,["modelValue"])])),_:1})],64)):V("",!0),"tencent-cos"===I.value.system["oss-type"]?(u(),o(s,{key:2},[E,d(v,{label:"存储桶名称"},{default:t((()=>[d(a,{modelValue:I.value["tencent-cos"].bucket,"onUpdate:modelValue":l[74]||(l[74]=e=>I.value["tencent-cos"].bucket=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"所属地域"},{default:t((()=>[d(a,{modelValue:I.value["tencent-cos"].region,"onUpdate:modelValue":l[75]||(l[75]=e=>I.value["tencent-cos"].region=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"secretID"},{default:t((()=>[d(a,{modelValue:I.value["tencent-cos"]["secret-id"],"onUpdate:modelValue":l[76]||(l[76]=e=>I.value["tencent-cos"]["secret-id"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"secretKey"},{default:t((()=>[d(a,{modelValue:I.value["tencent-cos"]["secret-key"],"onUpdate:modelValue":l[77]||(l[77]=e=>I.value["tencent-cos"]["secret-key"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"路径前缀"},{default:t((()=>[d(a,{modelValue:I.value["tencent-cos"]["path-prefix"],"onUpdate:modelValue":l[78]||(l[78]=e=>I.value["tencent-cos"]["path-prefix"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"访问域名"},{default:t((()=>[d(a,{modelValue:I.value["tencent-cos"]["base-url"],"onUpdate:modelValue":l[79]||(l[79]=e=>I.value["tencent-cos"]["base-url"]=e)},null,8,["modelValue"])])),_:1})],64)):V("",!0),"aliyun-oss"===I.value.system["oss-type"]?(u(),o(s,{key:3},[A,d(v,{label:"区域"},{default:t((()=>[d(a,{modelValue:I.value["aliyun-oss"].endpoint,"onUpdate:modelValue":l[80]||(l[80]=e=>I.value["aliyun-oss"].endpoint=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"accessKeyId"},{default:t((()=>[d(a,{modelValue:I.value["aliyun-oss"]["access-key-id"],"onUpdate:modelValue":l[81]||(l[81]=e=>I.value["aliyun-oss"]["access-key-id"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"accessKeySecret"},{default:t((()=>[d(a,{modelValue:I.value["aliyun-oss"]["access-key-secret"],"onUpdate:modelValue":l[82]||(l[82]=e=>I.value["aliyun-oss"]["access-key-secret"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"存储桶名称"},{default:t((()=>[d(a,{modelValue:I.value["aliyun-oss"]["bucket-name"],"onUpdate:modelValue":l[83]||(l[83]=e=>I.value["aliyun-oss"]["bucket-name"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"访问域名"},{default:t((()=>[d(a,{modelValue:I.value["aliyun-oss"]["bucket-url"],"onUpdate:modelValue":l[84]||(l[84]=e=>I.value["aliyun-oss"]["bucket-url"]=e)},null,8,["modelValue"])])),_:1})],64)):V("",!0),"huawei-obs"===I.value.system["oss-type"]?(u(),o(s,{key:4},[D,d(v,{label:"路径"},{default:t((()=>[d(a,{modelValue:I.value["hua-wei-obs"].path,"onUpdate:modelValue":l[85]||(l[85]=e=>I.value["hua-wei-obs"].path=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"存储桶名称"},{default:t((()=>[d(a,{modelValue:I.value["hua-wei-obs"].bucket,"onUpdate:modelValue":l[86]||(l[86]=e=>I.value["hua-wei-obs"].bucket=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"区域"},{default:t((()=>[d(a,{modelValue:I.value["hua-wei-obs"].endpoint,"onUpdate:modelValue":l[87]||(l[87]=e=>I.value["hua-wei-obs"].endpoint=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"accessKey"},{default:t((()=>[d(a,{modelValue:I.value["hua-wei-obs"]["access-key"],"onUpdate:modelValue":l[88]||(l[88]=e=>I.value["hua-wei-obs"]["access-key"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"secretKey"},{default:t((()=>[d(a,{modelValue:I.value["hua-wei-obs"]["secret-key"],"onUpdate:modelValue":l[89]||(l[89]=e=>I.value["hua-wei-obs"]["secret-key"]=e)},null,8,["modelValue"])])),_:1})],64)):V("",!0)])),_:1}),d(N,{title:"Excel上传配置",name:"11"},{default:t((()=>[d(v,{label:"合成目标地址"},{default:t((()=>[d(a,{modelValue:I.value.excel.dir,"onUpdate:modelValue":l[90]||(l[90]=e=>I.value.excel.dir=e)},null,8,["modelValue"])])),_:1})])),_:1}),d(N,{title:"自动化代码配置",name:"12"},{default:t((()=>[d(v,{label:"是否自动重启(linux)"},{default:t((()=>[d(S,{modelValue:I.value.autocode["transfer-restart"],"onUpdate:modelValue":l[91]||(l[91]=e=>I.value.autocode["transfer-restart"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"root(项目根路径)"},{default:t((()=>[d(a,{modelValue:I.value.autocode.root,"onUpdate:modelValue":l[92]||(l[92]=e=>I.value.autocode.root=e),disabled:""},null,8,["modelValue"])])),_:1}),d(v,{label:"Server(后端代码地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["transfer-restart"],"onUpdate:modelValue":l[93]||(l[93]=e=>I.value.autocode["transfer-restart"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"SApi(后端api文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["server-api"],"onUpdate:modelValue":l[94]||(l[94]=e=>I.value.autocode["server-api"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"SInitialize(后端Initialize文件夹)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["server-initialize"],"onUpdate:modelValue":l[95]||(l[95]=e=>I.value.autocode["server-initialize"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"SModel(后端Model文件地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["server-model"],"onUpdate:modelValue":l[96]||(l[96]=e=>I.value.autocode["server-model"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"SRequest(后端Request文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["server-request"],"onUpdate:modelValue":l[97]||(l[97]=e=>I.value.autocode["server-request"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"SRouter(后端Router文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["server-router"],"onUpdate:modelValue":l[98]||(l[98]=e=>I.value.autocode["server-router"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"SService(后端Service文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["server-service"],"onUpdate:modelValue":l[99]||(l[99]=e=>I.value.autocode["server-service"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"Web(前端文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode.web,"onUpdate:modelValue":l[100]||(l[100]=e=>I.value.autocode.web=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"WApi(后端WApi文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["web-api"],"onUpdate:modelValue":l[101]||(l[101]=e=>I.value.autocode["web-api"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"WForm(后端WForm文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["web-form"],"onUpdate:modelValue":l[102]||(l[102]=e=>I.value.autocode["web-form"]=e)},null,8,["modelValue"])])),_:1}),d(v,{label:"WTable(后端WTable文件夹地址)"},{default:t((()=>[d(a,{modelValue:I.value.autocode["web-table"],"onUpdate:modelValue":l[103]||(l[103]=e=>I.value.autocode["web-table"]=e)},null,8,["modelValue"])])),_:1})])),_:1})])),_:1},8,["modelValue"])])),_:1},8,["model"]),p("div",M,[d(H,{type:"primary",size:"small",onClick:T},{default:t((()=>[m("立即更新")])),_:1}),d(H,{type:"primary",size:"small",onClick:W},{default:t((()=>[m("重启服务（开发中）")])),_:1})])])}}});export{S as default};
