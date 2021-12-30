<script>
    export let id;
    export let title;
    export let description;
    export let category;
    export let author;

    async function deleteBook() {
        const res = await fetch(`http://localhost:8080/api/deletebook/${id}`, {
            method: 'DELETE',
            headers: { 
                "Content-Type": "application/json"
            }
        });
    }

</script>

<div class="col-sm-12 col-md-6 col-lg-4 mb-3">
    <div class="card">
        <div class="card-body">
            <h5 class="card-title">{ title } - id: { id }</h5>
            <h6 class="card-subtitle mb-1">Author: { author }</h6>
            <h6 class="card-subtitle mb-2 text-muted">{ category }</h6>
            <p class="card-text">{ description }</p>
            
            <form class="my-2">
                <button type="submit" class="btn btn-primary">Edit</button>
            </form>
            <form class="my-2" on:submit={deleteBook}>
                <button type="submit" class="btn btn-danger">Delete</button>
            </form>
        </div>
    </div>
</div>

<style>

</style>