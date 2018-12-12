
<header>
    <div class="wrap header-wrap">
        <a href="/"><h2>Exchange Rates</h2></a>
    </div>
</header>

<div class="form-window">
<div class = "wrap">
    <div class="side-box">
    <div class="side">
    <h2>{{.TitleBuy }}</h2>
    {{range .Buy}}
        <div class="bank">
            <ul>
                <h3 class="bank-name">{{.BankName}}</h3>
                <li>{{.CodeAlpha}}: {{.RateBuy}}</li>
            </ul>
            
        </div>
        {{end}}
    </div>


    <div>
    <h2>{{.TitleSale }}</h2>
    {{range .Sale}}
        <div class="bank">
            <ul>
            
                <h3 class="bank-name">{{.BankName}}</h3>
                <li>{{.CodeAlpha}}: {{.RateSale}}</li>
                
            </ul>
        </div>
        {{end}}
    </div>
    
    </div>
    

    <a href="/" class="button button-center">Back</a>
</div>
</div> 