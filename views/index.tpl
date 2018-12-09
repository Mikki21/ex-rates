<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="../static/css/main.css" rel="stylesheet">
</head>

<body>
  <header>
    <div class="wrap header-wrap">
      <div class="hamburger-logo"></div>
      <div class="header-logo">Exchange Rates</div>
    </div>
  </header>
    <div class="mycontainer">
      <div class="form-window">
        <div class="wrap">
        <form action = "/best" name=f1 method = "GET">
          <hr>
        <div class="choose-box">
          <div id="currency">           
            <h2>Choose your currency</h2>
            <p><input onClick="setAllCheckboxes('currency', this);" type="checkbox" class="check-box" id="all_sel"><label for="all_sel">Select All</label></p>
            <p><input type="checkbox" name="currency" value="usd" class="check-box" id="usd"><label for="usd">USD</label></p>
            <p><input type="checkbox" name="currency" value="eur" class="check-box" id="eur"><label for="eur">EUR</label></p>
            <div class="error">Choose currency</div>
          </div>
          <div class="option">
            <h2>Buy or Sale</h2>
            <p><input type="radio" name="option" value="buy" id="buy"><label for="buy">Buy</label></p>
            <p><input type="radio" name="option" value="sale" id="sale"><label for="sale">Sale</label></p>
            <p><input type="radio" name="option" value="both" id="both" checked><label for="both">Both</label></p>
          </div>
          <div id="banks">
            <h2>Choose your banks</h2>
            <p><input onClick="setAllCheckboxes('banks', this);" type="checkbox" id="all_sel_bank"><label for="all_sel_bank">Select All</label></p>
            <p><input type="checkbox" name="bank" value="privat" id="prb"><label for="prb">Privat Bank</label></p>
            <p><input type="checkbox" name="bank" value="otp" id="otpb"><label for="otpb">OTP Bank</label></p>
            <p><input type="checkbox" name="bank" value="pireus" id="pb"><label for="pb">Pireus Bank</label></p>
            <div class="error">Choose bank</div>
          </div>
        </div>
        <hr class="bot_hr">
        <div class="bot-buttons">
          <input class="button" type="submit" onclick="f1.action='/best';" value="Best choice">      
          <input class="button" type="submit" onclick="f1.action='/comparision';" value="Compare">
        </div>
        </form>
      </div>
      <div class="wrap">
        <a class="button button-center" href = "/statistics">Get statistics</a>
      </div>
    </div>   
    </div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
