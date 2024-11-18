<!-- Most of the template are copied from Alex Edwards' Let's Go book -->
<!-- I just do some small modification and add some small features -->
{{template "base" .}}
{{define "title"}}Home{{end}}
{{define "main"}}
<h2>Latest Articles:</h2>
{{if gt (len .Articles) 0}}
    <table>
        <tr>
            <th>Title</th>
            <th>Created</th>
            <th>ID</th>
            <th style="padding-right: 50px; text-align: left;">Tag</th>
        </tr>
        {{range .Articles}}
        <tr>
            <td><a href='/beegoblog/{{.Id}}'>{{.Title}}</a></td>
            <td>{{.CreateTime}}</td>
            <td>#{{.Id}}</td>
            <td style="padding-right: 50px; text-align: left;">{{.TagString}}</td>
        </tr>
        {{end}}
    </table>
    {{template "pagination" .}} <!-- use pagination template -->
{{else}}
    <p>There's nothing to see here... yet!</p>
{{end}}
{{end}}
