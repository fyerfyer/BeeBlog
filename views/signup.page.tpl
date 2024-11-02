<!-- Most of the template are copied from Alex Edwards' Let's Go book -->
<!-- I just do some small modification and add some small features -->
{{template "base" .}}
{{define "title"}}Signup{{end}}
{{define "main"}}
<form action='/user/signup' method='POST' novalidate>
    <div>
        <label>Name:</label>
        {{if .NameError.HasError}}
            <label class='error'>{{.NameError.ErrMsg}}</label>
        {{end}}
        <input type='text' name='name' value="{{.Name}}">
    </div>
    <div>
        <label>Email:</label>
        {{if .EmailError.HasError}}
            <label class='error'>{{.EmailError.ErrMsg}}</label>
        {{end}}
        <input type='email' name='email' value="{{.Email}}">
    </div>
    <div>
        <label>Password:</label>
        {{if .PasswordError.HasError}}
            <label class='error'>{{.PasswordError.ErrMsg}}</label>
        {{end}}
        <input type='password' name='password' value="{{.Password}}">
    </div>
    <div>
        <input type='submit' value='Signup'>
    </div>
</form>
{{end}}