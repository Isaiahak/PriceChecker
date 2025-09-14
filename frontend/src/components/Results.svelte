<script lang="ts">
	let name: string = "results"

	let results: number[] = $state([1,2,3,4])

	function generateResults(): number[]{
		const random = Math.floor(Math.random() * 10)
		const newResults: number[] = []
		for (let i = 0; i < random; i++) {
			newResults.push(i)
		}
		results = newResults
	}

	$inspect(results)

</script>


<div class="flex flex-col self-center gap-4 mt-4 rounded-lg border-1 px-4">
	<h1 class=" text-2xl text-center">{name.toUpperCase()}</h1>
	<button class="justify center" onclick={() => generateResults()}>
		generate results
	</button>
	<ul class="flex flex-col items-start gap-4">
		{#each results as result}
			<li>the number is {result}</li>
		{:else}
			<p>No results</p>
		{/each}
	</ul>
</div>