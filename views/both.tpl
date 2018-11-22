<!DOCTYPE html>
<html>
<head>
    <title>Best</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="/static/css/main.css" rel="stylesheet">
</head>
<body>
<div class="bestBank">
<h1>Продажа</h1>
<ul>
{{range .Sale}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>
</div>

<div class="bestBank">
<h1>Покупка</h1>
<ul>
{{range .Buy}}
<li>
Банк: {{.BankName}}<br>Валюта: {{.CodeAlpha}}<br>Покупка: {{.RateBuy}}<br>Продажа: {{.RateSale}}<br>
</li>
<br>
{{end}}
</ul>
</div>
</body>
</html>
