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


<h1 class=" text-2xl">{name.toUpperCase()}</h1>
<div class="flex flex-col self-center ">
	<button class="justify center" onClick={() => generateResults()}>
		generate results
	</button>
	<ul>
		{#each results as result}
			<li>the number is {result}</li>
		{:else}
			<p>No results</p>
		{/each}
</ul>
</div>