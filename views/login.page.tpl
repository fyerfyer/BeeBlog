<!-- Most of the template are copied from Alex Edwards' Let's Go book -->
<!-- I just do some small modification and add some small features -->
{{template "base" .}}
{{define "title"}}Login{{end}}
{{define "main"}}
    <form action='/user/login' method='POST' novalidate>
        {{if .Error.HasError}}
            <div class='error'>{{.Error.ErrMsg}}</div>
        {{end}}
        <div>
            <label>Email:</label>
            <input type='email' name='email' value="{{.Email}}">
        </div>
        <div>
            <label>Password:</label>
            <input type='password' name='password'>
        </div>
        <div>
            <input type='submit' value='Login'>
        </div>
    </form>
{{end}}
