<!DOCTYPE html>
<html>
<head>
    <meta charset="utf-8">
    <title>菜鸟教程(runoob.com)</title>
    <script src="https://cdn.bootcss.com/jquery/1.10.2/jquery.min.js"></script>
    <script>
        $(document).ready(function(){
            $("button").click(function(){
                $.get("/api/v1/user/show?name=利利&age=34",function(data,status){
                    console.log("数据: " + data + "\n状态: " + status);
                    console.log(data);
                });
            });
        });
    </script>
</head>
<body>
<div>
    <h1>欢迎来到</h1>
    <button>获取你给数据</button>
</div>
</body>
</html>