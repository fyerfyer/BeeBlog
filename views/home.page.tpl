{{template "base" .}}
{{define "title"}}Home{{end}}
{{define "main"}}
<h2>Latest Snippets</h2>
{{if .Articles}}
    <table>
        <tr>
            <th>Title</th>
            <th>Created</th>
            <th>ID</th>
        </tr>
        {{range .Articles}}
        <tr>
            <td><a href='/beeblog/{{.Id}}'>{{.Title}}</a></td>
            <td>{{.CreateTime}}</td>
            <td>#{{.Id}}</td>
        </tr>
        {{end}}
    </table>

    <!-- pagination -->
    <div style="text-align: center;">
    <div class="pagination">
        {{if gt .CurrentPage 1}}
            <a href="/{{sub .CurrentPage 1}}">Page Up</a>
        {{end}}
        
        {{$var := .TotalPages}}
        {{range $i := until $var}}
            {{if eq $i (sub $var 1)}}
                <a class="active">{{$i | add 1}}</a> <!-- current page -->
            {{else}}
                <a href="/{{$i | add 1}}">{{$i | add 1}}</a>
            {{end}}
        {{end}}
        
        {{if lt .CurrentPage .TotalPages}}
            <a href="/{{add .CurrentPage 1}}">Page Down</a>
        {{end}}
    </div>
    </div>

{{else}}
    <p>There's nothing to see here... yet!</p>
{{end}}
{{end}}
