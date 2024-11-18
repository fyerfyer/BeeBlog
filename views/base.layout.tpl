<!-- Most of the template are copied from Alex Edwards' Let's Go book -->
<!-- I just do some small modification and add some small features -->
{{define "base"}}
<!doctype html>
<html lang='en'>
    <head>
        <meta charset='utf-8'>
        <title>{{template "title" .}} - BeegoBlog</title>
        <link rel='stylesheet' href='/static/css/main.css'>
        <link rel='shortcur icon' href='/static/img/favicon.ico' type='image/x-icon'>
        <link rel='stylesheet' href='https://fonts.googleapis.com/css?family=Ubuntu+Mono:400,700'>
    </head>
    <body>
        <header>
            <h1><a href='/'>BeegoBlog</a></h1>
        </header>
        <nav>
            <div>
                <a href='/'>Home</a>
                {{if .IsAuthenticated}}
                <a href='/beegoblog/create'>Create</a>
                <a href='/beegoblog/search'>Search</a>
                {{end}}
            </div>
            <div>
                {{if .IsAuthenticated}}
                <a href='/user/profile'>Profile</a>
                <form action='/user/logout' method='POST'>
                    <button>Logout</button>
                </form>
                {{else}}
                <a href='/user/signup'>Signup</a>
                <a href='/user/login'>Login</a>
                {{end}}
            </div>
        </nav>
        <main>
        {{if .Flash.success}}
        <div class='flash'>{{.Flash.success}}</div>
        {{end}}
        {{template "main" .}}
        </main>
        {{template "footer" .}}
        <script src="/static/js/main.js" type="text/javascript"></script>
    </body>
</html>
{{end}}