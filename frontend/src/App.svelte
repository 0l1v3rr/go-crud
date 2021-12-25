<script>
	import Book from './Book.svelte';

	async function getBooks() {
		let response = await fetch("http://localhost:8080/api/books");
		let books = await response.json();
		return books;
	}

	const promise = getBooks();
</script>

<main class="container">
	<div class="row justify-content-center my-4">
		<div class="col-12">
			<h1 class="display-4">BOOKS</h1>
			<p>This is a simple CRUD application with Fiber, Svelte and Bootstrap.</p>
		</div>
		<hr>
	</div>
	<div class="row">
		{#await promise}
			<p>Loading the books...</p>
		{:then book}
			{#each book as book}
				<Book id={book.id} title={book.title} description={book.description} category={book.category} author={book.author}/>
			{/each}
		{:catch error}
			<p style="color: #ff5959">{error.message}</p>
		{/await}
	</div>
</main>

<style>

</style>