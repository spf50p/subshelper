package tpl

var IndexHTMLTpl = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
	<meta name="description" content="Project repository: https://github.com/spf50p/subshelper">
  <title>{{ .Title }}</title>
  <style>
    :root { --bg:#0f1115; --fg:#eaeef2; --muted:#98a1b3; --card:#151923; --ring:#2a2f3a; }
    * { box-sizing: border-box; }
    html, body { height: 100%; }
    body {
      margin: 0; font: 16px/1.4 system-ui, -apple-system, Segoe UI, Roboto, sans-serif;
      background: var(--bg); color: var(--fg);
      display: grid; place-items: center;
    }
    .wrap { width: min(450px, 92vw); display: grid; gap: 16px; text-align: center; }
    h1 { font-size: 1.6rem; font-weight: 700; margin-bottom: 24px; }
    .field {
      background: var(--card); border: 1px solid var(--ring); border-radius: 14px;
      padding: 16px 14px; display: flex; justify-content: space-between; align-items: center;
    }
    hr {
      border: none; border-top: 1px solid var(--ring); margin: 20px 0;
    }
    .title { font-weight: 600; letter-spacing: .2px; }
    button {
      background: transparent; color: var(--fg); border: 1px solid var(--ring);
      padding: 6px 10px; border-radius: 8px; cursor: pointer; display: flex; align-items: center; justify-content: center;
    }
    button:hover { background: #1a2030; }
    button:active { transform: translateY(1px); }
    svg.clipboard { width: 18px; height: 18px; stroke: var(--fg); }
    .sr-only { position:absolute; width:1px; height:1px; padding:0; margin:-1px; overflow:hidden; clip:rect(0,0,0,0); white-space:nowrap; border:0; }
  </style>
</head>
<body>
  <main class="wrap">
    <h1>{{ .Title }}</h1>
    <section class="field" data-value="{{ .Url }}">
      <div class="title">{{ .TitleUrlText }}</div>
      <button type="button" class="copy">
        <svg class="clipboard" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"></rect><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"></path></svg>
      </button>
    </section>
    <hr>
		{{- range .SubLinks }}
    <section class="field" data-value="{{ .Link }}">
      <div class="title">{{ .Title }}</div>
      <button type="button" class="copy">
        <svg class="clipboard" xmlns="http://www.w3.org/2000/svg" width="16" height="16" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" stroke-linecap="round" stroke-linejoin="round"><rect width="14" height="14" x="8" y="8" rx="2" ry="2"></rect><path d="M4 16c-1.1 0-2-.9-2-2V4c0-1.1.9-2 2-2h10c1.1 0 2 .9 2 2"></path></svg>
      </button>
    </section>
		{{- end }}
  </main>
  <div id="live" class="sr-only" role="status" aria-live="polite"></div>
  <script>
    async function copyText(text) {
      try {
        await navigator.clipboard.writeText(text);
        return true;
      } catch (e) {
        const ta = document.createElement('textarea');
        ta.value = text; ta.setAttribute('readonly', '');
        ta.style.position = 'fixed'; ta.style.left = '-9999px';
        document.body.appendChild(ta); ta.select();
        let ok = false;
        try { ok = document.execCommand('copy'); } catch(_) {}
        document.body.removeChild(ta);
        return ok;
      }
    }
    function flash(el) {
      el.style.transition = 'background 300ms ease';
      const orig = getComputedStyle(el).backgroundColor;
      el.style.background = '#1e273a';
      setTimeout(() => el.style.background = orig, 350);
    }
    document.querySelectorAll('.field').forEach((card) => {
      const btn = card.querySelector('.copy');
      const value = card.dataset.value || '';
      btn.addEventListener('click', async () => {
        const ok = await copyText(value);
        const live = document.getElementById('live');
        if (ok) {
          live.textContent = 'Copied';
          flash(card);
        } else {
          live.textContent = 'Copy failed';
        }
      });
    });
  </script>
</body>
</html>
`

type HtmlIndex struct {
	Title        string
	Url          string
	TitleUrlText string
	SubLinks     []HtmlIndexSubLink
}

type HtmlIndexSubLink struct {
	Title string
	Link  string
}
