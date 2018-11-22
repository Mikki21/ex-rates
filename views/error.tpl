<!DOCTYPE html>

<html>
<head>
  <title>Beego</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <link href="/static/css/main.css" rel="stylesheet">
</head>

<body>
    <div class="mycontainer">
      <div class="form-window">
        <form action = "/best" name=f1 method = "GET">
        <h2 color='#FFFFFF'>{{ .error}}</h2>
          <h2>Choose your currency</h2>
          <p><label><input onClick="setAllCheckboxes('currency', this);" type="checkbox" />Select All</label></p>
          <div id="currency">
            <p><label><input type="checkbox" name="currency" value="usd">USD</label></p>
            <p><label><input type="checkbox" name="currency" value="eur">EUR</label></p>
          </div>
          <h2>Buy or Sale</h2>
          <p><label><input type="radio" name="option" value="buy">Buy</label></p>
          <p><label><input type="radio" name="option" value="sale">Sale</label></p>
          <p><label><input type="radio" name="option" value="both" checked>Both</label></p>
          <h2>Choose your banks</h2>
          <p><label><input onClick="setAllCheckboxes('banks', this);" type="checkbox" />Select All</label></p>
          <div id="banks">
            <p><label><input type="checkbox" name="bank" value="privat">Privat Bank</label></p>
            <p><label><input type="checkbox" name="bank" value="otp">OTP Bank</label></p>
            <p><label><input type="checkbox" name="bank" value="pireus">Pireus Bank</label></p>
          </div>
          <input class="button" type="submit" onclick="f1.action='/best';" value="Best choice">      
          <input class="button" type="submit" onclick="f1.action='/comparision';" value="Compare">
        </form>
      </div>
    </div>
  <script src="/static/js/reload.min.js"></script>
</body>
</html>
