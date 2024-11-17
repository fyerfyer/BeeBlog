<!-- Most of the template are copied from Alex Edwards' Let's Go book -->
<!-- I just do some small modification and add some small features -->
{{template "base" .}}
{{define "title"}}Create a New Snippet{{end}}
{{define "main"}}
<form action='/beegoblog/create' method='POST'>
    <div>
        <label>Title:</label>
        <input type='text' name='title'>
    </div>
    <div>
        <label>Content:</label>
        <textarea name='content'></textarea>
        </div>
    <div>
        <label>Delete in:</label>
        <input type='radio' name='expires' value='365' checked> One Year
        <input type='radio' name='expires' value='7'> One Week
        <input type='radio' name='expires' value='1'> One Day
    </div>
    <div style="display: flex; flex-direction: column; align-items: flex-start;">
        <label>Tags:</label>
        <textarea name="tags" placeholder="Enter tags separated by commas"
                  style="width: 100%; padding: 10px; height: auto; 
                  max-height: 200px; resize: none; overflow-y: auto;"></textarea>
    </div>
    <div>
        <input type='submit' value='Publish snippet'>
    </div>
</form>
{{end}}