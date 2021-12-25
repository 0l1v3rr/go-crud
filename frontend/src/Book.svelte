<script>
    export let id;
    export let title;
    export let description;
    export let category;
    export let author;
</script>

<div class="col-sm-12 col-md-6 col-lg-4 mb-3">
    <div class="card">
        <div class="card-body">
            <h5 class="card-title">{ title } - id: { id }</h5>
            <h6 class="card-subtitle mb-1">Author: { author }</h6>
            <h6 class="card-subtitle mb-2 text-muted">{ category }</h6>
            <p class="card-text">{ description }</p>
            <a href="/" class="card-link">Edit</a>
            <a href="/" class="card-link">Delete</a>
        </div>
    </div>
</div>

<style>

</style>