{{template "base" .}}
{{define "title"}}Search{{end}}

{{define "main"}}
<h2>Search Articles</h2>
<form action="/beegoblog/search" method="POST">
    <div>
        <label for="search">Search:</label>
        <input type="text" id="search" name="tag" placeholder="Enter tags separated by comma..." />
    </div>
    <input type="submit" value="Search" />
</form>
{{end}}
