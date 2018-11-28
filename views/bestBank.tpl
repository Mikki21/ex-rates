<div class = "mycontainer">

<h1>Best Banks</h1>
<h2>{{.TitleBuy}}</h2>
{{range .Buy}}
<div class="bank">
<ul>
<h2 class="bank-name">{{.BankName}}</h2>
<li>
{{.CodeAlpha}}: {{.RateBuy}}
</li>
<br>
</ul>
</div>
{{end}}

<h2>{{.TitleSale}}</h1>
{{range .Sale}}
<div class="bank">
<ul>
<h2 class="bank-name">{{.BankName}}</h2>
<li>
{{.CodeAlpha}}: {{.RateSale}}
</li>
<br>
</ul>
</div>
{{end}}
<a href="/" class="back">Back</a>
</div>