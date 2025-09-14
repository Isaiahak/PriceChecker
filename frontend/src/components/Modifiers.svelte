<script lang="ts">
	let name = "modifiers"

	// modifiers come in via request from backend
	let modifiers: [] = $state([
		{"life":[{value:80},{toggled:true}]},
		{"armour":[{value:90},{toggled:true}]}
		])


	// once modifiers are adjusted they can be send via erquest to backend
	function searchItem(){

	}

</script>


<div class="flex flex-col gap-4 self-center items-center border-1
rounded-lg">
	<h1 class="text-center text-2xl">{name.toUpperCase()}</h1>
	<ul>
	{#each modifiers as modifier}

		{#if modifier.toggled == false}

		<li class="flex flex-row gap-4">
			<span>{modifier}</span>
			<span>modifier.value</span>
			<button> toggle</button>
		</li>
		{:else}
		<li class="flex flex-row gap-4">
			<span>{modifier}</span>
			<span> modifier.value</span>
			<button> toggle</button>
		</li>
		{/if}
	{:else}
		<p> no current modifiers</p>
	{/each}
	</ul>
</div>

