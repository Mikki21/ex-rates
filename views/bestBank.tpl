
<header>
    <div class="wrap header-wrap">
        <div class="hamburger-logo"></div>
        <div class="header-logo">Exchange Rates</div>
    </div>
</header>

<div class="form-window">
    <hr>
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
    <h2>{{.TitleSale }}</h1>
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
    

    <hr class="bot_hr">
    <a href="/" class="button button-center">Back</a>
</div>
</div> 