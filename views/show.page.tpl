<!-- Most of the template are copied from Alex Edwards' Let's Go book -->
<!-- I just do some small modification and add some small features -->
{{template "base" .}}
{{define "title"}}Snippet #{{.Id}}{{end}}
{{define "main"}}
<div class='snippet'>
    <div class='metadata'>
        <strong>{{.Title}}</strong>
        <span>#{{.Id}}</span>
    </div>
    <pre><code>{{.Content}}</code></pre>
    <div class='metadata'>
        <time>Created: {{.CreateTime}}</time>
        <time>Expires: {{.ExpireTime}}</time>
    </div>
</div>
{{end}}
