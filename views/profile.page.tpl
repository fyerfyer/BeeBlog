{{template "base" .}}
{{define "title"}}User Profile{{end}}
    {{define "main"}}
        <h2>User Profile</h2>
    {{with .User}}
        <table>
            <tr>
                <th>Name</th>
                <td>{{.Username}}</td>
            </tr>
            <tr>
                <th>Email</th>
                <td>{{.Email}}</td>
            </tr>
            <tr>
                <th>Joined</th>
                <td>{{.CreateTime}}</td>
            </tr>
        </table>
    {{end}}
{{end}}