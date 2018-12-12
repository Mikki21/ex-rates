<!DOCTYPE html>
<html>
<head>
    <title></title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
    <link href="../static/css/main.css" rel="stylesheet">
    {{.HtmlHead}}
</head>
<body>

    
    <div class="container">
        {{.Second}}
    </div>

    <div class="container">
        {{.LayoutContent}}
    </div>
    
    <div>
        {{.SideBar}}
    </div>

    <script src="../static/js/jquery-3.3.1.min.js"></script>
        <script>
          $(window).scroll(function() {
            if ($(this).scrollTop() > 580){  
                $('header').addClass("header_screen_2");
            }else{
                $('header').removeClass("header_screen_2");
            }
            if ($(this).scrollTop() > 1420){  
                $('header').addClass("header_screen_3");
            }else{
                $('header').removeClass("header_screen_3");
            }
            });
            </script>
            {{.Scroll}}
</body>
</html>