{{template "base" .}}
{{define "title"}}Search{{end}}
{{define "main"}}
<h2>Finding result:</h2>
{{if gt (len .Articles) 0}}
    <table>
        <tr>
            <th>Title</th>
            <th>Created</th>
            <th>ID</th>
        </tr>
        {{range .Articles}}
        <tr>
            <td><a href='/beegoblog/{{.Id}}'>{{.Title}}</a></td>
            <td>{{.CreateTime}}</td>
            <td>#{{.Id}}</td>
        </tr>
        {{end}}
    </table>
    {{template "pagination" .}} <!-- use pagination template -->
{{else}}
    <p>Haven't created article with such tags... yet!</p>
{{end}}
{{end}}
