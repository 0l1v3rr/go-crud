<script>
    import { onMount } from 'svelte';

    export let id;
    export let title;
    export let description;
    export let category;
    export let author;

    let defaultId = id;
    let defaultTitle = title;
    let defaultDescription = description;
    let defaultCategory = category;
    let defaultAuthor = author;

    let card;
    let form;
    let editBtn;
    let cancel;

    onMount(() => {
        form.style.display = 'none';

        editBtn.onclick = () => {
            card.style.display = 'none';
            form.style.display = 'block';
        };

        cancel.onclick = (e) => {
            e.preventDefault();
            card.style.display = 'block';
            form.style.display = 'none';
            reset();
        }
    });

    function reset() {
        id = defaultId;
        title = defaultTitle;
        description = defaultDescription;
        category = defaultCategory;
        author = defaultAuthor;
    }

    async function deleteBook() {
        const res = await fetch(`http://localhost:8080/api/deletebook/${id}`, {
            method: 'DELETE',
            headers: { 
                "Content-Type": "application/json"
            }
        });
    }

    async function updateBook() {
        const res = await fetch(`http://localhost:8080/api/updatebook/${id}`, {
            method: 'PUT',
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

</script>

<div class="col-sm-12 col-md-6 col-lg-4 mb-3">
    <div bind:this={card} class="card">
        <div class="card-body">
            <h5 class="card-title">{ title }</h5>
            <h6 class="card-subtitle mb-1">Author: { author }</h6>
            <h6 class="card-subtitle mb-2 text-muted">{ category }</h6>
            <p class="card-text">{ description }</p>
            
            <button bind:this={editBtn} class="btn btn-primary">Edit</button>
            <form class="my-2" on:submit={deleteBook}>
                <button type="submit" class="btn btn-danger">Delete</button>
            </form>
        </div>
    </div>

    <div bind:this={form} class="card">
        <div class="card-body">
            <h5 class="card-title">Edit Book - { defaultTitle }</h5>
            <form on:submit={updateBook}>
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
                <button bind:this={cancel} class="btn btn-secondary">Cancel</button>
            </form>
        </div>
    </div>
</div>

<style>

</style>