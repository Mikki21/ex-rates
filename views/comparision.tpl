<header>
    <div class="wrap header-wrap">
        <a href="/"><h2>Exchange Rates</h2></a>
    </div>
</header>
<div class="form-window">
<div class = "wrap">
        <h2 class="page-name nametitle">Rates Rating</h2>
        {{range $key,$val := .Banks}}
            <div class="bank andbank">
                <h3 class="bank-name">{{$val.BankName}}</h3>
                <ul class="bank-data">
                    <li>Currency: {{$val.CodeAlpha}}</li>
                    <li>Rate_Buy: {{$val.RateBuy}}</li>
                    <li>Rate_Sale: {{$val.RateSale}}</li>
                </ul>
            </div>
            {{end}}
        <a href="/" class="button button-center">Back</a>
</div>
</div>
