<script>
    let id = "";
    let title = "";
    let description = "";
    let category = "";
    let author = "";
    
    async function addBook() {
        id = uuid();

        const res = await fetch("http://localhost:8080/api/addbook", {
            method: 'POST',
            headers: { 
                "Content-Type": "application/json"
            },
            body: JSON.stringify({
                id,
                title,
                description,
                category,
                author
            })
        });
    }

    function uuid() {
        return ([1e7]+-1e3+-4e3+-8e3+-1e11).replace(/[018]/g, c =>
            (c ^ crypto.getRandomValues(new Uint8Array(1))[0] & 15 >> c / 4).toString(16)
        );
    }
</script>

<main>
    <form on:submit={addBook} method="post">
        <div class="mb-3">
            <label for="title" class="form-label">Title</label>
            <input bind:value={title} type="text" class="form-control" id="title" aria-describedby="title">
        </div>
        <div class="mb-3">
            <label for="description" class="form-label">Description</label>
            <input bind:value={description} type="text" class="form-control" id="description" aria-describedby="description">
        </div>
        <div class="mb-3">
            <label for="category" class="form-label">Category</label>
            <input bind:value={category} type="text" class="form-control" id="category" aria-describedby="category">
        </div>
        <div class="mb-3">
            <label for="author" class="form-label">Author</label>
            <input bind:value={author} type="text" class="form-control" id="author" aria-describedby="author">
        </div>
        <button type="submit" class="btn btn-primary">Submit</button>
    </form>
</main>

<style>

</style>