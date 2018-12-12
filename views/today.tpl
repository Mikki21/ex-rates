 <div class="form-window form_1st">
 <div class="wrap">
    <h2>Exchange</h2>
    <h2>Rates</h2>
    <hr id="hr-1st">
    <hr id="hr-2nd">
    <hr id="hr-3rd">
    <div class="form_1st_flex">
        <div class="form_1st_number">
            <b>{{.MainNumber}}</b>
        </div>
        <ul class="form_1st_title">
            <li><h3>&nbsp;</h3></li>
            <li><h5>Purchase</h5></li>
            <li><h5>Sale</h5></li>
        </ul>
        <ul>
            <li><h3>{{.NBU.BankName}}</h3></li>
            <li><h5>{{.NBU.RateBuy}}</h5></li>
            <li><h5>{{.NBU.RateSale}}</h5></li>
        </ul>
        <ul>
            <li><h3>{{.Others.BankName}}</h3></li>
            <li><h5>{{.Others.RateBuy}}</h5></li>
            <li><h5>{{.Others.RateSale}}</h5></li>
        </ul>
    </div>
    <a href="#screen_3" class="form_1st_pointer"></a>
</div>
</div>