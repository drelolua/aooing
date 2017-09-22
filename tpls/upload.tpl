<!doctype html>
<html lang="en">
<head>
    <meta charset="utf-8">
    <title>Multiple file upload</title>
    <script src="https://cdn.bootcss.com/jquery/1.10.2/jquery.min.js"></script>
    <link type="text/css" rel="stylesheet" href="/static/css/aooing.css">
</head>
<body>
<h1>上传文件</h1>

<form enctype="multipart/form-data">
    Name: <input type="text" name="name"><br>
    Email: <input type="email" name="email"><br>
    <input type="file" id="avatar" name="files" multiple>
    <button type="button">保存</button>
</form>
<div id="div1">
    <!--<img id="img1" src=""></div>-->
</div>
<script>
    $('button').click(function(){
        var files = $('#avatar').prop('files');

        var data = new FormData();
        for(var i=0; i< files.length;i++){
            data.append("files", files[i])
        }


        $.ajax({
            url: '/file/upload',
            type: 'POST',
            data: data,
            cache: false,
            processData: false,
            contentType: false,
            success: function (result) {
                for(var i=0;i<result.length;i++){
                    $('#div1').append('<img src="'+result[i].url+'"></img>')
                }
            }
        });
    });
</script>
</body>
</html>