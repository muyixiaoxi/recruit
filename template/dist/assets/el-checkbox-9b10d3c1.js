import{c5 as H,bR as k,c6 as ve,bS as pn,aU as I,bQ as W,c7 as mn,aW as _,bO as Ne,aT as G,c8 as le,c9 as N,ca as je,cb as E,bP as yn,cc as Ue,c3 as hn,c4 as xn,cd as Tn,a$ as ze,ab as q,ac as ge,a7 as pe,b7 as X,x as D,y as T,a6 as Y,b8 as wn,$ as oe,C as Ve,a2 as Ke,P as He,H as We,r as Z,a_ as B,a9 as Sn,z as me,a1 as ye,bc as qe,J as O,bp as Qe,K as ie,b as S,i as Je,w as Xe,d as he,T as w,D as f,Q as j,c as $,ba as U,U as z,ce as V,S as ue,F as $n,k as Ye,t as Ze,ar as ee,av as en,I as de,W as An,aN as Cn,aO as _n,am as kn,N as En,aw as On,ax as nn}from"./index-5439ded0.js";function Ln(e){return e}var Pn=H(k,"WeakMap");const ne=Pn;function In(e,n,t){switch(t.length){case 0:return e.call(n);case 1:return e.call(n,t[0]);case 2:return e.call(n,t[0],t[1]);case 3:return e.call(n,t[0],t[1],t[2])}return e.apply(n,t)}var Bn=800,Dn=16,Mn=Date.now;function Fn(e){var n=0,t=0;return function(){var a=Mn(),s=Dn-(a-t);if(t=a,s>0){if(++n>=Bn)return arguments[0]}else n=0;return e.apply(void 0,arguments)}}function Rn(e){return function(){return e}}var Gn=ve?function(e,n){return ve(e,"toString",{configurable:!0,enumerable:!1,value:Rn(n),writable:!0})}:Ln;const Nn=Gn;var jn=Fn(Nn);const Un=jn;var xe=Math.max;function zn(e,n,t){return n=xe(n===void 0?e.length-1:n,0),function(){for(var a=arguments,s=-1,r=xe(a.length-n,0),i=Array(r);++s<r;)i[s]=a[n+s];s=-1;for(var l=Array(n+1);++s<n;)l[s]=a[s];return l[n]=t(i),In(e,this,l)}}var Vn=9007199254740991;function ce(e){return typeof e=="number"&&e>-1&&e%1==0&&e<=Vn}function Kn(e){return e!=null&&ce(e.length)&&!pn(e)}var Hn=Object.prototype;function Wn(e){var n=e&&e.constructor,t=typeof n=="function"&&n.prototype||Hn;return e===t}function qn(e,n){for(var t=-1,a=Array(e);++t<e;)a[t]=n(t);return a}var Qn="[object Arguments]";function Te(e){return I(e)&&W(e)==Qn}var tn=Object.prototype,Jn=tn.hasOwnProperty,Xn=tn.propertyIsEnumerable,Yn=Te(function(){return arguments}())?Te:function(e){return I(e)&&Jn.call(e,"callee")&&!Xn.call(e,"callee")};const fe=Yn;function Zn(){return!1}var an=typeof exports=="object"&&exports&&!exports.nodeType&&exports,we=an&&typeof module=="object"&&module&&!module.nodeType&&module,et=we&&we.exports===an,Se=et?k.Buffer:void 0,nt=Se?Se.isBuffer:void 0,tt=nt||Zn;const te=tt;var at="[object Arguments]",rt="[object Array]",st="[object Boolean]",lt="[object Date]",ot="[object Error]",it="[object Function]",ut="[object Map]",dt="[object Number]",ct="[object Object]",ft="[object RegExp]",bt="[object Set]",vt="[object String]",gt="[object WeakMap]",pt="[object ArrayBuffer]",mt="[object DataView]",yt="[object Float32Array]",ht="[object Float64Array]",xt="[object Int8Array]",Tt="[object Int16Array]",wt="[object Int32Array]",St="[object Uint8Array]",$t="[object Uint8ClampedArray]",At="[object Uint16Array]",Ct="[object Uint32Array]",h={};h[yt]=h[ht]=h[xt]=h[Tt]=h[wt]=h[St]=h[$t]=h[At]=h[Ct]=!0;h[at]=h[rt]=h[pt]=h[st]=h[mt]=h[lt]=h[ot]=h[it]=h[ut]=h[dt]=h[ct]=h[ft]=h[bt]=h[vt]=h[gt]=!1;function _t(e){return I(e)&&ce(e.length)&&!!h[W(e)]}function kt(e){return function(n){return e(n)}}var rn=typeof exports=="object"&&exports&&!exports.nodeType&&exports,P=rn&&typeof module=="object"&&module&&!module.nodeType&&module,Et=P&&P.exports===rn,Q=Et&&mn.process,Ot=function(){try{var e=P&&P.require&&P.require("util").types;return e||Q&&Q.binding&&Q.binding("util")}catch(n){}}();const $e=Ot;var Ae=$e&&$e.isTypedArray,Lt=Ae?kt(Ae):_t;const sn=Lt;var Pt=Object.prototype,It=Pt.hasOwnProperty;function Bt(e,n){var t=_(e),a=!t&&fe(e),s=!t&&!a&&te(e),r=!t&&!a&&!s&&sn(e),i=t||a||s||r,l=i?qn(e.length,String):[],c=l.length;for(var o in e)(n||It.call(e,o))&&!(i&&(o=="length"||s&&(o=="offset"||o=="parent")||r&&(o=="buffer"||o=="byteLength"||o=="byteOffset")||Ne(o,c)))&&l.push(o);return l}function Dt(e,n){return function(t){return e(n(t))}}var Mt=Dt(Object.keys,Object);const Ft=Mt;var Rt=Object.prototype,Gt=Rt.hasOwnProperty;function Nt(e){if(!Wn(e))return Ft(e);var n=[];for(var t in Object(e))Gt.call(e,t)&&t!="constructor"&&n.push(t);return n}function jt(e){return Kn(e)?Bt(e):Nt(e)}function ln(e,n){for(var t=-1,a=n.length,s=e.length;++t<a;)e[s+t]=n[t];return e}var Ce=G?G.isConcatSpreadable:void 0;function Ut(e){return _(e)||fe(e)||!!(Ce&&e&&e[Ce])}function on(e,n,t,a,s){var r=-1,i=e.length;for(t||(t=Ut),s||(s=[]);++r<i;){var l=e[r];n>0&&t(l)?n>1?on(l,n-1,t,a,s):ln(s,l):a||(s[s.length]=l)}return s}function zt(e){var n=e==null?0:e.length;return n?on(e,1):[]}function Vt(e){return Un(zn(e,void 0,zt),e+"")}function Kt(){this.__data__=new le,this.size=0}function Ht(e){var n=this.__data__,t=n.delete(e);return this.size=n.size,t}function Wt(e){return this.__data__.get(e)}function qt(e){return this.__data__.has(e)}var Qt=200;function Jt(e,n){var t=this.__data__;if(t instanceof le){var a=t.__data__;if(!N||a.length<Qt-1)return a.push([e,n]),this.size=++t.size,this;t=this.__data__=new je(a)}return t.set(e,n),this.size=t.size,this}function A(e){var n=this.__data__=new le(e);this.size=n.size}A.prototype.clear=Kt;A.prototype.delete=Ht;A.prototype.get=Wt;A.prototype.has=qt;A.prototype.set=Jt;function Xt(e,n){for(var t=-1,a=e==null?0:e.length,s=0,r=[];++t<a;){var i=e[t];n(i,t,e)&&(r[s++]=i)}return r}function Yt(){return[]}var Zt=Object.prototype,ea=Zt.propertyIsEnumerable,_e=Object.getOwnPropertySymbols,na=_e?function(e){return e==null?[]:(e=Object(e),Xt(_e(e),function(n){return ea.call(e,n)}))}:Yt;const ta=na;function aa(e,n,t){var a=n(e);return _(e)?a:ln(a,t(e))}function ke(e){return aa(e,jt,ta)}var ra=H(k,"DataView");const ae=ra;var sa=H(k,"Promise");const re=sa;var la=H(k,"Set");const se=la;var Ee="[object Map]",oa="[object Object]",Oe="[object Promise]",Le="[object Set]",Pe="[object WeakMap]",Ie="[object DataView]",ia=E(ae),ua=E(N),da=E(re),ca=E(se),fa=E(ne),C=W;(ae&&C(new ae(new ArrayBuffer(1)))!=Ie||N&&C(new N)!=Ee||re&&C(re.resolve())!=Oe||se&&C(new se)!=Le||ne&&C(new ne)!=Pe)&&(C=function(e){var n=W(e),t=n==oa?e.constructor:void 0,a=t?E(t):"";if(a)switch(a){case ia:return Ie;case ua:return Ee;case da:return Oe;case ca:return Le;case fa:return Pe}return n});const Be=C;var ba=k.Uint8Array;const De=ba;var va="__lodash_hash_undefined__";function ga(e){return this.__data__.set(e,va),this}function pa(e){return this.__data__.has(e)}function K(e){var n=-1,t=e==null?0:e.length;for(this.__data__=new je;++n<t;)this.add(e[n])}K.prototype.add=K.prototype.push=ga;K.prototype.has=pa;function ma(e,n){for(var t=-1,a=e==null?0:e.length;++t<a;)if(n(e[t],t,e))return!0;return!1}function ya(e,n){return e.has(n)}var ha=1,xa=2;function un(e,n,t,a,s,r){var i=t&ha,l=e.length,c=n.length;if(l!=c&&!(i&&c>l))return!1;var o=r.get(e),v=r.get(n);if(o&&v)return o==n&&v==e;var m=-1,y=!0,d=t&xa?new K:void 0;for(r.set(e,n),r.set(n,e);++m<l;){var u=e[m],b=n[m];if(a)var g=i?a(b,u,m,n,e,r):a(u,b,m,e,n,r);if(g!==void 0){if(g)continue;y=!1;break}if(d){if(!ma(n,function(p,x){if(!ya(d,x)&&(u===p||s(u,p,t,a,r)))return d.push(x)})){y=!1;break}}else if(!(u===b||s(u,b,t,a,r))){y=!1;break}}return r.delete(e),r.delete(n),y}function Ta(e){var n=-1,t=Array(e.size);return e.forEach(function(a,s){t[++n]=[s,a]}),t}function wa(e){var n=-1,t=Array(e.size);return e.forEach(function(a){t[++n]=a}),t}var Sa=1,$a=2,Aa="[object Boolean]",Ca="[object Date]",_a="[object Error]",ka="[object Map]",Ea="[object Number]",Oa="[object RegExp]",La="[object Set]",Pa="[object String]",Ia="[object Symbol]",Ba="[object ArrayBuffer]",Da="[object DataView]",Me=G?G.prototype:void 0,J=Me?Me.valueOf:void 0;function Ma(e,n,t,a,s,r,i){switch(t){case Da:if(e.byteLength!=n.byteLength||e.byteOffset!=n.byteOffset)return!1;e=e.buffer,n=n.buffer;case Ba:return!(e.byteLength!=n.byteLength||!r(new De(e),new De(n)));case Aa:case Ca:case Ea:return yn(+e,+n);case _a:return e.name==n.name&&e.message==n.message;case Oa:case Pa:return e==n+"";case ka:var l=Ta;case La:var c=a&Sa;if(l||(l=wa),e.size!=n.size&&!c)return!1;var o=i.get(e);if(o)return o==n;a|=$a,i.set(e,n);var v=un(l(e),l(n),a,s,r,i);return i.delete(e),v;case Ia:if(J)return J.call(e)==J.call(n)}return!1}var Fa=1,Ra=Object.prototype,Ga=Ra.hasOwnProperty;function Na(e,n,t,a,s,r){var i=t&Fa,l=ke(e),c=l.length,o=ke(n),v=o.length;if(c!=v&&!i)return!1;for(var m=c;m--;){var y=l[m];if(!(i?y in n:Ga.call(n,y)))return!1}var d=r.get(e),u=r.get(n);if(d&&u)return d==n&&u==e;var b=!0;r.set(e,n),r.set(n,e);for(var g=i;++m<c;){y=l[m];var p=e[y],x=n[y];if(a)var be=i?a(x,p,y,n,e,r):a(p,x,y,e,n,r);if(!(be===void 0?p===x||s(p,x,t,a,r):be)){b=!1;break}g||(g=y=="constructor")}if(b&&!g){var M=e.constructor,F=n.constructor;M!=F&&"constructor"in e&&"constructor"in n&&!(typeof M=="function"&&M instanceof M&&typeof F=="function"&&F instanceof F)&&(b=!1)}return r.delete(e),r.delete(n),b}var ja=1,Fe="[object Arguments]",Re="[object Array]",R="[object Object]",Ua=Object.prototype,Ge=Ua.hasOwnProperty;function za(e,n,t,a,s,r){var i=_(e),l=_(n),c=i?Re:Be(e),o=l?Re:Be(n);c=c==Fe?R:c,o=o==Fe?R:o;var v=c==R,m=o==R,y=c==o;if(y&&te(e)){if(!te(n))return!1;i=!0,v=!1}if(y&&!v)return r||(r=new A),i||sn(e)?un(e,n,t,a,s,r):Ma(e,n,c,t,a,s,r);if(!(t&ja)){var d=v&&Ge.call(e,"__wrapped__"),u=m&&Ge.call(n,"__wrapped__");if(d||u){var b=d?e.value():e,g=u?n.value():n;return r||(r=new A),s(b,g,t,a,r)}}return y?(r||(r=new A),Na(e,n,t,a,s,r)):!1}function dn(e,n,t,a,s){return e===n?!0:e==null||n==null||!I(e)&&!I(n)?e!==e&&n!==n:za(e,n,t,a,dn,s)}function Va(e,n){return e!=null&&n in Object(e)}function Ka(e,n,t){n=Ue(n,e);for(var a=-1,s=n.length,r=!1;++a<s;){var i=hn(n[a]);if(!(r=e!=null&&t(e,i)))break;e=e[i]}return r||++a!=s?r:(s=e==null?0:e.length,!!s&&ce(s)&&Ne(i,s)&&(_(e)||fe(e)))}function Ha(e,n){return e!=null&&Ka(e,n,Va)}function Wa(e,n){return dn(e,n)}function qa(e,n,t){for(var a=-1,s=n.length,r={};++a<s;){var i=n[a],l=xn(e,i);t(l,i)&&Tn(r,Ue(i,e),l)}return r}function Qa(e,n){return qa(e,n,function(t,a){return Ha(e,a)})}var Ja=Vt(function(e,n){return e==null?{}:Qa(e,n)});const Xa=Ja,cn={modelValue:{type:[Number,String,Boolean],default:void 0},label:{type:[String,Boolean,Number,Object]},indeterminate:Boolean,disabled:Boolean,checked:Boolean,name:{type:String,default:void 0},trueLabel:{type:[String,Number],default:void 0},falseLabel:{type:[String,Number],default:void 0},id:{type:String,default:void 0},controls:{type:String,default:void 0},border:Boolean,size:ze,tabindex:[String,Number],validateEvent:{type:Boolean,default:!0}},fn={[q]:e=>ge(e)||pe(e)||X(e),change:e=>ge(e)||pe(e)||X(e)},L=Symbol("checkboxGroupContextKey"),Ya=({model:e,isChecked:n})=>{const t=D(L,void 0),a=T(()=>{var r,i;const l=(r=t==null?void 0:t.max)==null?void 0:r.value,c=(i=t==null?void 0:t.min)==null?void 0:i.value;return!Y(l)&&e.value.length>=l&&!n.value||!Y(c)&&e.value.length<=c&&n.value});return{isDisabled:wn(T(()=>(t==null?void 0:t.disabled.value)||a.value)),isLimitDisabled:a}},Za=(e,{model:n,isLimitExceeded:t,hasOwnLabel:a,isDisabled:s,isLabeledByFormItem:r})=>{const i=D(L,void 0),{formItem:l}=oe(),{emit:c}=We();function o(u){var b,g;return u===e.trueLabel||u===!0?(b=e.trueLabel)!=null?b:!0:(g=e.falseLabel)!=null?g:!1}function v(u,b){c("change",o(u),b)}function m(u){if(t.value)return;const b=u.target;c("change",o(b.checked),u)}async function y(u){t.value||!a.value&&!s.value&&r.value&&(u.composedPath().some(p=>p.tagName==="LABEL")||(n.value=o([!1,e.falseLabel].includes(n.value)),await He(),v(n.value,u)))}const d=T(()=>(i==null?void 0:i.validateEvent)||e.validateEvent);return Ve(()=>e.modelValue,()=>{d.value&&(l==null||l.validate("change").catch(u=>Ke()))}),{handleChange:m,onClickRoot:y}},er=e=>{const n=Z(!1),{emit:t}=We(),a=D(L,void 0),s=T(()=>Y(a)===!1),r=Z(!1);return{model:T({get(){var l,c;return s.value?(l=a==null?void 0:a.modelValue)==null?void 0:l.value:(c=e.modelValue)!=null?c:n.value},set(l){var c,o;s.value&&B(l)?(r.value=((c=a==null?void 0:a.max)==null?void 0:c.value)!==void 0&&l.length>(a==null?void 0:a.max.value),r.value===!1&&((o=a==null?void 0:a.changeEvent)==null||o.call(a,l))):(t(q,l),n.value=l)}}),isGroup:s,isLimitExceeded:r}},nr=(e,n,{model:t})=>{const a=D(L,void 0),s=Z(!1),r=T(()=>{const o=t.value;return X(o)?o:B(o)?Sn(e.label)?o.map(me).some(v=>Wa(v,e.label)):o.map(me).includes(e.label):o!=null?o===e.trueLabel:!!o}),i=ye(T(()=>{var o;return(o=a==null?void 0:a.size)==null?void 0:o.value}),{prop:!0}),l=ye(T(()=>{var o;return(o=a==null?void 0:a.size)==null?void 0:o.value})),c=T(()=>!!(n.default||e.label));return{checkboxButtonSize:i,isChecked:r,isFocused:s,checkboxSize:l,hasOwnLabel:c}},tr=(e,{model:n})=>{function t(){B(n.value)&&!n.value.includes(e.label)?n.value.push(e.label):n.value=e.trueLabel||!0}e.checked&&t()},bn=(e,n)=>{const{formItem:t}=oe(),{model:a,isGroup:s,isLimitExceeded:r}=er(e),{isFocused:i,isChecked:l,checkboxButtonSize:c,checkboxSize:o,hasOwnLabel:v}=nr(e,n,{model:a}),{isDisabled:m}=Ya({model:a,isChecked:l}),{inputId:y,isLabeledByFormItem:d}=qe(e,{formItemContext:t,disableIdGeneration:v,disableIdManagement:s}),{handleChange:u,onClickRoot:b}=Za(e,{model:a,isLimitExceeded:r,hasOwnLabel:v,isDisabled:m,isLabeledByFormItem:d});return tr(e,{model:a}),{inputId:y,isLabeledByFormItem:d,isChecked:l,isDisabled:m,isFocused:i,checkboxButtonSize:c,checkboxSize:o,hasOwnLabel:v,model:a,handleChange:u,onClickRoot:b}},ar=["tabindex","role","aria-checked"],rr=["id","aria-hidden","name","tabindex","disabled","true-value","false-value"],sr=["id","aria-hidden","disabled","value","name","tabindex"],lr=O({name:"ElCheckbox"}),or=O({...lr,props:cn,emits:fn,setup(e){const n=e,t=Qe(),{inputId:a,isLabeledByFormItem:s,isChecked:r,isDisabled:i,isFocused:l,checkboxSize:c,hasOwnLabel:o,model:v,handleChange:m,onClickRoot:y}=bn(n,t),d=ie("checkbox"),u=T(()=>[d.b(),d.m(c.value),d.is("disabled",i.value),d.is("bordered",n.border),d.is("checked",r.value)]),b=T(()=>[d.e("input"),d.is("disabled",i.value),d.is("checked",r.value),d.is("indeterminate",n.indeterminate),d.is("focus",l.value)]);return(g,p)=>(S(),Je(en(!f(o)&&f(s)?"span":"label"),{class:w(f(u)),"aria-controls":g.indeterminate?g.controls:null,onClick:f(y)},{default:Xe(()=>[he("span",{class:w(f(b)),tabindex:g.indeterminate?0:void 0,role:g.indeterminate?"checkbox":void 0,"aria-checked":g.indeterminate?"mixed":void 0},[g.trueLabel||g.falseLabel?j((S(),$("input",{key:0,id:f(a),"onUpdate:modelValue":p[0]||(p[0]=x=>U(v)?v.value=x:null),class:w(f(d).e("original")),type:"checkbox","aria-hidden":g.indeterminate?"true":"false",name:g.name,tabindex:g.tabindex,disabled:f(i),"true-value":g.trueLabel,"false-value":g.falseLabel,onChange:p[1]||(p[1]=(...x)=>f(m)&&f(m)(...x)),onFocus:p[2]||(p[2]=x=>l.value=!0),onBlur:p[3]||(p[3]=x=>l.value=!1),onClick:p[4]||(p[4]=z(()=>{},["stop"]))},null,42,rr)),[[V,f(v)]]):j((S(),$("input",{key:1,id:f(a),"onUpdate:modelValue":p[5]||(p[5]=x=>U(v)?v.value=x:null),class:w(f(d).e("original")),type:"checkbox","aria-hidden":g.indeterminate?"true":"false",disabled:f(i),value:g.label,name:g.name,tabindex:g.tabindex,onChange:p[6]||(p[6]=(...x)=>f(m)&&f(m)(...x)),onFocus:p[7]||(p[7]=x=>l.value=!0),onBlur:p[8]||(p[8]=x=>l.value=!1),onClick:p[9]||(p[9]=z(()=>{},["stop"]))},null,42,sr)),[[V,f(v)]]),he("span",{class:w(f(d).e("inner"))},null,2)],10,ar),f(o)?(S(),$("span",{key:0,class:w(f(d).e("label"))},[ue(g.$slots,"default"),g.$slots.default?ee("v-if",!0):(S(),$($n,{key:0},[Ye(Ze(g.label),1)],64))],2)):ee("v-if",!0)]),_:3},8,["class","aria-controls","onClick"]))}});var ir=de(or,[["__file","/home/runner/work/element-plus/element-plus/packages/components/checkbox/src/checkbox.vue"]]);const ur=["name","tabindex","disabled","true-value","false-value"],dr=["name","tabindex","disabled","value"],cr=O({name:"ElCheckboxButton"}),fr=O({...cr,props:cn,emits:fn,setup(e){const n=e,t=Qe(),{isFocused:a,isChecked:s,isDisabled:r,checkboxButtonSize:i,model:l,handleChange:c}=bn(n,t),o=D(L,void 0),v=ie("checkbox"),m=T(()=>{var d,u,b,g;const p=(u=(d=o==null?void 0:o.fill)==null?void 0:d.value)!=null?u:"";return{backgroundColor:p,borderColor:p,color:(g=(b=o==null?void 0:o.textColor)==null?void 0:b.value)!=null?g:"",boxShadow:p?"-1px 0 0 0 ".concat(p):void 0}}),y=T(()=>[v.b("button"),v.bm("button",i.value),v.is("disabled",r.value),v.is("checked",s.value),v.is("focus",a.value)]);return(d,u)=>(S(),$("label",{class:w(f(y))},[d.trueLabel||d.falseLabel?j((S(),$("input",{key:0,"onUpdate:modelValue":u[0]||(u[0]=b=>U(l)?l.value=b:null),class:w(f(v).be("button","original")),type:"checkbox",name:d.name,tabindex:d.tabindex,disabled:f(r),"true-value":d.trueLabel,"false-value":d.falseLabel,onChange:u[1]||(u[1]=(...b)=>f(c)&&f(c)(...b)),onFocus:u[2]||(u[2]=b=>a.value=!0),onBlur:u[3]||(u[3]=b=>a.value=!1),onClick:u[4]||(u[4]=z(()=>{},["stop"]))},null,42,ur)),[[V,f(l)]]):j((S(),$("input",{key:1,"onUpdate:modelValue":u[5]||(u[5]=b=>U(l)?l.value=b:null),class:w(f(v).be("button","original")),type:"checkbox",name:d.name,tabindex:d.tabindex,disabled:f(r),value:d.label,onChange:u[6]||(u[6]=(...b)=>f(c)&&f(c)(...b)),onFocus:u[7]||(u[7]=b=>a.value=!0),onBlur:u[8]||(u[8]=b=>a.value=!1),onClick:u[9]||(u[9]=z(()=>{},["stop"]))},null,42,dr)),[[V,f(l)]]),d.$slots.default||d.label?(S(),$("span",{key:2,class:w(f(v).be("button","inner")),style:An(f(s)?f(m):void 0)},[ue(d.$slots,"default",{},()=>[Ye(Ze(d.label),1)])],6)):ee("v-if",!0)],2))}});var vn=de(fr,[["__file","/home/runner/work/element-plus/element-plus/packages/components/checkbox/src/checkbox-button.vue"]]);const br=Cn({modelValue:{type:_n(Array),default:()=>[]},disabled:Boolean,min:Number,max:Number,size:ze,label:String,fill:String,textColor:String,tag:{type:String,default:"div"},validateEvent:{type:Boolean,default:!0}}),vr={[q]:e=>B(e),change:e=>B(e)},gr=O({name:"ElCheckboxGroup"}),pr=O({...gr,props:br,emits:vr,setup(e,{emit:n}){const t=e,a=ie("checkbox"),{formItem:s}=oe(),{inputId:r,isLabeledByFormItem:i}=qe(t,{formItemContext:s}),l=async o=>{n(q,o),await He(),n("change",o)},c=T({get(){return t.modelValue},set(o){l(o)}});return kn(L,{...Xa(En(t),["size","min","max","disabled","validateEvent","fill","textColor"]),modelValue:c,changeEvent:l}),Ve(()=>t.modelValue,()=>{t.validateEvent&&(s==null||s.validate("change").catch(o=>Ke()))}),(o,v)=>{var m;return S(),Je(en(o.tag),{id:f(r),class:w(f(a).b("group")),role:"group","aria-label":f(i)?void 0:o.label||"checkbox-group","aria-labelledby":f(i)?(m=f(s))==null?void 0:m.labelId:void 0},{default:Xe(()=>[ue(o.$slots,"default")]),_:3},8,["id","class","aria-label","aria-labelledby"])}}});var gn=de(pr,[["__file","/home/runner/work/element-plus/element-plus/packages/components/checkbox/src/checkbox-group.vue"]]);const yr=On(ir,{CheckboxButton:vn,CheckboxGroup:gn});nn(vn);nn(gn);export{yr as E,A as S,De as U,ln as a,aa as b,Be as c,kt as d,te as e,ke as f,ta as g,on as h,Wa as i,zt as j,jt as k,Un as l,Ln as m,$e as n,zn as o,Kn as p,Wn as q,Bt as r,Yt as s,Dt as t,sn as u,fe as v,dn as w,Ha as x};
