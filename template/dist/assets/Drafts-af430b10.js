import{_ as A,az as x,r,o as b,b as d,c as i,d as a,j as _,w as I,F as w,f as E,aQ as S,aR as L,q as N,k as B,U as M,t as h}from"./index-5439ded0.js";import{E as T}from"./el-checkbox-9b10d3c1.js";import{I as V}from"./inform-3c9f8d85.js";const F={class:"add"},H={class:"top-all"},J={class:"button"},O={class:"xtx-cart-page"},R={class:"container m-top-20"},j={class:"cart"},q=["onClick"],z={class:"tc message"},Q={class:"title"},U=["innerHTML"],W={class:"tc"},Z={__name:"Drafts",setup($){const p=V(),m=x(),o=r([]),l=r([]),c=r("false"),u=async()=>{const t=await S();t.data.forEach(e=>{if(e.CreatedAt!==null&&e.CreatedAt!=="0001-01-01T00:00:00Z"){const n=new Date(e.CreatedAt);e.CreatedAt=e.CreatedAt=n.toLocaleString(void 0,{year:"numeric",month:"numeric",day:"numeric",hour:"numeric",minute:"numeric"})}else e.CreatedAt=""}),l.value=t.data,g()},g=()=>{o.value=JSON.parse(JSON.stringify(l.value)),o.value.forEach(t=>{t.content=t.content.replace(/<br\s*\/?>/g,"。 ").replace(/<[^>]+>/g,"")})},C=(t,e,n)=>{t.is_read=e,o.value.filter(s=>s.is_read===!1).length<=0?c.value=!0:c.value=!1},D=t=>{c.value=t,o.value.forEach(e=>e.is_read=t)},k=async()=>{const t=o.value.filter(e=>e.is_read===!0).map(e=>e.ID);t.length>0&&(await L({id:t}),u())},y=t=>{console.log("草稿箱",t.ID,l.value);const e=l.value.find(n=>n.ID===t.ID);p.Anew(e),m.push("/recruit/word")};return b(()=>{u()}),(t,e)=>{const n=T,f=N;return d(),i("div",F,[a("div",H,[_(n,{"model-value":c.value,onChange:D},null,8,["model-value"]),a("div",J,[_(f,{onClick:k},{default:I(()=>[B("删除草稿")]),_:1})])]),a("div",O,[a("div",R,[a("div",j,[a("table",null,[a("tbody",null,[(d(!0),i(w,null,E(o.value,s=>(d(),i("tr",{key:s.ID,onClick:v=>y(s)},[a("td",null,[_(n,{"model-value":s.is_read,onChange:(v,G)=>C(s,v),onClick:e[0]||(e[0]=M(()=>{},["stop"]))},null,8,["model-value","onChange"])]),a("td",z,[a("div",Q,h(s.title),1),a("div",{class:"content",innerHTML:s.content},null,8,U)]),a("td",W,h(s.CreatedAt),1)],8,q))),128))])])])])])])}}},Y=A(Z,[["__scopeId","data-v-798d4a5a"]]);export{Y as default};