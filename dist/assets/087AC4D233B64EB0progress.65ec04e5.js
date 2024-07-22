/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{x as e,y as t,z as a,A as s,B as r,bt as n,bu as o,bv as c,bw as i,bx as l,aT as u,aw as p,o as d,c as f,C as h,D as y,b,n as v,G as g,t as k,h as x,g as w,w as m,p as B,q as D,H as I,I as N}from"./087AC4D233B64EB0index.39f1eb31.js";const F=e({type:{type:String,default:"line",values:["line","circle","dashboard"]},percentage:{type:Number,default:0,validator:e=>e>=0&&e<=100},status:{type:String,default:"",values:["","success","exception","warning"]},indeterminate:{type:Boolean,default:!1},duration:{type:Number,default:3},strokeWidth:{type:Number,default:6},strokeLinecap:{type:t(String),default:"round"},textInside:{type:Boolean,default:!1},width:{type:Number,default:126},showText:{type:Boolean,default:!0},color:{type:t([String,Array,Function]),default:""},striped:Boolean,stripedFlow:Boolean,format:{type:t(Function),default:e=>"".concat(e,"%")}}),T=["aria-valuenow"],S={viewBox:"0 0 100 100"},W=["d","stroke","stroke-linecap","stroke-width"],$=["d","stroke","opacity","stroke-linecap","stroke-width"],A={key:0},C=a({name:"ElProgress"});const E=N(I(a({...C,props:F,setup(e){const t=e,a={success:"#13ce66",exception:"#ff4949",warning:"#e6a23c",default:"#20a0ff"},I=s("progress"),N=r((()=>({width:"".concat(t.percentage,"%"),animationDuration:"".concat(t.duration,"s"),backgroundColor:V(t.percentage)}))),F=r((()=>(t.strokeWidth/t.width*100).toFixed(1))),C=r((()=>["circle","dashboard"].includes(t.type)?Number.parseInt("".concat(50-Number.parseFloat(F.value)/2),10):0)),E=r((()=>{const e=C.value,a="dashboard"===t.type;return"\n          M 50 50\n          m 0 ".concat(a?"":"-").concat(e,"\n          a ").concat(e," ").concat(e," 0 1 1 0 ").concat(a?"-":"").concat(2*e,"\n          a ").concat(e," ").concat(e," 0 1 1 0 ").concat(a?"":"-").concat(2*e,"\n          ")})),L=r((()=>2*Math.PI*C.value)),_=r((()=>"dashboard"===t.type?.75:1)),z=r((()=>{const e=-1*L.value*(1-_.value)/2;return"".concat(e,"px")})),M=r((()=>({strokeDasharray:"".concat(L.value*_.value,"px, ").concat(L.value,"px"),strokeDashoffset:z.value}))),P=r((()=>({strokeDasharray:"".concat(L.value*_.value*(t.percentage/100),"px, ").concat(L.value,"px"),strokeDashoffset:z.value,transition:"stroke-dasharray 0.6s ease 0s, stroke 0.6s ease, opacity ease 0.6s"}))),j=r((()=>{let e;return e=t.color?V(t.percentage):a[t.status]||a.default,e})),q=r((()=>"warning"===t.status?n:"line"===t.type?"success"===t.status?o:c:"success"===t.status?i:l)),G=r((()=>"line"===t.type?12+.4*t.strokeWidth:.111111*t.width+2)),H=r((()=>t.format(t.percentage)));const V=e=>{var a;const{color:s}=t;if(u(s))return s(e);if(p(s))return s;{const t=function(e){const t=100/e.length;return e.map(((e,a)=>p(e)?{color:e,percentage:(a+1)*t}:e)).sort(((e,t)=>e.percentage-t.percentage))}(s);for(const a of t)if(a.percentage>e)return a.color;return null==(a=t[t.length-1])?void 0:a.color}};return(e,t)=>(d(),f("div",{class:h([y(I).b(),y(I).m(e.type),y(I).is(e.status),{[y(I).m("without-text")]:!e.showText,[y(I).m("text-inside")]:e.textInside}]),role:"progressbar","aria-valuenow":e.percentage,"aria-valuemin":"0","aria-valuemax":"100"},["line"===e.type?(d(),f("div",{key:0,class:h(y(I).b("bar"))},[b("div",{class:h(y(I).be("bar","outer")),style:v({height:"".concat(e.strokeWidth,"px")})},[b("div",{class:h([y(I).be("bar","inner"),{[y(I).bem("bar","inner","indeterminate")]:e.indeterminate},{[y(I).bem("bar","inner","striped")]:e.striped},{[y(I).bem("bar","inner","striped-flow")]:e.stripedFlow}]),style:v(y(N))},[(e.showText||e.$slots.default)&&e.textInside?(d(),f("div",{key:0,class:h(y(I).be("bar","innerText"))},[g(e.$slots,"default",{percentage:e.percentage},(()=>[b("span",null,k(y(H)),1)]))],2)):x("v-if",!0)],6)],6)],2)):(d(),f("div",{key:1,class:h(y(I).b("circle")),style:v({height:"".concat(e.width,"px"),width:"".concat(e.width,"px")})},[(d(),f("svg",S,[b("path",{class:h(y(I).be("circle","track")),d:y(E),stroke:"var(".concat(y(I).cssVarName("fill-color-light"),", #e5e9f2)"),"stroke-linecap":e.strokeLinecap,"stroke-width":y(F),fill:"none",style:v(y(M))},null,14,W),b("path",{class:h(y(I).be("circle","path")),d:y(E),stroke:y(j),fill:"none",opacity:e.percentage?1:0,"stroke-linecap":e.strokeLinecap,"stroke-width":y(F),style:v(y(P))},null,14,$)]))],6)),!e.showText&&!e.$slots.default||e.textInside?x("v-if",!0):(d(),f("div",{key:2,class:h(y(I).e("text")),style:v({fontSize:"".concat(y(G),"px")})},[g(e.$slots,"default",{percentage:e.percentage},(()=>[e.status?(d(),w(y(D),{key:1},{default:m((()=>[(d(),w(B(y(q))))])),_:1})):(d(),f("span",A,k(y(H)),1))]))],6))],10,T))}}),[["__file","progress.vue"]]));export{E};