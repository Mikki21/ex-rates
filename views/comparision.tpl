    <div class = "mycontainer">
        <h1 class="page-name">Rates rating</h1>
        {{range $key, $val := .Banks}}
            <div class="bank">
                <h2 class="bank-name">{{$val.BankName}}</h2>
                <ul class="bank-data">
                    <li>Code Alpha: {{$val.CodeAlpha}}</li>
                    <li>Rate Buy: {{$val.RateBuy}}</li>
                    <li>Rate Sale: {{$val.RateSale}}</li>
                </ul>
            </div>
        {{end}}
        <a href="/" class="back">Back</a>
    </div>
