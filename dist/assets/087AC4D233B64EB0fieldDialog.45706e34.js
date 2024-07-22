/*! 
 Build based on gin-vue-admin 
 Time : 1717039643000 */
import{r as e,o as l,c as a,d as o,w as u,F as d,m as t,g as i,b as r,j as s,k as m}from"./087AC4D233B64EB0index.39f1eb31.js";import{a as p,E as n}from"./087AC4D233B64EB0form-item.a86c41d2.js";import{E as f}from"./087AC4D233B64EB0switch.5bfb00d1.js";import{E as c}from"./087AC4D233B64EB0checkbox.a3853383.js";/* empty css                              */import"./087AC4D233B64EB0tag.d3c1c3e4.js";import{E as v,a as b}from"./087AC4D233B64EB0select.7f5b5538.js";import"./087AC4D233B64EB0scrollbar.3894915f.js";/* empty css                               *//* empty css                               */import{a as y,t as B}from"./087AC4D233B64EB0stringFun.24b05cf5.js";import{g as V}from"./087AC4D233B64EB0sysDictionary.e0472ec6.js";import{_ as g}from"./087AC4D233B64EB0warningBar.beb718a3.js";import"./087AC4D233B64EB0castArray.209ac14b.js";import"./087AC4D233B64EB0_baseClone.a73122b0.js";import"./087AC4D233B64EB0_Uint8Array.02b83471.js";import"./087AC4D233B64EB0_initCloneObject.0c0ca3ea.js";import"./087AC4D233B64EB0isEqual.e024b3a6.js";import"./087AC4D233B64EB0hasIn.9f12b2a1.js";import"./087AC4D233B64EB0index.d8bfc0de.js";import"./087AC4D233B64EB0isUndefined.d281c39a.js";import"./087AC4D233B64EB0index.5d58b83c.js";import"./087AC4D233B64EB0strings.2df7e9e1.js";import"./087AC4D233B64EB0debounce.0aab18a9.js";const E=r("span",{style:{"font-size":"12px"}},"自动填充",-1),D=Object.assign({name:"FieldDialog"},{__name:"fieldDialog",props:{dialogMiddle:{type:Object,default:function(){return{}}},typeOptions:{type:Array,default:function(){return[]}},typeSearchOptions:{type:Array,default:function(){return[]}}},setup(r,{expose:D}){const j=r,_=e({}),A=e([]),C=e({fieldName:[{required:!0,message:"请输入字段英文名",trigger:"blur"}],fieldDesc:[{required:!0,message:"请输入字段中文名",trigger:"blur"}],fieldJson:[{required:!0,message:"请输入字段格式化json",trigger:"blur"}],columnName:[{required:!0,message:"请输入数据库字段",trigger:"blur"}],fieldType:[{required:!0,message:"请选择字段类型",trigger:"blur"}]});(async()=>{_.value=j.dialogMiddle;const e=await V({page:1,pageSize:999999});A.value=e.data})();const T=()=>{_.value.fieldJson=y(_.value.fieldName),_.value.columnName=B(_.value.fieldJson)},h=e=>{const l=_.value.fieldType;return"string"!==l&&"LIKE"===e||"int"!==l&&"time.Time"!==l&&"float64"!==l&&("BETWEEN"===e||"NOT BETWEEN"===e)},U=()=>{_.value.fieldSearchType="",_.value.dictType=""},N=e(null);return D({fieldDialogFrom:N}),(e,y)=>{const B=s,V=m,D=p,j=v,w=b,x=c,O=f,k=n;return l(),a("div",null,[o(g,{title:"id , created_at , updated_at , deleted_at 会自动生成请勿重复创建。搜索时如果条件为LIKE只支持字符串"}),o(k,{ref_key:"fieldDialogFrom",ref:N,model:_.value,"label-width":"120px","label-position":"right",rules:C.value,class:"grid grid-cols-2"},{default:u((()=>[o(D,{label:"字段名称",prop:"fieldName"},{default:u((()=>[o(B,{modelValue:_.value.fieldName,"onUpdate:modelValue":y[0]||(y[0]=e=>_.value.fieldName=e),autocomplete:"off",style:{width:"80%"}},null,8,["modelValue"]),o(V,{style:{width:"18%","margin-left":"2%"},size:"small",onClick:T},{default:u((()=>[E])),_:1})])),_:1}),o(D,{label:"字段中文名",prop:"fieldDesc"},{default:u((()=>[o(B,{modelValue:_.value.fieldDesc,"onUpdate:modelValue":y[1]||(y[1]=e=>_.value.fieldDesc=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),o(D,{label:"字段JSON",prop:"fieldJson"},{default:u((()=>[o(B,{modelValue:_.value.fieldJson,"onUpdate:modelValue":y[2]||(y[2]=e=>_.value.fieldJson=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),o(D,{label:"数据库字段名",prop:"columnName"},{default:u((()=>[o(B,{modelValue:_.value.columnName,"onUpdate:modelValue":y[3]||(y[3]=e=>_.value.columnName=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),o(D,{label:"数据库字段描述",prop:"comment"},{default:u((()=>[o(B,{modelValue:_.value.comment,"onUpdate:modelValue":y[4]||(y[4]=e=>_.value.comment=e),autocomplete:"off"},null,8,["modelValue"])])),_:1}),o(D,{label:"字段类型",prop:"fieldType"},{default:u((()=>[o(w,{modelValue:_.value.fieldType,"onUpdate:modelValue":y[5]||(y[5]=e=>_.value.fieldType=e),style:{width:"100%"},placeholder:"请选择字段类型",clearable:"",onChange:U},{default:u((()=>[(l(!0),a(d,null,t(r.typeOptions,(e=>(l(),i(j,{key:e.value,label:e.label,value:e.value,disabled:e.disabled},null,8,["label","value","disabled"])))),128))])),_:1},8,["modelValue"])])),_:1}),o(D,{label:"enum"===_.value.fieldType?"枚举值":"类型长度",prop:"dataTypeLong"},{default:u((()=>[o(B,{modelValue:_.value.dataTypeLong,"onUpdate:modelValue":y[6]||(y[6]=e=>_.value.dataTypeLong=e),placeholder:"enum"===_.value.fieldType?"例:'北京','天津'":"数据库类型长度"},null,8,["modelValue","placeholder"])])),_:1},8,["label"]),o(D,{label:"字段查询条件",prop:"fieldSearchType"},{default:u((()=>[o(w,{modelValue:_.value.fieldSearchType,"onUpdate:modelValue":y[7]||(y[7]=e=>_.value.fieldSearchType=e),disabled:"json"===_.value.fieldType,style:{width:"100%"},placeholder:"请选择字段查询条件",clearable:""},{default:u((()=>[(l(!0),a(d,null,t(r.typeSearchOptions,(e=>(l(),i(j,{key:e.value,label:e.label,value:e.value,disabled:h(e.value)},null,8,["label","value","disabled"])))),128))])),_:1},8,["modelValue","disabled"])])),_:1}),o(D,{label:"关联字典",prop:"dictType"},{default:u((()=>[o(w,{modelValue:_.value.dictType,"onUpdate:modelValue":y[8]||(y[8]=e=>_.value.dictType=e),style:{width:"100%"},disabled:"int"!==_.value.fieldType&&"string"!==_.value.fieldType,placeholder:"请选择字典",clearable:""},{default:u((()=>[(l(!0),a(d,null,t(A.value,(e=>(l(),i(j,{key:e.type,label:"".concat(e.type,"(").concat(e.name,")"),value:e.type},null,8,["label","value"])))),128))])),_:1},8,["modelValue","disabled"])])),_:1}),o(D,{label:"主键"},{default:u((()=>[o(x,{modelValue:_.value.primaryKey,"onUpdate:modelValue":y[9]||(y[9]=e=>_.value.primaryKey=e)},null,8,["modelValue"])])),_:1}),o(D,{label:"是否排序"},{default:u((()=>[o(O,{modelValue:_.value.sort,"onUpdate:modelValue":y[10]||(y[10]=e=>_.value.sort=e)},null,8,["modelValue"])])),_:1}),o(D,{label:"是否必填"},{default:u((()=>[o(O,{modelValue:_.value.require,"onUpdate:modelValue":y[11]||(y[11]=e=>_.value.require=e)},null,8,["modelValue"])])),_:1}),o(D,{label:"是否可清空"},{default:u((()=>[o(O,{modelValue:_.value.clearable,"onUpdate:modelValue":y[12]||(y[12]=e=>_.value.clearable=e)},null,8,["modelValue"])])),_:1}),o(D,{label:"校验失败文案"},{default:u((()=>[o(B,{modelValue:_.value.errorText,"onUpdate:modelValue":y[13]||(y[13]=e=>_.value.errorText=e)},null,8,["modelValue"])])),_:1})])),_:1},8,["model","rules"])])}}});export{D as default};