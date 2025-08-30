<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="utf-8"/><meta name="viewport" content="width=device-width, initial-scale=1"/>
<title>Gracias — Next Level & Levelo</title>
<style>
  :root{--brand:#e63946;--ink:#111827;--muted:#6b7280;--bg:#fafbfc}
  *{box-sizing:border-box} body{margin:0;font-family:system-ui,Arial,sans-serif;background:var(--bg);color:var(--ink);text-align:center}
  .wrap{max-width:720px;margin:0 auto;padding:30px 18px}
  .card{background:#fff;border:1px solid #eee;border-radius:16px;box-shadow:0 10px 28px rgba(0,0,0,.06);padding:28px}
  h1{margin:0 0 10px} .muted{color:var(--muted)}
  .btn{display:inline-block;background:var(--brand);color:#fff;padding:14px 18px;border-radius:10px;font-weight:700;text-decoration:none}
  .ghost{display:inline-block;border:1px solid #ddd;background:#fff;color:#111827;padding:12px 14px;border-radius:10px;text-decoration:none}
  .panda{margin-top:10px;font-style:italic;color:#444}
  /* Modal */
  .modal{position:fixed;inset:0;background:rgba(17,24,39,.6);display:none;align-items:center;justify-content:center;padding:16px}
  .modal .box{background:#fff;max-width:560px;border-radius:16px;padding:22px;border:1px solid #eee;text-align:left}
  .modal h3{margin-top:0}
  .close{float:right;border:none;background:transparent;font-size:22px;cursor:pointer}
  ul{margin:10px 0 16px}
</style>
</head>
<body>
<div class="wrap">
  <div class="card">
    <h1 data-i18n="t1">¡Pago recibido! 🎉</h1>
    <p class="muted" data-i18n="t2">En 24h te pedimos 2–3 detalles. Tu Demo llegará en 48–72h.</p>
    <div style="display:flex;gap:12px;justify-content:center;flex-wrap:wrap;margin:16px 0">
      <a class="btn" href="[ENLACE_FACTURA]" data-i18n="b1">Descargar factura</a>
      <a class="ghost" href="#" id="upsellBtn" data-i18n="b2">Añadir Kit Comercial +$49</a>
    </div>
    <p class="panda">🐼 Levelo: <span data-i18n="bubble">¿Le damos un empujón extra a tu demo? ✨</span></p>
  </div>
  <p style="margin:16px 0"><a href="/" class="muted" data-i18n="home">Volver a la página principal</a></p>
  <div style="margin-top:8px">
    <button id="es" style="margin-right:8px">ES</button>
    <button id="en">EN</button>
  </div>
</div>

<!-- Modal Upsell -->
<div class="modal" id="upsell">
  <div class="box">
    <button class="close" id="close">×</button>
    <h3 data-i18n="u_title">Upgrade: Kit Comercial (+$49)</h3>
    <p data-i18n="u_sub">Convierte tu demo en ventas con materiales listos:</p>
    <ul>
      <li data-i18n="u_1">Mini-landing 1 sección</li>
      <li data-i18n="u_2">Guion de pitch (1 min)</li>
      <li data-i18n="u_3">5 mockups sociales</li>
    </ul>
    <a class="btn" href="[LINK_CHECKOUT_UPSELL]" data-i18n="u_cta">Añadir por $49 (1 clic)</a>
    <p class="muted" style="margin-top:8px" data-i18n="u_note">Pago seguro con Stripe/PayPal · Entrega en 72h</p>
  </div>
</div>

<script>
const dict={
 es:{t1:"¡Pago recibido! 🎉",t2:"En 24h te pedimos 2–3 detalles. Tu Demo llegará en 48–72h.",
     b1:"Descargar factura",b2:"Añadir Kit Comercial +$49",bubble:"¿Le damos un empujón extra a tu demo? ✨",
     home:"Volver a la página principal",
     u_title:"Upgrade: Kit Comercial (+$49)",u_sub:"Convierte tu demo en ventas con materiales listos:",
     u_1:"Mini-landing 1 sección",u_2:"Guion de pitch (1 min)",u_3:"5 mockups sociales",
     u_cta:"Añadir por $49 (1 clic)",u_note:"Pago seguro con Stripe/PayPal · Entrega en 72h"},
 en:{t1:"Payment received! 🎉",t2:"Within 24h we’ll ask for 2–3 details. Your Demo arrives in 48–72h.",
     b1:"Download invoice",b2:"Add Commercial Kit +$49",bubble:"Shall we give your demo an extra push? ✨",
     home:"Back to Home",
     u_title:"Upgrade: Commercial Kit (+$49)",u_sub:"Turn your demo into sales with ready-to-use assets:",
     u_1:"1-section mini landing",u_2:"1-minute pitch script",u_3:"5 social mockups",
     u_cta:"Add for $49 (1-click)",u_note:"Secure payment via Stripe/PayPal · 72h delivery"}
};
function set(l){document.querySelectorAll("[data-i18n]").forEach(el=>{
  const k=el.getAttribute("data-i18n"); el.textContent=dict[l][k]||el.textContent;
}); document.documentElement.lang=l;}
document.getElementById('es').onclick=()=>set('es');
document.getElementById('en').onclick=()=>set('en');
set('es');

const m=document.getElementById('upsell');
document.getElementById('upsellBtn').onclick=(e)=>{e.preventDefault(); m.style.display='flex'};
document.getElementById('close').onclick=()=>{m.style.display='none'};
m.addEventListener('click',(e)=>{ if(e.target===m) m.style.display='none';});
</script>
</body>
</html>
