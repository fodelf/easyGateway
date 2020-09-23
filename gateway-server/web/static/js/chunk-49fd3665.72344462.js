(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-49fd3665"],{"1d00":function(e,t,r){"use strict";r.r(t);var s=function(){var e=this,t=e.$createElement,r=e._self._c||t;return r("el-form",{ref:"ruleForm",staticClass:"projectAdd",attrs:{model:e.ruleForm,rules:e.serviceRules,inline:!0,"label-width":"150px","label-position":"left",disabled:"check"==e.$route.query.type}},[r("div",{staticStyle:{width:"80px"},on:{click:e.cancel}},[r("i",{staticClass:"el-icon-d-arrow-left",staticStyle:{color:"white","font-size":"14px",cursor:"pointer","margin-bottom":"20px"}},[e._v("返回")])]),r("el-tabs",{model:{value:e.baseInfo,callback:function(t){e.baseInfo=t},expression:"baseInfo"}},[r("el-tab-pane",{attrs:{label:"基本信息",name:"baseInfo"}})],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"服务名称",prop:"serviceName"}},[r("el-input",{attrs:{type:"text",placeholder:"请输入英文或者数字"},model:{value:e.ruleForm.serviceName,callback:function(t){e.$set(e.ruleForm,"serviceName",t)},expression:"ruleForm.serviceName"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"服务类型",prop:"serviceType"}},[r("el-select",{attrs:{placeholder:"请选择服务类型"},model:{value:e.ruleForm.serviceType,callback:function(t){e.$set(e.ruleForm,"serviceType",t)},expression:"ruleForm.serviceType"}},e._l(e.serverList,(function(e){return r("el-option",{key:e.value,attrs:{label:e.label,value:e.value}})})),1)],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"服务地址",prop:"serviceAddress"}},[r("el-input",{attrs:{type:"text",placeholder:"示例：127.0.0.1"},model:{value:e.ruleForm.serviceAddress,callback:function(t){e.$set(e.ruleForm,"serviceAddress",t)},expression:"ruleForm.serviceAddress"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"服务端口",prop:"servicePort"}},[r("el-input",{attrs:{type:"text",placeholder:"示例：3000"},model:{value:e.ruleForm.servicePort,callback:function(t){e.$set(e.ruleForm,"servicePort",e._n(t))},expression:"ruleForm.servicePort"}})],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"熔断",prop:"serviceBreak"}},[r("el-input",{attrs:{type:"text",placeholder:"请求超时时间"},model:{value:e.ruleForm.serviceBreak,callback:function(t){e.$set(e.ruleForm,"serviceBreak",t)},expression:"ruleForm.serviceBreak"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"限流",prop:"serviceLimit"}},[r("el-input",{attrs:{type:"text",placeholder:"请求次数限制"},model:{value:e.ruleForm.serviceLimit,callback:function(t){e.$set(e.ruleForm,"serviceLimit",t)},expression:"ruleForm.serviceLimit"}})],1)],1)],1),r("el-row",[r("el-form-item",{staticClass:"agency-item",attrs:{label:"代理规则（必须至少一个代理规则）","label-width":"250px"}},[r("el-button",{attrs:{type:"primary",icon:"el-icon-plus",circle:""},on:{click:e.addRule}})],1)],1),e._l(e.ruleForm.serviceRules,(function(t,s){return r("el-row",{key:s,attrs:{gutter:20}},[r("el-row",{staticStyle:{"padding-left":"10px"},attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"拦截地址",prop:"serviceRules."+s+".url",rules:e.rules.apiUrl}},[r("el-input",{attrs:{type:"text",placeholder:"示例：/api,不能以ui开头"},model:{value:t.url,callback:function(r){e.$set(t,"url",r)},expression:"item.url"}})],1)],1),e.ruleForm.serviceRules.length>1?r("el-col",{attrs:{span:6}},[r("el-button",{attrs:{type:"primary",icon:"el-icon-minus",circle:""},on:{click:function(t){return e.deleteRule(s)}}})],1):e._e()],1),r("el-row",{staticStyle:{"padding-left":"10px"},attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"重写前缀",prop:"serviceRules."+s+".pathReWriteBefore",rules:e.rules.pathReWriteBefore}},[r("el-input",{attrs:{type:"text",placeholder:"示例：/api"},model:{value:t.pathReWriteBefore,callback:function(r){e.$set(t,"pathReWriteBefore",r)},expression:"item.pathReWriteBefore"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"重写地址",prop:"serviceRules."+s+".pathReWriteUrl",rules:e.rules.pathReWriteUrl}},[r("el-input",{attrs:{type:"text",placeholder:"示例：/api"},model:{value:t.pathReWriteUrl,callback:function(r){e.$set(t,"pathReWriteUrl",r)},expression:"item.pathReWriteUrl"}})],1)],1)],1)],1)})),r("el-tabs",{model:{value:e.useConsul,callback:function(t){e.useConsul=t},expression:"useConsul"}},[r("el-tab-pane",{attrs:{label:"注册中心服务信息配合consul使用",name:"useConsul"}})],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"服务ID"}},[r("el-input",{attrs:{type:"text",placeholder:"请输入英文或者数字"},model:{value:e.ruleForm.useConsulId,callback:function(t){e.$set(e.ruleForm,"useConsulId",t)},expression:"ruleForm.useConsulId"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"服务标签"}},[r("el-input",{attrs:{type:"text",placeholder:"请输入英文或者数字"},model:{value:e.ruleForm.useConsulTag,callback:function(t){e.$set(e.ruleForm,"useConsulTag",t)},expression:"ruleForm.useConsulTag"}})],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"监控检查接口",prop:"useConsulCheckPath"}},[r("el-input",{attrs:{type:"text",placeholder:"示例：/checkHealth"},model:{value:e.ruleForm.useConsulCheckPath,callback:function(t){e.$set(e.ruleForm,"useConsulCheckPath",t)},expression:"ruleForm.useConsulCheckPath"}})],1)],1)],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"健康检查间隔(秒)",prop:"useConsulInterval"}},[r("el-input",{attrs:{type:"text",placeholder:"默认10s"},model:{value:e.ruleForm.useConsulInterval,callback:function(t){e.$set(e.ruleForm,"useConsulInterval",t)},expression:"ruleForm.useConsulInterval"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"健康检查超时时间(秒)",prop:"useConsulTimeout"}},[r("el-input",{attrs:{type:"text",placeholder:"默认3s"},model:{value:e.ruleForm.useConsulTimeout,callback:function(t){e.$set(e.ruleForm,"useConsulTimeout",t)},expression:"ruleForm.useConsulTimeout"}})],1)],1)],1),r("el-tabs",{model:{value:e.messageWarn,callback:function(t){e.messageWarn=t},expression:"messageWarn"}},[r("el-tab-pane",{attrs:{label:"钉钉信息（消息告警）",name:"messageWarn"}})],1),r("el-row",{attrs:{gutter:20}},[r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"钉钉accessToken"}},[r("el-input",{attrs:{type:"text",placeholder:"请输入"},model:{value:e.ruleForm.dingdingAccessToken,callback:function(t){e.$set(e.ruleForm,"dingdingAccessToken",t)},expression:"ruleForm.dingdingAccessToken"}})],1)],1),r("el-col",{attrs:{span:8}},[r("el-form-item",{attrs:{label:"钉钉secret"}},[r("el-input",{attrs:{type:"text",placeholder:"请输入"},model:{value:e.ruleForm.dingdingSercet,callback:function(t){e.$set(e.ruleForm,"dingdingSercet",t)},expression:"ruleForm.dingdingSercet"}})],1)],1)],1),r("el-row",[r("el-form-item",{attrs:{label:"钉钉联系人手机号"}},[e._l(e.ruleForm.dingdingList,(function(t){return r("el-tag",{key:t,attrs:{type:"warning",closable:"","disable-transitions":!1},on:{close:function(r){return e.handleClose(t)}}},[e._v(" "+e._s(t)+" ")])})),e.inputVisible?r("el-input",{ref:"saveTagInput",staticClass:"input-new-tag",attrs:{size:"small"},on:{blur:e.handleInputConfirm},nativeOn:{keyup:function(t){return!t.type.indexOf("key")&&e._k(t.keyCode,"enter",13,t.key,"Enter")?null:e.handleInputConfirm(t)}},model:{value:e.inputValue,callback:function(t){e.inputValue=t},expression:"inputValue"}}):r("el-button",{staticClass:"button-new-tag",attrs:{type:"primary",size:"small",icon:"el-icon-plus",circle:""},on:{click:e.showInput}})],2)],1),r("el-form-item",["add"===e.$route.query.type?r("el-button",{attrs:{type:"primary"},on:{click:e.saveRule}},[e._v("保存")]):e._e(),"edit"===e.$route.query.type?r("el-button",{attrs:{type:"primary"},on:{click:e.updateRule}},[e._v("保存")]):e._e(),"check"!=e.$route.query.type?r("el-button",{attrs:{type:"default"},on:{click:e.cancel}},[e._v("取消")]):e._e()],1)],2)},l=[],i=r("5530"),n=(r("c975"),r("a434"),r("a9e3"),r("8ba4"),r("d093")),a={name:"projectAdd",data:function(){var e=function(e,t,r){""===t?r(new Error("请输入服务名称")):(/[^\w\.\/]/g.test(t)&&r(new Error("只能输入英文字母和数字")),r())},t=function(e,t,r){""===t?r(new Error("请输入服务地址")):(/^(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])\.(\d{1,2}|1\d\d|2[0-4]\d|25[0-5])$/.test(t)||/^(?=^.{3,255}$)(http(s)?:\/\/)?(www\.)?[a-zA-Z0-9][-a-zA-Z0-9]{0,62}(\.[a-zA-Z0-9][-a-zA-Z0-9]{0,62})+(:\d+)*(\/\w+\.\w+)*$/.test(t)||r(new Error("请输入正确的服务地址")),r())},r=function(e,t,r){""===t||Number.isInteger(1*t)||r(new Error("只能输入整数")),r()},s=function(e,t,r){""===t||isNaN(Number(t))&&r(new Error("只能输入数字")),r()},l=function(e,t,r){""===t||/^\/[0-9a-zA-Z]*$/.test(t)||r(new Error("路径格式不正确")),r()};return{ruleForm:{serviceName:"",serviceType:"",serviceAddress:"",servicePort:"",serviceRules:[],serviceBreak:"",serviceLimit:"",useConsulId:"",useConsulTag:"",useConsulCheckPath:"",useConsulPort:"",useConsulInterval:"",useConsulTimeout:"",dingdingAccessToken:"",dingdingSecret:"",dingdingList:[]},serviceRules:{serviceName:[{required:!0,validator:e,trigger:"blur"}],serviceType:[{required:!0,message:"请选择服务类型",trigger:"blur"}],serviceAddress:[{required:!0,validator:t,trigger:"blur"}],servicePort:[{validator:r,trigger:"blur"}],serviceLimit:[{validator:r,trigger:"blur"}],serviceBreak:[{validator:s,trigger:"blur"}],useConsulId:[{validator:e,trigger:"blur"}],useConsulTag:[{validator:e,trigger:"blur"}],useConsulCheckPath:[{validator:l,trigger:"blur"}],useConsulInterval:[{validator:s,trigger:"blur"}],useConsulTimeout:[{validator:s,trigger:"blur"}]},rules:{apiUrl:[{required:!0,message:"请输入正确的拦截地址",validator:l,trigger:"blur"}],pathReWriteBefore:[{validator:l,trigger:"blur"}],pathReWriteUrl:[{validator:l,trigger:"blur"}]},inputVisible:!1,inputValue:"",baseInfo:"baseInfo",useConsul:"useConsul",messageWarn:"messageWarn",serverList:[]}},methods:{handleClose:function(e){"check"!=this.$route.query.type&&this.ruleForm.dingdingList.splice(this.ruleForm.dingdingList.indexOf(e),1)},showInput:function(){var e=this;this.inputVisible=!0,this.$nextTick((function(t){e.$refs.saveTagInput.$refs.input.focus()}))},handleInputConfirm:function(){var e=this.inputValue;e&&this.ruleForm.dingdingList.push(e),this.inputVisible=!1,this.inputValue=""},addRule:function(){this.ruleForm.serviceRules.push({interceptLoc:"",locationReset:"",pathReWriteBefore:"",pathReWriteUrl:""})},deleteRule:function(e){this.ruleForm.serviceRules.splice(e,1)},cancel:function(){this.$router.push({name:"projectManage"})},queryProjectType:function(){var e=this;Object(n["e"])().then((function(t){e.serverList=t.serverTypeList||[]}))},saveRule:function(){var e=this;0!=this.ruleForm.serviceRules.length?this.$refs["ruleForm"].validate((function(t){if(!t)return!1;var r=JSON.parse(JSON.stringify(e.ruleForm));r.servicePort=1*r.servicePort,r.serviceLimit=1*r.serviceLimit,r.serviceBreak=1*r.serviceBreak,r.useConsulPort=1*r.useConsulPort,r.useConsulInterval=1*r.useConsulInterval,r.useConsulTimeout=1*r.useConsulTimeout,Object(n["a"])(r).then((function(t){e.$router.push({name:"projectManage"})}))})):this.$message({message:"拦截规则不能为空",type:"warning"})},initDetail:function(){var e=this,t=this.$route.query.id;Object(n["f"])(t).then((function(t){e.ruleForm=JSON.parse(JSON.stringify(t))}))},updateRule:function(){var e=this;0!=this.ruleForm.serviceRules.length?this.$refs["ruleForm"].validate((function(t){if(!t)return!1;var r=JSON.parse(JSON.stringify(e.ruleForm));r.servicePort=1*r.servicePort,r.serviceLimit=1*r.serviceLimit,r.serviceBreak=1*r.serviceBreak,r.useConsulPort=1*r.useConsulPort,r.useConsulInterval=1*r.useConsulInterval,r.useConsulTimeout=1*r.useConsulTimeout,Object(n["g"])(r).then((function(t){e.$router.push({name:"projectManage"})}))})):this.$message({message:"拦截规则不能为空",type:"warning"})}},created:function(){this.queryProjectType(),this.$route.query.id?this.initDetail():this.ruleForm={serviceName:"",serviceType:"http",serviceAddress:"",servicePort:"",serviceRules:[],serviceBreak:"",serviceLimit:"",useConsulId:"",useConsulTag:"",useConsulCheckPath:"",useConsulPort:"",useConsulInterval:"",useConsulTimeout:"",dingdingAccessToken:"",dingdingSecret:"",dingdingList:[]}}},o=Object(i["a"])({},a),u=o,c=(r("b7e2"),r("2877")),p=Object(c["a"])(u,s,l,!1,null,"161b0760",null);t["default"]=p.exports},5899:function(e,t){e.exports="\t\n\v\f\r                　\u2028\u2029\ufeff"},"58a8":function(e,t,r){var s=r("1d80"),l=r("5899"),i="["+l+"]",n=RegExp("^"+i+i+"*"),a=RegExp(i+i+"*$"),o=function(e){return function(t){var r=String(s(t));return 1&e&&(r=r.replace(n,"")),2&e&&(r=r.replace(a,"")),r}};e.exports={start:o(1),end:o(2),trim:o(3)}},"5e89":function(e,t,r){var s=r("861d"),l=Math.floor;e.exports=function(e){return!s(e)&&isFinite(e)&&l(e)===e}},7156:function(e,t,r){var s=r("861d"),l=r("d2bb");e.exports=function(e,t,r){var i,n;return l&&"function"==typeof(i=t.constructor)&&i!==r&&s(n=i.prototype)&&n!==r.prototype&&l(e,n),e}},"7a78":function(e,t,r){},"8ba4":function(e,t,r){var s=r("23e7"),l=r("5e89");s({target:"Number",stat:!0},{isInteger:l})},a434:function(e,t,r){"use strict";var s=r("23e7"),l=r("23cb"),i=r("a691"),n=r("50c4"),a=r("7b0b"),o=r("65f0"),u=r("8418"),c=r("1dde"),p=r("ae40"),d=c("splice"),m=p("splice",{ACCESSORS:!0,0:0,1:2}),f=Math.max,v=Math.min,h=9007199254740991,g="Maximum allowed length exceeded";s({target:"Array",proto:!0,forced:!d||!m},{splice:function(e,t){var r,s,c,p,d,m,b=a(this),y=n(b.length),C=l(e,y),k=arguments.length;if(0===k?r=s=0:1===k?(r=0,s=y-C):(r=k-2,s=v(f(i(t),0),y-C)),y+r-s>h)throw TypeError(g);for(c=o(b,s),p=0;p<s;p++)d=C+p,d in b&&u(c,p,b[d]);if(c.length=s,r<s){for(p=C;p<y-s;p++)d=p+s,m=p+r,d in b?b[m]=b[d]:delete b[m];for(p=y;p>y-s+r;p--)delete b[p-1]}else if(r>s)for(p=y-s;p>C;p--)d=p+s-1,m=p+r-1,d in b?b[m]=b[d]:delete b[m];for(p=0;p<r;p++)b[p+C]=arguments[p+2];return b.length=y-s+r,c}})},a9e3:function(e,t,r){"use strict";var s=r("83ab"),l=r("da84"),i=r("94ca"),n=r("6eeb"),a=r("5135"),o=r("c6b6"),u=r("7156"),c=r("c04e"),p=r("d039"),d=r("7c73"),m=r("241c").f,f=r("06cf").f,v=r("9bf2").f,h=r("58a8").trim,g="Number",b=l[g],y=b.prototype,C=o(d(y))==g,k=function(e){var t,r,s,l,i,n,a,o,u=c(e,!1);if("string"==typeof u&&u.length>2)if(u=h(u),t=u.charCodeAt(0),43===t||45===t){if(r=u.charCodeAt(2),88===r||120===r)return NaN}else if(48===t){switch(u.charCodeAt(1)){case 66:case 98:s=2,l=49;break;case 79:case 111:s=8,l=55;break;default:return+u}for(i=u.slice(2),n=i.length,a=0;a<n;a++)if(o=i.charCodeAt(a),o<48||o>l)return NaN;return parseInt(i,s)}return+u};if(i(g,!b(" 0o1")||!b("0b1")||b("+0x1"))){for(var F,x=function(e){var t=arguments.length<1?0:e,r=this;return r instanceof x&&(C?p((function(){y.valueOf.call(r)})):o(r)!=g)?u(new b(k(t)),r,x):k(t)},I=s?m(b):"MAX_VALUE,MIN_VALUE,NaN,NEGATIVE_INFINITY,POSITIVE_INFINITY,EPSILON,isFinite,isInteger,isNaN,isSafeInteger,MAX_SAFE_INTEGER,MIN_SAFE_INTEGER,parseFloat,parseInt,isInteger".split(","),T=0;I.length>T;T++)a(b,F=I[T])&&!a(x,F)&&v(x,F,f(b,F));x.prototype=y,y.constructor=x,n(l,g,x)}},b7e2:function(e,t,r){"use strict";var s=r("7a78"),l=r.n(s);l.a},d093:function(e,t,r){"use strict";r.d(t,"d",(function(){return l})),r.d(t,"c",(function(){return i})),r.d(t,"e",(function(){return n})),r.d(t,"a",(function(){return a})),r.d(t,"f",(function(){return o})),r.d(t,"g",(function(){return u})),r.d(t,"b",(function(){return c}));var s=r("b775");function l(){return Object(s["a"])({url:"/uiApi/v1/service/serviceSum",method:"GET"})}function i(e){return Object(s["a"])({url:"/uiApi/v1/service/serviceList",method:"GET",params:e})}function n(){return Object(s["a"])({url:"/uiApi/v1/eumn/serverTypeList",method:"GET"})}function a(e){return Object(s["a"])({url:"/uiApi/v1/service/addService",method:"POST",params:e})}function o(e){return Object(s["a"])({url:"/uiApi/v1/service/serviceDetail?serverId=".concat(e),method:"GET"})}function u(e){return Object(s["a"])({url:"/uiApi/v1/service/editService",method:"POST",params:e})}function c(e){return Object(s["a"])({url:"/uiApi/v1/service/deleteService",method:"POST",params:e})}}}]);